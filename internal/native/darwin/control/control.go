package control

//#include <stdlib.h>
//#include "control.h"
import "C"

import (
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util/ustr"
)

func NSControl_initWithFrame(c native.Handle, x, y, w, h int) native.Handle {
	rect := C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(w), C.CGFloat(h))
	return native.Handle(C.NSControl_initWithFrame(C.OBJC_PTR(c), rect))
}

func NSControl_setTarget(c, target native.Handle) {
	C.NSControl_setTarget(C.OBJC_PTR(c), C.OBJC_PTR(target))
}

func NSControl_target(c native.Handle) native.Handle {
	return native.Handle(C.NSControl_target(C.OBJC_PTR(c)))
}

func NSControl_action(c native.Handle) uintptr {
	return uintptr(C.NSControl_action(C.OBJC_PTR(c)))
}

func NSControl_setAction(c native.Handle, sel uintptr) {
	C.NSControl_setAction(C.OBJC_PTR(c), C.UINTPTR(sel))
}

func NSControl_setStringValue(c native.Handle, value string) {
	C.NSControl_setStringValue(C.OBJC_PTR(c), (*C.char)(ustr.CStringUtf8(value)))
}

func NSControl_stringValue(c native.Handle) string {
	return C.GoString(C.NSControl_stringValue(C.OBJC_PTR(c)))
}

func NSControl_isEnabled(c native.Handle) bool {
	return bool(C.NSControl_isEnabled(C.OBJC_PTR(c)))
}

func Control_setEnabled(c native.Handle, enabled bool) {
	var value C.bool
	if enabled {
		value = C.bool(true)
	}
	C.NSControl_setEnabled(C.OBJC_PTR(c), value)
}

func NSControl_currentEditor(c native.Handle) native.Handle {
	return native.Handle(C.NSControl_currentEditor(C.OBJC_PTR(c)))
}
