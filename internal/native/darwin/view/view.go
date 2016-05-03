package view

//#include "view.h"
import "C"

import (
	"github.com/kevin-yuan/rw/native"
	nativecolor "github.com/kevin-yuan/rw/internal/native/darwin/color"
	"image/color"
)

var ViewFrameDidChangeNotification = C.GoString(C.NSView_NSViewFrameDidChangeNotification())

func RWFlippedView_alloc() native.Handle {
	return native.Handle(C.RWFlippedView_alloc())
}

func RWFlippedView_backgroundColor(v native.Handle) color.Color {
	return nativecolor.NSColor_getRGBA(native.Handle(C.RWFlippedView_backgroundColor(C.OBJC_PTR(v))))
}

func RWFlippedView_setBackgroundColor(v native.Handle, c color.Color) {
	C.RWFlippedView_setBackgroundColor(C.OBJC_PTR(v), C.OBJC_PTR(nativecolor.NSColor_colorWithRGB(c)))
}

func RWFlippedView_setAcceptFirstResponder(v native.Handle, accept bool) {
	C.RWFlippedView_setAcceptFirstResponder(C.OBJC_PTR(v), C.bool(accept))
}

func NSView_initWithFrame(handle native.Handle, x, y, w, h int) native.Handle {
	return native.Handle(C.NSView_initWithFrame(C.OBJC_PTR(handle), C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(w), C.CGFloat(h))))
}

func NSView_addSubview(v native.Handle, subView native.Handle) {
	C.NSView_addSubview(C.OBJC_PTR(v), C.OBJC_PTR(subView))
}

func NSView_removeFromSuperview(v native.Handle) {
	C.NSView_removeFromSuperview(C.OBJC_PTR(v))
}

func NSView_superview(v native.Handle) native.Handle {
	return native.Handle(C.NSView_superview(C.OBJC_PTR(v)))
}

func NSView_window(v native.Handle) native.Handle {
	return native.Handle(C.NSView_window(C.OBJC_PTR(v)))
}

func NSView_frame(v native.Handle)(x, y, width, height int) {
	rect := C.NSView_frame(C.OBJC_PTR(v))
	return int(rect.origin.x), int(rect.origin.y), int(rect.size.width), int(rect.size.height)
}

func NSView_setFrameSize(v native.Handle, cx, cy int) {
	C.NSView_setFrameSize(C.OBJC_PTR(v), C.CGSizeMake(C.CGFloat(cx), C.CGFloat(cy)))
}

func NSView_setFrameOrigin(v native.Handle, x, y int) {
	C.NSView_setFrameOrigin(C.OBJC_PTR(v), C.CGPointMake(C.CGFloat(x), C.CGFloat(y)))
}

func NSView_isHidden(v native.Handle) bool {
	return C.NSView_isHidden(C.OBJC_PTR(v)) != false;
}

func NSView_setHidden(v native.Handle, hidden bool) {
	var h C.bool = false
	if hidden {
		h = true
	}
	C.NSView_setHidden(C.OBJC_PTR(v), h)
}

func NSView_subviews(v native.Handle) native.Handle {
	return native.Handle(C.NSView_subviews(C.OBJC_PTR(v)))
}

func NSView_display(v native.Handle) {
	C.NSView_display((C.OBJC_PTR(v)))
}

func NSView_displayIfNeeded(v native.Handle) {
	C.NSView_displayIfNeeded(C.OBJC_PTR(v))
}

func NSView_needsDisplay(v native.Handle) bool {
	return bool(C.NSView_needsDisplay(C.OBJC_PTR(v)))
}

func NSView_setNeedsDisplay(v native.Handle, needs bool) {
	C.NSView_setNeedsDisplay(C.OBJC_PTR(v), C.bool(needs))
}

func NSView_nextKeyView(v native.Handle) native.Handle {
	return native.Handle(C.NSView_nextKeyView(C.OBJC_PTR(v)))
}

func NSView_setNextKeyView(v, k native.Handle) {
	C.NSView_setNextKeyView(C.OBJC_PTR(v), C.OBJC_PTR(k))
}

func NSView_previousKeyView(v native.Handle) native.Handle {
	return native.Handle(C.NSView_previousKeyView(C.OBJC_PTR(v)))
}


