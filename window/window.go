package window

import (
	"github.com/mkch/rw"
)

// New creates a window.
func New() rw.Window {
	w := Alloc()
	rw.Init(w)
	return w
}

// Alloc creates an uninitialized window.
func Alloc() rw.Window {
	w := rw.AllocWindow(createWindow)
	return w
}
