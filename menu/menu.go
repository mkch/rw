package menu

import (
	"github.com/mkch/rw"
)

var menuHM = &rw.MenuHandleManager{}

func New() rw.Menu {
	m := Alloc()
	rw.Init(m)
	return m
}

func Alloc() rw.Menu {
	m := rw.NewMenuTemplate()
	m.Wrapper().SetHandleManager(menuHM)
	return m
}
