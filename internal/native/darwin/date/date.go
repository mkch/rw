package date

//#include "date.h"
import "C"

import (
	"github.com/mkch/rw/native"
)

func NSDate_distantFuture() native.Handle {
	return native.Handle(C.NSDate_distantFuture())
}
