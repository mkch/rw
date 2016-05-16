package messagebox

import (
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util/ustr"
)

//#include <windows.h>
import "C"

type Type uint32

const (
	MB_ABORTRETRYIGNORE  = Type(C.MB_ABORTRETRYIGNORE)
	MB_CANCELTRYCONTINUE = Type(C.MB_CANCELTRYCONTINUE)
	MB_HELP              = Type(C.MB_HELP)
	MB_OK                = Type(C.MB_OK)
	MB_OKCANCEL          = Type(C.MB_OKCANCEL)
	MB_RETRYCANCEL       = Type(C.MB_RETRYCANCEL)
	MB_YESNO             = Type(C.MB_YESNO)
	MB_YESNOCANCEL       = Type(C.MB_YESNOCANCEL)

	MB_ICONWARNING     = Type(C.MB_ICONWARNING)
	MB_ICONINFORMATION = Type(C.MB_ICONINFORMATION)
	MB_ICONQUESTION    = Type(C.MB_ICONQUESTION)
	MB_ICONERROR       = Type(C.MB_ICONERROR)

	MB_DEFBUTTON1 = Type(C.MB_DEFBUTTON1)
	MB_DEFBUTTON2 = Type(C.MB_DEFBUTTON2)
	MB_DEFBUTTON3 = Type(C.MB_DEFBUTTON3)
	MB_DEFBUTTON4 = Type(C.MB_DEFBUTTON4)

	MB_APPLMODAL   = Type(C.MB_APPLMODAL)
	MB_SYSTEMMODAL = Type(C.MB_SYSTEMMODAL)
	MB_TASKMODAL   = Type(C.MB_TASKMODAL)
)

type ReturnCode int

const (
	IDABORT    = ReturnCode(C.IDABORT)
	IDCANCEL   = ReturnCode(C.IDCANCEL)
	IDCONTINUE = ReturnCode(C.IDCONTINUE)
	IDIGNORE   = ReturnCode(C.IDIGNORE)
	IDNO       = ReturnCode(C.IDNO)
	IDOK       = ReturnCode(C.IDOK)
	IDRETRY    = ReturnCode(C.IDRETRY)
	IDTRYAGAIN = ReturnCode(C.IDTRYAGAIN)
	IDYES      = ReturnCode(C.IDYES)
)

func MessageBox(hwnd native.Handle, text, caption string, utype Type) ReturnCode {
	return ReturnCode(C.MessageBox(C.HWND(C.PVOID(hwnd)), C.LPWSTR(ustr.CStringUtf16(text)), C.LPWSTR(ustr.CStringUtf16(caption)), C.UINT(utype)))
}
