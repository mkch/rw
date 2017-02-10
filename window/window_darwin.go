package window

import (
	"github.com/mkch/rw/internal/native/darwin/deallochook"
	"github.com/mkch/rw/internal/native/darwin/window"
	"github.com/mkch/rw/internal/native/darwin/windowstyle"
	s "github.com/mkch/rw/internal/windowstyle"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
)

func createWindow(util.Bundle) native.Handle {
	style := windowstyle.WindowStyle(&s.WindowStyleFeatures{
		HasBorder:         true,
		HasTitle:          true,
		HasCloseButton:    true,
		HasMinimizeButton: true,
		HasMaximizeButton: true,
		Resizable:         true,
	})
	w := window.NewRWWindow(0, 0, 300, 200, style)
	window.NSWindow_center(w)
	return deallochook.Apply(w)
}
