package rw

import (
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/native"
)

func centerWindowToScreen(handle native.Handle) {
	// https://msdn.microsoft.com/en-us/library/dd162826(v=vs.85).aspx
	monitorInfo := window.GetMonitorInfo(window.MonitorFromWindow(handle, window.MONITOR_DEFAULTTOPRIMARY))
	x, y, w, h := window.GetWindowRect(handle)
	x = x + (monitorInfo.WorkRectWidth-w)/2
	y = y + (monitorInfo.WorkRectHeight-h)/2
	setFrame(handle, x, y, w, h)
}

func setFrame(handle native.Handle, x, y, w, h int) {
	window.SetWindowPos(handle, 0, x, y, w, h, window.SWP_NOACTIVATE|window.SWP_NOZORDER|window.SWP_NOOWNERZORDER)
}

func setWndProc(handle native.Handle) uintptr {
	oldProc := window.SetWindowLongPtr(handle, window.GWLP_WNDPROC, nativeWndProc)
	if oldProc == nativeWndProc {
		panic("Window procedure is already set")
	}
	return oldProc
}

// Used by nativeInit in app_windows.go
func wndProc(handle native.Handle, msg uint, wParam, lParam uintptr) uintptr {
	return defaultObjectTable.Query(handle).(Windows_WindowMessageReceiver).Windows_WndProc(handle, msg, wParam, lParam)
}

// Windows_WindowMessageReceiver can be used to receive and process window messages.
// Available on Windows platform only.
type Windows_WindowMessageReceiver interface {
	// Windows_WndProc receives all window messages sent to the window procedure.
	// Only available on Windows platform.
	Windows_WndProc(handle native.Handle, msg uint, wParam, lParam uintptr) uintptr
	// Windows_PreTranslateMessage gives the window a chance to process the window messages before they are translated and dispatched.
	// msg.Hwnd() does not need to be equal with the handle of this window.
	// Only available on Windows platform.
	Windows_PreTranslateMessage(msg window.PMsg) bool
}

// hwndManagerBase is the common building block of HandleManager of HWND.
// Create method should be added to make a HandleManager.
type hwndManagerBase struct{}

func (m hwndManagerBase) Destroy(handle native.Handle) {
	window.DestroyWindow(handle)
}

func (m hwndManagerBase) Valid(handle native.Handle) bool {
	return handle != 0
}

func (m hwndManagerBase) Table() util.ObjectTable {
	return defaultObjectTable
}
