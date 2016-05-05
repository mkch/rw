package button

//#include <stdlib.h>
//#include "button.h"
import "C"


import (
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/internal/native/darwin/control"
	"github.com/mkch/rw/util/ustr"
)

func NSButton_setButtonType(b native.Handle, buttonType uint) {
	C.NSButton_setButtonType(C.OBJC_PTR(b), C.UINTPTR(buttonType))
}

func NSButton_setBezelStyle(b native.Handle, bezelStyle uint) {
	C.NSButton_setBezelStyle(C.OBJC_PTR(b), C.UINTPTR(bezelStyle))
}

func NSButton_title(b native.Handle) string {
	return C.GoString(C.NSButton_title(C.OBJC_PTR(b)))
}

func NSButton_setTitle(b native.Handle, title string) {
	C.NSButton_setTitle(C.OBJC_PTR(b), (*C.char)(ustr.CStringUtf8(title)))
}

func NewButton(x, y, w, h int) native.Handle {
	handle := control.NSControl_initWithFrame(native.Handle(C.NSButton_alloc()), x, y, w, h)
	NSButton_setButtonType(handle, 7) //NSMomentaryPushInButton		= 7
	NSButton_setBezelStyle(handle, 1) //NSRoundedBezelStyle          = 1
	return handle
}