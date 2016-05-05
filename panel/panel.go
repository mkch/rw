package panel

import (
	"github.com/mkch/rw"
)

type HandleManager struct {
	rw.ControlHandleManagerBase
}

var panelHM = &HandleManager{}

func New() rw.Container {
	p := Alloc()
	rw.Init(p)
	return p
}

func Alloc() rw.Container {
	p := rw.NewContainerTemplate()
	p.Wrapper().SetHandleManager(panelHM)
	return p
}