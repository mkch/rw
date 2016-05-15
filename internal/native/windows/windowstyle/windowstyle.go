package windowstyle

//#include "windowstyle.h"
import "C"

import (
	"github.com/mkch/rw/internal/windowstyle"
	"github.com/mkch/rw/native"
)

func WindowStyle(features *windowstyle.WindowStyleFeatures) (style uint, exStyle uint, applyAfterCreate func(native.Handle)) {
	if !features.HasCloseButton {
		applyAfterCreate = func(h native.Handle) {
			C.disableCloseButton(C.HWND(C.Ptr(h)))
		}
	}

	var s, e C.UINT
	C.getWindowStyle(&C.WindowStyleFeatures{
		C.bool(features.HasBorder),
		C.bool(features.HasTitle),
		C.bool(features.HasCloseButton),
		C.bool(features.HasMinimizeButton),
		C.bool(features.HasMaximizeButton),
		C.bool(features.Resizable),
	}, &s, &e)
	style = uint(s)
	exStyle = uint(e)
	return
}
