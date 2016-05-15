package runtime

//#include <stdlib.h>
//#include "runtime.h"
import "C"

import (
	"github.com/mkch/rw/util/ustr"
)

func RegisterSelector(name string) uintptr {
	return uintptr(C.registerSelector((*C.char)(ustr.CStringUtf8(name))))
}
