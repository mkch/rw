package panel

import (
	"github.com/mkch/rw"
)

func New() rw.Container {
	p := Alloc()
	rw.Init(p)
	return p
}

func Alloc() rw.Container {
	p := rw.AllocContainer(createPanel)
	return p
}
