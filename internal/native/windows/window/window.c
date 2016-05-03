#include <windows.h>
#include <stdio.h>
#include "window.h"
#include "cgo_export.h"

static BOOL CALLBACK EnumChildProc(HWND hwnd, LPARAM lParam) {
    return callEnumChildWindowCallback(hwnd, lParam);
}

WNDENUMPROC GetEnumChildProc() {
    return EnumChildProc;
}

// When compiling for 32-bit Windows, SetWindowLongPtr is defined as a call to the SetWindowLong function.
//
// Last parameter of SetWindowLong function is of type LONG instend of LONG_PTR, this means we can't call
// this function consistently on both 32-bit and 64-bit Windows.
LONG_PTR WINAPI SetWindowLongPtr__LONG_PTR_HACK(HWND hWnd, int nIndex, LONG_PTR dwNewLong) {
    return SetWindowLongPtr(hWnd, nIndex, dwNewLong);
}

LONG_PTR WINAPI SetClassLongPtr__LONG_PTR_HACK(HWND hWnd, int nIndex, LONG_PTR dwNewLong) {
    return SetClassLongPtr(hWnd, nIndex, dwNewLong);
}

WORD FnHIWORD(DWORD value) {
    return HIWORD(value);
}

WORD FnLOWWORD(DWORD value) {
    return LOWORD(value);
}

int GetCW_USEDEFAULT() {
	// int overflow in go code.
	return CW_USEDEFAULT;
}

WNDPROC GetDefWindowProcPtr() {
	return DefWindowProc;
}

LPCTSTR VarIDC_ARROW = IDC_ARROW;
UINT VarHTTRANSPARENT = HTTRANSPARENT;
HWND VarHWND_BOTTOM = HWND_BOTTOM;
HWND VarHWND_NOTOPMOST = HWND_NOTOPMOST;
HWND VarHWND_TOP = HWND_TOP;
HWND VarHWND_TOPMOST = HWND_TOPMOST;