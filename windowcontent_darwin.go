package rw

import (
	"github.com/mkch/rw/internal/native/darwin/deallochook"
	"github.com/mkch/rw/internal/native/darwin/view"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
)

func createWinContent(util.Bundle) native.Handle {
	return deallochook.Apply(view.NSView_initWithFrame(view.RWFlippedView_alloc(), 0, 0, 0, 0))
}

type winContent struct {
	Container
}

func newWindowContent() Container {
	c := &winContent{AllocContainer(createWinContent)}
	Init(c)
	return c
}
