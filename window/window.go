package window

import (
	"github.com/kevin-yuan/rw"
)

var hm = &rw.WindowHandleManager{}

// New creates a window.
func New() rw.Window {
	w := Alloc()
	rw.Init(w)
	return w
}

// Alloc creates an uninitialized window.
func Alloc() rw.Window {
	w := rw.NewWindowTemplate()
	w.Wrapper().SetHandleManager(hm)
	return w
}
