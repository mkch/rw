package rw

import (
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/internal/native/windows/acceltable"
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/internal/native/windows/window/winutil"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"unicode"
)

type WindowPlatformSpecific Windows_WindowMessageReceiver

type windowExtra interface {
	addMenuItemAccelerator(item MenuItem)
	removeMenuItemAccelerator(id uint16)
}

type windowBase struct {
	objectBase
	wrapper       util.WrapperImpl
	content       Container
	menu          Menu
	dialogResult  interface{}
	accel         acceltable.AccelTable
	prevWndProc   uintptr
	onClose       event.Hub
	menuItemAccel map[uint16]MenuItem // Accel id to MenuItem
}

func (w *windowBase) addMenuItemAccelerator(item MenuItem) {
	var id uint16
	// Use 100 as the first id to avoid some predefined IDs, ID_OK etc.
	for id = uint16(100); id < 0xFFFF; id++ {
		if _, exists := w.menuItemAccel[id]; !exists {
			break
		}
	}
	if id == 0 {
		panic("Run out of accelerator id")
	}
	mod, key := item.KeyboardShortcut()
	// http://stackoverflow.com/questions/23592079/why-does-createacceleratortable-not-work-without-fvirtkey
	//https://msdn.microsoft.com/en-us/library/windows/desktop/dd375731(v=vs.85).aspx
	var fVirt byte = acceltable.FVIRTKEY
	k := unicode.ToUpper(key) // Virtual key code.
	if mod&ControlKey != 0 {
		fVirt |= acceltable.FCONTROL
	}
	if mod&AltKey != 0 {
		fVirt |= acceltable.FALT
	}
	if mod&ShiftKey != 0 {
		fVirt |= acceltable.FSHIFT
	}
	w.accel.Add(fVirt, uint16(k), id)
	w.menuItemAccel[id] = item
	item.setId(id)
}

func (w *windowBase) removeMenuItemAccelerator(id uint16) {
	w.accel.Remove(id)
	w.menuItemAccel[id].setId(0)
	delete(w.menuItemAccel, id)
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
	handle := w.Wrapper().Handle()
	if menu != nil {
		menu.setWindow(w.Self().(Window))
		window.SetMenu(handle, menu.Wrapper().Handle())
		menu.addAccelerators(w)
	} else {
		window.SetMenu(handle, 0)
	}
	window.DrawMenuBar(handle)
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
		if w.menu != nil {
			winMenu := w.menu
			w.SetMenu(nil)
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
			high, low := window.HIWORD(uint(wParam)), window.LOWORD(uint(wParam))
			switch high {
			case 0: // Menu, low is menu id
				// Do nothing. WM_MENUCOMMAND is used for menus items(see SetMenuInfo and MNS_NOTIFYBYPOS).
			case 1: // Accelerator, low is accelerator id
				w.handleAccelCommand(low)
			}
		}
	case window.WM_MENUCOMMAND:
		w.handleMenuCommand(native.Handle(lParam), int(wParam))
	}
	return window.CallWindowProc(w.prevWndProc, handle, msg, wParam, lParam)
}

func (w *windowBase) handleAccelCommand(id uint16) {
	if menuItem, exists := w.menuItemAccel[id]; exists && menuItem.OnClick().HasHandler() {
		menuItem.OnClick().Send(&simpleEvent{sender: menuItem})
	}
}

func (w *windowBase) handleMenuCommand(menuHandle native.Handle, index int) {
	if menu, ok := menuTable.Query(menuHandle).(Menu); ok {
		if menuItem := menu.Item(index); menuItem.OnClick().HasHandler() {
			menuItem.OnClick().Send(&simpleEvent{sender: menuItem})
		}
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

type windowHandleManager struct {
	hwndManager
}

func (m windowHandleManager) Create(b util.Bundle) native.Handle {
	if b != nil {
		if handle, ok := b["rw:dlg-handle"].(native.Handle); ok && m.Valid(handle) {
			return handle
		}
	}
	return m.hwndManager.Create(b)
}

func initWindow(w *windowBase, createHandleFunc func(util.Bundle) native.Handle) *windowBase {
	w.menuItemAccel = make(map[uint16]MenuItem)
	w.wrapper.SetHandleManager(windowHandleManager{hwndManager(createHandleFunc)})
	w.wrapper.AfterRegistered().AddHook(w.afterRegistered)
	return w
}

func allocWindow(createHandleFunc func(util.Bundle) native.Handle) Window {
	return initWindow(&windowBase{}, createHandleFunc)
}
