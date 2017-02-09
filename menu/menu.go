package menu

import (
	"github.com/mkch/rw"
)

func New() rw.Menu {
	m := Alloc()
	rw.Init(m)
	return m
}

func Alloc() rw.Menu {
	m := rw.AllocMenu(createMenu)
	return m
}
