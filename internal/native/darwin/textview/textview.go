package textview

//#include "textview.h"
import "C"

import (
	"github.com/mkch/rw/internal/native/darwin/array"
	"github.com/mkch/rw/internal/native/darwin/value"
	"github.com/mkch/rw/native"
)

func NSTextView_setSelectedRange(v native.Handle, start, length uint) {
	C.NSTextView_setSelectedRange(C.OBJC_PTR(v), C.ulong(start), C.ulong(length))
}

func NSTextView_selectedRanges(v native.Handle) (ranges []uint) /*[]uint{location1, length1, location2, length2, ...}*/ {
	rs := C.NSTextView_selectedRanges(C.OBJC_PTR(v))
	count := array.NSArraycount(native.Handle(rs))
	for i := uint(0); i < count; i++ {
		r := array.NSArrayobjectAtIndex(native.Handle(rs), i)
		loc, length := value.NSValue_rangeValue(r)
		ranges = append(append(ranges, loc), length)
	}
	return
}
