package ustrings

//#include <windows.h>
import "C"

import (
    "unsafe"
    "fmt"
    "github.com/kevin-yuan/rw/internal/mem"
)

type Unicode unsafe.Pointer

func toUnicode(s string, autoFree bool) Unicode {
    str := C.LPCSTR(mem.CStringAutoFree(s))
    bufCharCount := C.MultiByteToWideChar(C.CP_UTF8, 0, str, -1, nil, 0)
    if bufCharCount <= 0 {
        panic(fmt.Errorf("ToUnicode() failed: %d", uint(C.GetLastError())))
    }
    bufSize := uintptr(bufCharCount) * unsafe.Sizeof(C.WCHAR(0))
    var buf unsafe.Pointer
    if autoFree {
        buf = mem.AllocAutoFree(bufSize)
    } else {
        buf = mem.Alloc(bufSize)
    }
    if C.MultiByteToWideChar(C.CP_UTF8, 0, str, -1, C.LPWSTR(buf), bufCharCount) == 0 {
        panic(fmt.Errorf("ToUnicode() failed: %d", uint(C.GetLastError())))
    }
    return Unicode(buf)
}

func FromUnicode(wStr Unicode) string {
    bufSize := C.WideCharToMultiByte(C.CP_UTF8, 0, C.LPWSTR(wStr), -1, nil, 0, nil, nil);
    if bufSize <= 0 {
        panic(fmt.Errorf("FromUnicode() failed: %d", uint(C.GetLastError())))
    }
    buf := mem.AllocAutoFree(uintptr(bufSize));
    if C.WideCharToMultiByte(C.CP_UTF8, 0, C.LPWSTR(wStr), -1, C.LPSTR(buf), bufSize, nil, nil) == 0 {
        panic(fmt.Errorf("FromUnicode() failed: %d", uint(C.GetLastError())))
    }
    return C.GoString((*C.char)(buf))
}


func ToUnicode(str string) (wStr Unicode) {
    return toUnicode(str, false)
}

func ToUnicodeAutoFree(str string) (wStr Unicode) {
     return toUnicode(str, true)
}

