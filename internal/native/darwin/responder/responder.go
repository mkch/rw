package responder

//#include "responder.h"
import "C"

import (
	"github.com/mkch/rw/native"
)

func NSResponsder_acceptsFirstResponder(r native.Handle) bool {
	return bool(C.NSResponder_acceptsFirstResponder(C.OBJC_PTR(r)))
}
