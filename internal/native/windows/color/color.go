package color

//#include <windows.h>
//#include "color.h"
import "C"

import (
	"image/color"
)

func Uint32(c color.Color) uint32 {
	r, g, b, _ := c.RGBA()
	r8, g8, b8 := C.byte(float32(r)/float32(0xFFFF)*float32(0xFF)), C.byte(float32(g)/float32(0xFFFF)*float32(0xFF)), C.byte(float32(b)/float32(0xFFFF)*float32(0xFF))
	return uint32(C.fRGB(r8, g8, b8))
}

func Color(c uint32) color.Color {
	r := uint8(C.fGetRValue(C.DWORD(c)))
	g := uint8(C.fGetGValue(C.DWORD(c)))
	b := uint8(C.fGetBValue(C.DWORD(c)))
	return color.RGBA{r, g, b, 0xFF}
}
