package rw

import (
	"fmt"
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/internal/native/windows/acceltable"
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/internal/native/windows/window/winutil"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/util/ustr"
	"unsafe"
)

type WindowPlatformSpecific Windows_WindowMessageReceiver

type windowExtra interface {
	accelTable() *acceltable.AccelTable
}

type windowBase struct {
	objectBase
	wrapper      util.WrapperImpl
	content      Container
	menu         Menu
	dialogResult interface{}
	accel        acceltable.AccelTable
	prevWndProc  uintptr
	onClose      event.Hub
}

func (w *windowBase) accelTable() *acceltable.AccelTable {
	return &w.accel
}

func (w *windowBase) translateAccelerator(msg window.PMsg) bool {
	if table := w.accel.Handle(); table != 0 {
		return window.TranslateAccelerator(w.Wrapper().Handle(), table, msg)
	}
	return false
}

func (w *windowBase) OnClose() *event.Hub {
	return &w.onClose
}

func (w *windowBase) Enabled() bool {
	return window.IsWindowEnabled(w.Wrapper().Handle())
}

func (w *windowBase) SetEnabled(enabled bool) {
	window.EnableWindow(w.Wrapper().Handle(), enabled)
}

func (w *windowBase) Visible() bool {
	return winutil.HasStyle(w.Wrapper().Handle(), window.WS_VISIBLE)
}

func (w *windowBase) SetVisible(v bool) {
	var cmd = window.SW_HIDE
	if v {
		cmd = window.SW_SHOWNA
	}
	window.ShowWindow(w.Wrapper().Handle(), cmd)
}

func (w *windowBase) SetFrame(frame Rect) {
	setFrame(w.Wrapper().Handle(), frame.X, frame.Y, frame.Width, frame.Height)
}

func (w *windowBase) CenterToScreen() {
	centerWindowToScreen(w.Wrapper().Handle())
}

func (w *windowBase) Title() string {
	return window.GetWindowText(w.Wrapper().Handle())
}

func (w *windowBase) SetTitle(title string) {
	window.SetWindowText(w.Wrapper().Handle(), title)
}

func (w *windowBase) Menu() Menu {
	return w.menu
}

func (w *windowBase) SetMenu(menu Menu) {
	if w.menu == menu {
		return
	}
	if w.menu != nil {
		w.menu.removeAccelerators(w)
		w.menu.setWindow(nil)
	}
	if menu != nil {
		menu.setWindow(w.Self().(Window))
		window.SetMenu(w.Wrapper().Handle(), menu.Wrapper().Handle())
		menu.addAccelerators(w)
	} else {
		window.SetMenu(w.Wrapper().Handle(), 0)
	}
	window.DrawMenuBar(w.Wrapper().Handle())
	w.menu = menu
}

func (w *windowBase) ShowActive() {
	window.ShowWindow(w.Wrapper().Handle(), window.SW_SHOW)
}

func (w *windowBase) Close() {
	window.SendMessage(w.Wrapper().Handle(), window.WM_CLOSE, 0, 0)
}

func (w *windowBase) Frame() Rect {
	x, y, width, height := window.GetWindowRect(w.Wrapper().Handle())
	return Rect{x, y, width, height}
}

func (w *windowBase) Windows_PreTranslateMessage(msg window.PMsg) bool {
	switch msg.Message() {
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms646373(v=vs.85).aspx
	// TranslateAccelerator function
	// "The function translates a WM_KEYDOWN or WM_SYSKEYDOWN message to a WM_COMMAND or WM_SYSCOMMAND message ..."
	case window.WM_KEYDOWN, window.WM_SYSKEYDOWN:
		if w.translateAccelerator(msg) {
			return true
		}
	}
	return false
}

func (w *windowBase) Windows_WndProc(handle native.Handle, msg uint, wParam, lParam uintptr) uintptr {
	switch msg {
	case window.WM_DESTROY:
		if !w.wrapper.Recreating() {
			if w.onClose.HasHandler() {
				w.onClose.Send(&simpleEvent{sender: w.Self()})
			}
		}
		fmt.Printf("*hwnd=0X%X menu=%v\n", handle, w.menu)
		if w.menu != nil {
			winMenu := w.menu
			w.SetMenu(nil)
			fmt.Printf("*hwnd=0X%X destroy menu %v\n", handle, winMenu)
			winMenu.Release()
		}
		w.accel.Destroy()
		w.accel.SetOnChangedListener(nil)
	case window.WM_SIZE:
		w.resizeContent()
	case window.WM_ACTIVATE:
		if !w.wrapper.Recreating() {
			if wParam == 0 { // WA_INACTIVE
			} else {
			}
		}
	case window.WM_COMMAND:
		if lParam == 0 {
			h, l := window.HIWORD(uint(wParam)), window.LOWORD(uint(wParam))
			switch h {
			case 0: // Menu, l is menu id
				handleMenuCommand(native.Handle(l))
			case 1: // Accelerator, l is accelerator id
				handleMenuCommand(native.Handle(l))
			}
		}
	}
	return window.CallWindowProc(w.prevWndProc, handle, msg, wParam, lParam)
}

func handleMenuCommand(item native.Handle) {
	if menuItem, ok := menuItemTable.Query(item).(MenuItem); ok && menuItem.OnClick().HasHandler() {
		menuItem.OnClick().Send(&simpleEvent{sender: menuItem})
	}
}

func (w *windowBase) resizeContent() {
	if w.content == nil {
		return
	}
	x, y, width, height := window.GetClientRect(w.Wrapper().Handle())
	w.content.SetFrame(Rect{x, y, width, height})
}

func (w *windowBase) afterRegistered(event event.Event, nextHook event.Handler) bool {
	handle := w.Wrapper().Handle()
	w.prevWndProc = setWndProc(handle)
	if !event.(*util.WrapperEvent).Recreating() {
		// Create the content.
		content := newWindowContent()
		window.SetParent(content.Wrapper().Handle(), handle)
		w.content = content
		content.setWindow(w.Self().(Window))
	}
	return nextHook(event)
}

func (w *windowBase) SetContent(content Container) {
	if content == w.content {
		return
	}
	if w.content != nil {
		window.SetParent(w.content.Wrapper().Handle(), winutil.DockerWindow())
	}
	if content != nil {
		window.SetParent(content.Wrapper().Handle(), w.Wrapper().Handle())
		w.content = content
		content.setWindow(w.Self().(Window))
		w.resizeContent()
	} else {
		w.content = nil
	}
}

type WindowHandleManager struct {
	hwndManagerBase
}

var windowClsName unsafe.Pointer

func (m WindowHandleManager) Create(util.Bundle) native.Handle {
	moduleHandle := window.GetModuleHandle(nil)
	if windowClsName == nil {
		windowClsName = ustr.CStringUtf16("rw.Window")
		window.RegisterClassEx(&window.WndClassEx{
			WndProc:    window.DefWindowProcPtr(),
			Instance:   moduleHandle,
			Cursor:     window.LoadCursor(0, window.IDC_ARROW),
			Background: native.Handle(window.COLOR_WINDOW),
			ClassName:  windowClsName,
		})
	}
	return window.CreateWindowEx(0, uintptr(windowClsName), "Window", window.WS_OVERLAPPEDWINDOW,
		window.CW_USEDEFAULT, window.CW_USEDEFAULT, 300, 200,
		0, 0, moduleHandle, nil)
}

func newWindowTemplate() Window {
	w := &windowBase{}
	w.Wrapper().AfterRegistered().AddHook(w.afterRegistered)
	return w
}
