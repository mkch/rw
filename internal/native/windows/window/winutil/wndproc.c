#include <windows.h>
#include "cgo_export.h"
#include "wndproc.h"

static LRESULT CALLBACK NativeWndProc(HWND hwnd,  UINT uMsg, WPARAM wParam, LPARAM lParam) {
    return goWndProc(hwnd, uMsg, wParam, lParam);
}

WNDPROC GetNativeWndProc() {
	return NativeWndProc;
}