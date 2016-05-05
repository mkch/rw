package notificationcenter

//#include "notificationcenter.h"
//#include <stdlib.h>
import "C"

import (
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util/ustr"
)

func NSNotificationCenter_defaultCenter() native.Handle {
	return native.Handle(C.NSNotificationCenter_defaultCenter())
}

func NSNotificationCenter_addObserver_selector_name_object(c native.Handle, observer native.Handle, sel uintptr, name string, object native.Handle) {
	C.NSNotificationCenter_addObserver_selector_name_object(C.OBJC_PTR(c), C.OBJC_PTR(observer), C.PVOID(sel), (*C.char)(ustr.CStringUtf8(name)), C.OBJC_PTR(object))
}
