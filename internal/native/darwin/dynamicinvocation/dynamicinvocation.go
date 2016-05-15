package dynamicinvocation

//#include <stdlib.h>
//#include "dynamicinvocation.h"
import "C"

import (
	"github.com/mkch/rw/internal/stackescape"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util/ustr"
	"unsafe"
)

type Callback func(selector string, invocationArgs native.Handle)

//export callDynamicInvocationCallback
func callDynamicInvocationCallback(id uintptr, selector *C.char, args C.OBJC_PTR) {
	if selector == nil && args == nil { // called from dealloc
		stackescape.Remove(stackescape.Id(id))
		return
	}
	f := stackescape.Get(stackescape.Id(id)).(Callback)
	f(C.GoString(selector), native.Handle(args))
}

func RWDynamicInvocation_initWithMethodsCallback(methods []string, callback Callback) native.Handle {
	var array []byte
	for _, str := range methods {
		array = append(append(array, str...), 0)
	}
	return native.Handle(C.RWDynamicInvocation_initWithMethodsUserData((*C.char)(ustr.CStringUtf8(string(array))), C.UINTPTR(stackescape.Add(callback))))
}

func RWDynamicInvocation_addMethodwithSignature(handle native.Handle, method, signature string) {
	C.RWDynamicInvocation_addMethodwithSignature(C.OBJC_PTR(handle), (*C.char)(ustr.CStringUtf8(method)), (*C.char)(ustr.CStringUtf8(signature)))
}

func RWDynamicInvocation_delegate(handle native.Handle) native.Handle {
	return native.Handle(C.RWDynamicInvocation_delegate(C.OBJC_PTR(handle)))
}

func RWDynamicInvocation_setDelegateRetain(handle, delegate native.Handle) {
	C.RWDynamicInvocation_setDelegateRetain(C.OBJC_PTR(handle), C.OBJC_PTR(delegate))
}

func RWInvocationArguments_numberOfArguments(handle native.Handle) uint {
	return uint(C.RWInvocationArguments_numberOfArguments(C.OBJC_PTR(handle)))
}

func RWInvocationArguments_getArgumentAtIndex(handle native.Handle, index uint) unsafe.Pointer {
	return C.RWInvocationArguments_getArgumentAtIndex(C.OBJC_PTR(handle), C.uint(index))
}

func RWInvocationArguments_setReturnValue(handle native.Handle, value unsafe.Pointer) {
	C.RWInvocationArguments_setReturnValue(C.OBJC_PTR(handle), value)
}
