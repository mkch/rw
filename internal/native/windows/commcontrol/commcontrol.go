package commcontrol

//#include "commcontrol.h"
import "C"

import (
	"unsafe"
    "github.com/mkch/rw/internal/native/windows/nativeutil"
)


var WC_BUTTON = uintptr(unsafe.Pointer(C._WC_BUTTON))
var WC_EDIT = uintptr(unsafe.Pointer(C._WC_EDIT))
var WC_TREEVIEW = uintptr(unsafe.Pointer(C._WC_TREEVIEW))

func Initialize() {
    if C.InitializeCommonControls() == 0 {
        nativeutil.PanicWithLastError()
    }
}

