package deallochook

//#include "deallochook.h"
import "C"

import (
	"github.com/mkch/rw/native"
)

type DeallocHookFunc func(native.Handle)

var globalDeallocHook DeallocHookFunc

//export callGlobalDeallocHook
func callGlobalDeallocHook(ptr C.OBJC_PTR) {
	globalDeallocHook(native.Handle(ptr))
}

func Init(hook DeallocHookFunc) {
	globalDeallocHook = hook
}

func Apply(obj native.Handle) native.Handle {
	C.hookDealloc(C.OBJC_PTR(obj))
	return obj
}