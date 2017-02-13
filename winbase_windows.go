package rw

import (
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
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

// hwndManager is the common building block of all concrete HandleManagers for HWND.
// Create method calls hwndManager itself.
// Converting a `func(util.Bundle) native.Handle` to hwndManager makes a useable hwndManager.
type hwndManager func(util.Bundle) native.Handle

func (m hwndManager) Create(b util.Bundle) native.Handle {
	return m(b)
}

func (m hwndManager) Destroy(handle native.Handle) {
	window.DestroyWindow(handle)
}

func (m hwndManager) Valid(handle native.Handle) bool {
	return handle != m.Invalid()
}

func (m hwndManager) Invalid() native.Handle {
	return 0
}

func (m hwndManager) Table() *util.ObjectTable {
	return defaultObjectTable
}
