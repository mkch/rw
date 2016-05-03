package dialog

//#include <windows.h>
//#include "dialog.h"
import "C"

import (
    "unsafe"
    "github.com/kevin-yuan/rw/native"
    "github.com/kevin-yuan/rw/internal/native/windows/window"
    "github.com/kevin-yuan/rw/internal/native/windows/nativeutil"
)

const (
	WM_INITDIALOG = uint(C.WM_INITDIALOG)
	DS_CENTER = uint(C.DS_CENTER)
    DS_NOIDLEMSG = uint(C.DS_NOIDLEMSG)
    DS_SYSMODAL = uint(C.DS_SYSMODAL)
)

const (
	IDOK = uint16(C.IDOK)
	IDCANCEL = uint16(C.IDCANCEL)
)

type DialogProc func(handle native.Handle, msg uint, wParam, lParam uintptr) bool
type MsgFilterProc func(msg window.PMsg) bool

var dialogProc DialogProc
var msgFilterProc MsgFilterProc

func Init(dlgProc DialogProc, messageFilterProc MsgFilterProc) {
    dialogProc = dlgProc
    msgFilterProc = messageFilterProc
}

//export realDlgProc
func realDlgProc(hwnd C.HWND, msg C.UINT, wParam C.WPARAM, lParam C.LPARAM) C.LRESULT {
    if dialogProc(native.Handle(C.PVOID(hwnd)), uint(msg), uintptr(wParam), uintptr(lParam)) {
        return 1
    } else {
        return 0
    }
}

func CreateDialog(instance native.Handle, parent native.Handle, style uint, exStyle uint, lParam uintptr) (ret uintptr) {
    // https://msdn.microsoft.com/en-us/library/windows/desktop/ms645394(v=vs.85).aspx
    // Remarks:
    // "In a standard template for a dialog box, the DLGTEMPLATE structure is always immediately followed by three variable-length arrays that specify the menu, class, and title for the dialog box. "
    tmpl := (*C.DLGTEMPLATE)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.DLGTEMPLATE{})+unsafe.Sizeof(C.WORD(0))*3)[0]))
    tmpl.style = C.DWORD(style)
    tmpl.dwExtendedStyle = C.DWORD(exStyle)
    // cgo: could not determine kind of name for C.DialogBoxIndirect
    return uintptr(C.DialogBoxIndirectParam(C.HINSTANCE(C.PVOID(instance)), tmpl, C.HWND(C.PVOID(parent)), C.GetDlgProc(), C.LPARAM(lParam)));
}

func EndDialog(hwnd native.Handle, result uintptr) {
    if C.EndDialog(C.HWND(C.PVOID(hwnd)), C.int(result)) == 0 {
        nativeutil.PanicWithLastError()
    }
}

//export realtDialogMsgFilterMessageProc
func realtDialogMsgFilterMessageProc(msg *C.MSG) bool {
    return msgFilterProc(window.PMsg(unsafe.Pointer(msg)))
}

// Add TranslateAccelerator to handle menu accelerators.
// The return value is HHOOK. Use UnhookWindowsHookEx() to unhook when dialog ends.
func SetupMsgFilterHookToTranslateAccelerator() native.Handle {
    return native.Handle(C.PVOID(C.SetupDialogMsgFilterHook()))
}

