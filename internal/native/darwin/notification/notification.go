package notification

//#include "notification.h"
import "C"


import (
	"github.com/kevin-yuan/rw/native"
)

func NSNotification_object(n native.Handle) native.Handle {
	return native.Handle(C.NSNotification_object(C.OBJC_PTR(n)))
}