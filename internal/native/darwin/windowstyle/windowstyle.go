package windowstyle

//#include "windowstyle.h"
import "C"

import (
	"github.com/mkch/rw/internal/windowstyle"
)

func WindowStyle(features *windowstyle.WindowStyleFeatures) uint {
	return uint(C.getWindowStyle(&C.WindowStyleFeatures{
		C.bool(features.HasBorder),
		C.bool(features.HasTitle),
		C.bool(features.HasCloseButton),
		C.bool(features.HasMinimizeButton),
		C.bool(features.HasMaximizeButton),
		C.bool(features.Resizable),
	}))
}
