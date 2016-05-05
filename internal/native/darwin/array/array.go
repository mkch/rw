package array

//#include "array.h"
import "C"

import (
	"github.com/mkch/rw/native"
)

// NSNotFound, index for the object "not found".
const NotFound = ^uint(0) >> 1 // 7fffffffffffffff...

func NSArray_count(a native.Handle) uint {
	return uint(C.NSArray_count(C.OBJC_PTR(a)))
}

func NSArray_objectAtIndex(a native.Handle, index uint) native.Handle {
	return native.Handle(C.NSArray_objectAtIndex(C.OBJC_PTR(a), C.ulong(index)))
}