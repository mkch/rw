package ustr

import (
	"unsafe"
	"unicode/utf16"
)

// CStringUtf8 converts a go string to a C string(0 terminated []C.char)
// which can be passed to C.
// See "Passing pointers" section in https://golang.org/cmd/cgo/ for
// the restrictions of using the pointer in C.
func CStringUtf8(str string) unsafe.Pointer {
	// Go string is utf-8, but without terminated 0.
	zeroTerminated := append([]byte(str), 0)
	return unsafe.Pointer(&zeroTerminated[0])
}

// CStringUtf16 converts a go string to a UTF-16 C string(0 terminated []C.short)
// which can be passed to C.
// See "Passing pointers" section in https://golang.org/cmd/cgo/ for
// the restrictions of using the pointer in C.
func CStringUtf16(str string) unsafe.Pointer {
	zeroTerminated := append(utf16.Encode([]rune(str)), 0)
	return unsafe.Pointer(&zeroTerminated[0])
}

// GoStringFromUtf16 converts a UTF-16 C string(0 terminated []C.short) to go string.
// BUG(?): Hard coded 0x0FFFFFFF.
func GoStringFromUtf16(str unsafe.Pointer) string {
	var count uint
	for p:=(*uint16)(str); *p!=0; p=(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(p))+unsafe.Sizeof(uint16(0)))) {
		count++
	}
	return string(utf16.Decode((*[0x0FFFFFFF]uint16)(str)[:count]))
}
