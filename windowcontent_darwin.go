package rw

import (
	"github.com/kevin-yuan/rw/internal/native/darwin/deallochook"
	"github.com/kevin-yuan/rw/internal/native/darwin/view"
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/util"
)

type winContentHandleManager struct {
	objcHandleManagerBase
}

func (m *winContentHandleManager) Create(util.Bundle) native.Handle {
	return deallochook.Apply(view.NSView_initWithFrame(view.RWFlippedView_alloc(), 0, 0, 0, 0))
}

var winContentHM = &winContentHandleManager{}

type winContent struct {
	Container
}

func newWindowContent() Container {
	c := &winContent{NewContainerTemplate()}
	c.Wrapper().SetHandleManager(winContentHM)
	Init(c)
	return c
}
