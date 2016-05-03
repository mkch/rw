#include "post.h"
#include "cgo_export.h"

LRESULT CALLBACK postMsgWndProc(HWND hwnd,  UINT uMsg, WPARAM wParam, LPARAM lParam) {
	switch(uMsg) {
		case WM_GO_SAFE_POST:
			runSafePostedFunc(wParam);
			return 1;
		case WM_GO_UNSAFE_POST:
			runUnsafePostedFunc(wParam);
			return 1;
		default:
			return DefWindowProc(hwnd, uMsg, wParam, lParam);
	}
}

HWND createPostMessageOnlyWindow() {
	// https://msdn.microsoft.com/en-us/library/ms633574(v=VS.85).aspx#system
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms632599(v=vs.85).aspx#message_only
	HWND wnd = CreateWindowEx(0, L"BUTTON", L"", 0, WS_CHILD, 0, 0, 0, HWND_MESSAGE, NULL, GetModuleHandle(NULL), NULL);
	if(wnd) {
		SetWindowLongPtr(wnd, GWLP_WNDPROC, (LONG_PTR)postMsgWndProc);
	}
	return wnd;
}