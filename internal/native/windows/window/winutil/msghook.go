package winutil

//#include <windows.h>
//#include "msghook.h"
import "C"

import (
	"github.com/mkch/rw/internal/native/windows/nativeutil"
	"github.com/mkch/rw/native"
)

func SetupMessageHook(ncDestroy func(handle native.Handle)) {
	ncDestroyHandler = ncDestroy
	if C.SetupMessageHook() == nil {
		nativeutil.PanicWithLastError()
	}
}

var ncDestroyHandler func(handle native.Handle)

//export msgHookNcDestroy
func msgHookNcDestroy(hwnd C.HWND) {
	ncDestroyHandler(native.Handle(C.PVOID(hwnd)))
}
