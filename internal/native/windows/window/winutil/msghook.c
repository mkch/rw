#include <windows.h>
#include "msghook.h"
#include "cgo_export.h"
#include "stdio.h"

static LRESULT CALLBACK CallWndProcHookProc(int nCode, WPARAM wParam, LPARAM lParam) {
    //If nCode is HC_ACTION, the hook procedure must process the message. If nCode is less than zero, the hook procedure must pass the message to the CallNextHookEx function without further processing and should return the value returned by CallNextHookEx.
    LRESULT ret = CallNextHookEx(NULL, nCode, wParam, lParam);
    if(nCode >= 0) {
        CWPSTRUCT* pData = (CWPSTRUCT*)lParam;
        if(pData->message == WM_NCDESTROY) {
            msgHookNcDestroy(pData->hwnd);
        }
    }
    return ret;
}

static HHOOK SetupMessageHook_CALLWNDPROC() {
    // Do not use WH_CALLWNDPROCRET, it will cause "fatal error: exitsyscall: syscall frame is no longer valid" at exit,
    // although, if the hook proc does not call into go, aka. removeObject(), the error will not be triggered.
    return SetWindowsHookEx(WH_CALLWNDPROC, CallWndProcHookProc, NULL, GetCurrentThreadId());
}


static LRESULT CALLBACK CallWndRetProcHookProc(int nCode, WPARAM wParam, LPARAM lParam) {
    //If nCode is HC_ACTION, the hook procedure must process the message. If nCode is less than zero, the hook procedure must pass the message to the CallNextHookEx function without further processing and should return the value returned by CallNextHookEx.
    LRESULT ret = CallNextHookEx(NULL, nCode, wParam, lParam);
    if(nCode >= 0) {
        CWPRETSTRUCT* pData = (CWPRETSTRUCT*)lParam;
        if(pData->message == WM_NCDESTROY) {
            msgHookNcDestroy(pData->hwnd);
        }
    }
    return ret;
}

static HHOOK SetupMessageHook_CALLWNDPROCRET() {
    // Do not use WH_CALLWNDPROCRET, it will cause "fatal error: exitsyscall: syscall frame is no longer valid" at exit,
    // although, if the hook proc does not call into go, aka. removeObject(), the error will not be triggered.
    return SetWindowsHookEx(WH_CALLWNDPROCRET, CallWndRetProcHookProc, NULL, GetCurrentThreadId());
}

HHOOK SetupMessageHook() {
    SetupMessageHook_CALLWNDPROCRET();
}
