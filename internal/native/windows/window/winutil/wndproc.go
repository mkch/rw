package winutil

import (
	"github.com/kevin-yuan/rw/native"
)

//#include <windows.h>
//#include "wndproc.h"
import "C"

type WndProc func (hwnd native.Handle, msg uint, wParam, lParam uintptr) uintptr

var wndProc WndProc

// Return the native WNDPROC that calls callback.
func NativeWndProc(callback WndProc) uintptr {
	if wndProc != nil {
		panic("Callback is already set")
	}
	wndProc = callback
	return uintptr(C.PVOID(C.GetNativeWndProc()))
}

//export goWndProc
func goWndProc(hwnd C.HWND, msg C.UINT, wParam C.WPARAM, lParam C.LPARAM) (result C.LRESULT) {
	return C.LRESULT(wndProc(native.Handle(C.PVOID(hwnd)), uint(msg), uintptr(wParam), uintptr(lParam)))
}