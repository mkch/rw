package screen

//#include "screen.h"
import "C"

import (
	"github.com/mkch/rw/native"
)

func NSScreen_mainScreen() native.Handle {
	return native.Handle(C.NSScreen_mainScreen())
}

func NSScreen_visibleFrame(s native.Handle) (x, y, width, height int) {
	rect := C.NSScreen_visibleFrame(C.OBJC_PTR(s))
	return int(rect.origin.x), int(rect.origin.y), int(rect.size.width), int(rect.size.height)
}

func NSScreen_frame(s native.Handle) (x, y, width, height int) {
	rect := C.NSScreen_frame(C.OBJC_PTR(s))
	return int(rect.origin.x), int(rect.origin.y), int(rect.size.width), int(rect.size.height)
}