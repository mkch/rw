package runtime

//#include <stdlib.h>
//#include "runtime.h"
import "C"

import (
	"github.com/kevin-yuan/rw/internal/mem"
)

func RegisterSelector(name string) uintptr {
	return uintptr(C.registerSelector((*C.char)(mem.CStringAutoFree(name))))
}

