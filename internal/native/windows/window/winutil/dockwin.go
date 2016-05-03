package winutil

import (
	"unsafe"
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/internal/native/windows/nativeutil"
	"github.com/kevin-yuan/rw/internal/native/windows/nativeutil/ustrings"
	"github.com/kevin-yuan/rw/internal/native/windows/window"
)

//#include <windows.h>
//#include "dockwin.h"
import "C"

var dockerWinClassName ustrings.Unicode
var dockerWinHandle native.Handle

func DockerWindow() native.Handle {
	if dockerWinHandle == 0 {
		if dockerWinClassName == nil {
			dockerWinClassName = ustrings.ToUnicode("rw.Dock")
			window.RegisterClassEx(&window.WndClassEx {
				WndProc: unsafe.Pointer(C.GetDockWndProc()),
				Instance: window.GetModuleHandle(nil),
				Cursor: 0,
				Background: 0,
				ClassName:	dockerWinClassName,
				})
		}
		dockerWinHandle = window.CreateWindowEx(0, uintptr(dockerWinClassName), "", window.WS_OVERLAPPEDWINDOW, 0, 0, 0, 0, native.Handle(C.PVOID(C.VarHWND_MESSAGE)), 0, window.GetModuleHandle(nil), nil)
		if dockerWinHandle == 0 {
			nativeutil.PanicWithLastError()
		}
	}
	return dockerWinHandle
}
