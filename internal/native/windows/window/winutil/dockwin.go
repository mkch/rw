package winutil

import (
	"unsafe"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/internal/native/windows/nativeutil"
	"github.com/mkch/rw/util/ustr"
	"github.com/mkch/rw/internal/native/windows/window"
)

//#include <windows.h>
//#include "dockwin.h"
import "C"

var dockerWinHandle native.Handle

func DockerWindow() native.Handle {
	if dockerWinHandle == 0 {
		clsName := ustr.CStringUtf16("rw.Dock")
		window.RegisterClassEx(&window.WndClassEx {
			WndProc: unsafe.Pointer(C.GetDockWndProc()),
			Instance: window.GetModuleHandle(nil),
			Cursor: 0,
			Background: 0,
			ClassName:	clsName,
			})
		dockerWinHandle = window.CreateWindowEx(0, uintptr(clsName), "", window.WS_OVERLAPPEDWINDOW, 0, 0, 0, 0, native.Handle(C.PVOID(C.VarHWND_MESSAGE)), 0, window.GetModuleHandle(nil), nil)
		if dockerWinHandle == 0 {
			nativeutil.PanicWithLastError()
		}
	}
	return dockerWinHandle
}
