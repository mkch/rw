package nativeutil

//#include <windows.h>
//#include "../types.h"
import "C"

import (
    "unsafe"
    "fmt"
    "github.com/mkch/rw/util/ustr"

)

func GetLastError() uint {
    return uint(C.GetLastError())
}


func GetLastErrorMessage(lastError uint) (errorMessage string) {
    errCode := C.DWORD(lastError)
	var lastErrorMessageBuffer unsafe.Pointer
	if C.FormatMessage(C.FORMAT_MESSAGE_ALLOCATE_BUFFER|C.FORMAT_MESSAGE_FROM_SYSTEM, nil, errCode, 0, C.LPWSTR(unsafe.Pointer(&lastErrorMessageBuffer)), 0, nil) != 0 && lastErrorMessageBuffer != nil {
		defer C.LocalFree(C.HLOCAL(lastErrorMessageBuffer))
		return ustr.GoStringFromUtf16(lastErrorMessageBuffer)
	}
    panic("FormatMessage failed!")
}

// Call panic with the error message GetLastError() returns.
func PanicWithLastError() {
    code := uint(C.GetLastError())
    msg := GetLastErrorMessage(code)
	panic(fmt.Sprintf("GetLastError() returns %d: %v", code, msg))
}

func PanicWithLastErrorCode(code uint) {
    msg := GetLastErrorMessage(code)
    panic(fmt.Sprintf("GetLastError() returns %d: %v", code, msg))
}