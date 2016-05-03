package value

//#include "value.h"
import "C"

import (
	"github.com/kevin-yuan/rw/native"
)


func NSValue_rangeValue(v native.Handle) (location, length uint) {
	var loc, l C.ulong
	C.NSValue_rangeValue(C.OBJC_PTR(v), &loc, &l)
	return uint(loc), uint(l)
}
