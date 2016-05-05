package rw

import (
	"github.com/mkch/rw/internal/native/windows/commcontrol"
	"github.com/mkch/rw/internal/native/windows/dialog"
	"github.com/mkch/rw/internal/native/windows/post"
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/internal/native/windows/window/winutil"
	"github.com/mkch/rw/native"
	"io"
)

var nativeWndProc uintptr

func nativeInit() {
	commcontrol.Initialize()
	post.Init()
	winutil.SetupMessageHook(func(handle native.Handle) {
		defaultObjectTable.Remove(handle)
	})
	dialog.Init(dialogProc, dlgMsgFilterProc)
	nativeWndProc = winutil.NativeWndProc(wndProc)
}

func nativeRun(initializeCallback, terminateCallback func()) {
	initializeCallback()
	msgLoop()
	if terminateCallback != nil {
		terminateCallback()
	}
}

func msgLoop() uintptr {
	msg := window.AllocMsg()
	defer msg.Free()

	for window.GetMessage(msg, 0, 0, 0) {
		if msg.Hwnd() == 0 {
			continue // Thread message.
		}

		if obj := defaultObjectTable.Query(msg.Hwnd()); obj != nil {
			if obj.(Windows_WindowMessageReceiver).Windows_PreTranslateMessage(msg) {
				continue
			}
		}
		if rootWin := window.GetAncestor(msg.Hwnd(), window.GA_ROOT); window.IsDialogMessage(rootWin, msg) {
			continue
		}
		window.TranslateMessage(msg)
		window.DispatchMessage(msg)
	}
	return msg.WParam()
}

func nativeExit() {
	window.PostQuitMessage(0)
}

func objectsLeaked() bool {
	return !defaultObjectTable.Empty() || !menuTable.Empty() || !menuItemTable.Empty()
}

func printLeakedObjects(w io.Writer) {
	if !defaultObjectTable.Empty() {
		defaultObjectTable.Print("default object table", w)
	}
	if !menuTable.Empty() {
		menuTable.Print("menu table", w)
	}
	if !menuItemTable.Empty() {
		menuItemTable.Print("menu item table", w)
	}
}
