package textfield

//#include "textfield.h"
import "C"

import (
	"github.com/mkch/rw/internal/native/darwin/control"
	"github.com/mkch/rw/native"
)

func NewTextField(x, y, w, h int) native.Handle {
	return control.NSControl_initWithFrame(native.Handle(C.NSTextField_alloc()), x, y, w, h)
}

func NSTextField_isEditable(edt native.Handle) bool {
	return bool(C.NSTextField_isEditable(C.OBJC_PTR(edt)))
}

func NSTextField_setEditable(edt native.Handle, value bool) {
	C.NSTextField_setEditable(C.OBJC_PTR(edt), C.bool(value))
}

func NSTextField_isSelectable(edt native.Handle) bool {
	return bool(C.NSTextField_isSelectable(C.OBJC_PTR(edt)))
}

func NSTextField_setSelectable(edt native.Handle, value bool) {
	C.NSTextField_setSelectable(C.OBJC_PTR(edt), C.bool(value))
}

func NSTextField_textColor(edt native.Handle) native.Handle {
	return native.Handle(C.NSTextField_textColor(C.OBJC_PTR(edt)))
}

func NSTextField_setTextColor(edt, color native.Handle) {
	C.NSTextField_setTextColor(C.OBJC_PTR(edt), C.OBJC_PTR(color))
}

func NSTextField_backgroundColor(edt native.Handle) native.Handle {
	return native.Handle(C.NSTextField_backgroundColor(C.OBJC_PTR(edt)))
}

func NSTextField_setBackgroundColor(edt, color native.Handle) {
	C.NSTextField_setBackgroundColor(C.OBJC_PTR(edt), C.OBJC_PTR(color))
}

/////////////////////////////////////

func RWTextFieldDelegate_new() native.Handle {
	return native.Handle(C.RWTextFieldDelegate_init())
}

func RWTextFieldDelegate_multiline(handle native.Handle) bool {
	return C.RWTextFieldDelegate_multiline(C.OBJC_PTR(handle)) != 0
}

func RWTextFieldDelegate_setMultiline(handle native.Handle, value bool) {
	C.RWTextFieldDelegate_setMultiline(C.OBJC_PTR(handle), C.bool(value))
}
