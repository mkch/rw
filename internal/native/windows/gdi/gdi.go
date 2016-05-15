package gdi

//#include <windows.h>
//#include "gdi.h"
import "C"

import (
	"github.com/mkch/rw/internal/native/windows/nativeutil"
)

var (
	PS_NULL = uintptr(C.PS_NULL)
)

var (
	NULL_BRUSH = int(C.NULL_BRUSH)
	NULL_PEN   = int(C.NULL_PEN)
)

func GetStockObject(obj int) uintptr {
	return uintptr(C.GetStockObject(C.int(obj)))
}

func Rectangle(dc uintptr, l, t, r, b int) bool {
	return C.Rectangle(C.HDC(C.PVOID(dc)), C.int(l), C.int(t), C.int(r), C.int(b)) != 0
}

func SelectObject(dc, obj uintptr) uintptr {
	return uintptr(C.SelectObject(C.HDC(C.PVOID(dc)), C.HGDIOBJ(C.PVOID(obj))))
}

func DeleteObject(obj uintptr) {
	if C.DeleteObject(C.HGDIOBJ(C.PVOID(obj))) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func CreateSolidBrush(color uint32) uintptr {
	return uintptr(C.PVOID(C.CreateSolidBrush(C.COLORREF(color))))
}

func SetTextColor(dc uintptr, color uint32) {
	if C.SetTextColor(C.HDC(C.PVOID(dc)), C.COLORREF(color)) == C.CLR_INVALID {
		nativeutil.PanicWithLastError()
	}
}

func SetBkColor(dc uintptr, color uint32) {
	if C.SetBkColor(C.HDC(C.PVOID(dc)), C.COLORREF(color)) == C.CLR_INVALID {
		nativeutil.PanicWithLastError()
	}
}
