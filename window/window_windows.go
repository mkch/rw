package window

import (
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/util/ustr"
	"unsafe"
)

var windowClsName unsafe.Pointer

func createWindow(util.Bundle) native.Handle {
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
