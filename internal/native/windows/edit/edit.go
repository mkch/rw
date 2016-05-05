package edit

//#include <windows.h>
import "C"

import (
	"github.com/mkch/rw/native"
	"unsafe"
)

var (
	EN_CHANGE uint = uint(C.EN_CHANGE)

	ES_MULTILINE = uint(C.ES_MULTILINE)
	ES_READONLY = uint(C.ES_READONLY)
	ES_AUTOHSCROLL = uint(C.ES_AUTOHSCROLL)
	ES_AUTOVSCROLL = uint(C.ES_AUTOVSCROLL)
)

var (
	EM_SETSEL uint = uint(C.EM_SETSEL)
	EM_SETREADONLY uint = uint(C.EM_SETREADONLY)
)

var (
	WM_CTLCOLOREDIT = uint(C.WM_CTLCOLOREDIT)
	WM_CTLCOLORSTATIC = uint(C.WM_CTLCOLORSTATIC)
)

func Edit_SetSel(handle native.Handle, start, end int) {
	C.SendMessage(C.HWND(C.PVOID(handle)), C.EM_SETSEL, C.WPARAM(C.int(start)), C.LPARAM(C.int(end)))
}

func Edit_GetSel(handle native.Handle)(start, end int) {
	var dwStart, dwEnd C.DWORD;
	C.SendMessage(C.HWND(C.PVOID(handle)), C.EM_GETSEL, C.WPARAM(uintptr(unsafe.Pointer(&dwStart))), C.LPARAM(uintptr(unsafe.Pointer(&dwEnd))))
	return int(dwStart), int(dwEnd)
}