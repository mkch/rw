package object

//#include "object.h"
import "C"

import (
	"github.com/mkch/rw/native"
	"unsafe"
)

func NSObject_init(h native.Handle) native.Handle {
	return native.Handle(C.NSObject_init(C.OBJC_PTR(h)))
}

func NSObject_retain(h native.Handle) native.Handle {
	C.NSObject_retain(C.OBJC_PTR(h))
	return h
}

func NSObject_retainCount(h native.Handle) int {
	return int(C.NSObject_retainCount(C.OBJC_PTR(h)))
}

func NSObject_release(h native.Handle) {
	C.NSObject_release(C.OBJC_PTR(h))
}

func NSObject_autorelease(h native.Handle) native.Handle {
	return native.Handle(C.NSObject_autorelease(C.OBJC_PTR(h)))
}

func NSObject_description(h native.Handle) string {
	return C.GoString(C.NSObject_description(C.OBJC_PTR(h)))
}

func NSObject_setAssociatedObjectRetain(h native.Handle, key unsafe.Pointer, value native.Handle) {
	C.NSObject_setAssociatedObjectRetain(C.OBJC_PTR(h), key, C.OBJC_PTR(value))
}

func NSObject_getAssociatedObject(h native.Handle, key unsafe.Pointer) native.Handle {
	return native.Handle(C.NSObject_getAssociatedObject(C.OBJC_PTR(h), key))
}

func SetDelegateRetain(obj native.Handle, delegate native.Handle) {
	C.Object_SetDelegateRetain(C.OBJC_PTR(obj), C.OBJC_PTR(delegate))
}

func Delegate(obj native.Handle) native.Handle {
	return native.Handle(C.Object_getDelegate(C.OBJC_PTR(obj)))
}

func SetTargetRetain(obj native.Handle, target native.Handle) {
	C.Object_SetTargetRetain(C.OBJC_PTR(obj), C.OBJC_PTR(target))
}

func Target(obj native.Handle) native.Handle {
	return native.Handle(C.Object_getTarget(C.OBJC_PTR(obj)))
}
