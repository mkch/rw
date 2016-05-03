package mem

//#include "mem.h"
//#include <stdlib.h>
//#include <string.h>
import "C"

import (
	"unsafe"
)

var post func(func())

func Init(f func(func())) {
	post = f
}

const smallBufferSize = unsafe.Sizeof(C.smallBuffer)
var freeList []unsafe.Pointer

func drain() {
	C.smallBufferNext = &C.smallBuffer[0]
	for _, p := range freeList {
		Free(p)
	}
	freeList = freeList[0:0]
}

func isEmpty() bool {
	return C.smallBufferNext == &C.smallBuffer[0] && len(freeList) == 0
}

func Alloc(size uintptr) unsafe.Pointer {
	if size == 0 {
		panic("0 size to alloc")
	}
	return C.calloc(C.size_t(size), 1)
}

func Free(p unsafe.Pointer) {
	C.free(p)
}

func Realloc(p unsafe.Pointer, size uintptr) unsafe.Pointer {
	if size == 0 {
		panic("0 size to alloc")
	}
	return C.realloc(p, C.size_t(size))
}

// AutoFree deallocates the pointer with C.free() after the calling function returned.
func AutoFree(p unsafe.Pointer) unsafe.Pointer {
	postDrainIfNecessary()
	freeList = append(freeList, p)
	return p
}

// CStringAutoFree calls AutoFree(C.String(str)).
func CStringAutoFree(str string) unsafe.Pointer {
	p := C.CString(str)
	return AutoFree(unsafe.Pointer(p))
}

func postDrainIfNecessary() {
	if isEmpty() {
		post(drain)
	}
}

// AllocAutoFree allocates memory and dealloccates it after the calling function returned.
func AllocAutoFree(size uintptr) (p unsafe.Pointer) {
	if size == 0 {
		panic("0 size to alloc")
	}
	postDrainIfNecessary()
	end := uintptr(unsafe.Pointer(C.smallBufferNext)) + size
	if end <= uintptr(unsafe.Pointer(&C.smallBuffer[0])) + smallBufferSize {
		p = unsafe.Pointer(C.smallBufferNext)
		C.memset(p, 0, C.size_t(size))
		C.smallBufferNext = (*C.uchar)(unsafe.Pointer(uintptr(unsafe.Pointer(C.smallBufferNext))+size))
	} else {
		p = Alloc(size)
		freeList = append(freeList, p)
	}
	return
}