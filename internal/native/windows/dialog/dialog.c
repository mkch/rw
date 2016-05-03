#include <windows.h>
#include "cgo_export.h"

static INT_PTR CALLBACK DlgProc(HWND hwnd,  UINT uMsg, WPARAM wParam, LPARAM lParam) {
    // Call into go.
    return realDlgProc(hwnd, uMsg, wParam, lParam);
}

DLGPROC GetDlgProc() {
    return DlgProc;
}

static LRESULT CALLBACK DialogMsgFilterMessageProc(int code, WPARAM wParam, LPARAM lParam) {
    if(code >= 0 && realtDialogMsgFilterMessageProc((LPMSG)lParam)) {
        return TRUE;
    }
    return CallNextHookEx(NULL, code, wParam, lParam);
}

HHOOK SetupDialogMsgFilterHook() {
    return SetWindowsHookEx(WH_MSGFILTER, &DialogMsgFilterMessageProc, 0, GetCurrentThreadId());
}
