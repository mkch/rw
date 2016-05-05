package panel

import (
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/internal/native/darwin/view"
	"github.com/mkch/rw/internal/native/darwin/deallochook"
)

func (m *HandleManager) Create(util.Bundle) native.Handle {
	h := deallochook.Apply(view.NSView_initWithFrame(view.RWFlippedView_alloc(), 0, 0, 0, 0))
	view.RWFlippedView_setAcceptFirstResponder(h, true)
	return h
}