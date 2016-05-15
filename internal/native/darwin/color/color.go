package color

//#include "color.h"
import "C"

import (
	"github.com/mkch/rw/native"
	"image/color"
)

func NSColor_colorWithRGB(c color.Color) native.Handle {
	r, g, b, a := c.RGBA()
	fr, fg, fb, fa := float32(r)/float32(0xFFFF), float32(g)/float32(0xFFFF), float32(b)/float32(0xFFFF), float32(a)/float32(0xFFFF)
	return native.Handle(C.NSColor_colorWithRGB(C.CGFloat(fr), C.CGFloat(fg), C.CGFloat(fb), C.CGFloat(fa)))
}

func NSColor_getRGBA(c native.Handle) color.Color {
	var fr, fg, fb, fa C.CGFloat
	C.NSColor_getRGBA(C.OBJC_PTR(c), &fr, &fg, &fb, &fa)
	r, g, b, a := uint8(0xFF*fr), uint8(0xFF*fg), uint8(0xFF*fb), uint8(0xFF*fa)
	return color.RGBA{r, g, b, a}
}
