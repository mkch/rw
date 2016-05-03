package menu

import (
	"github.com/kevin-yuan/rw"
)

var itemtemHM = &rw.MenuItemHandleManager{}
var separatorHM = &rw.MenuSeparatorHandleManager{}

func newMenuItem(separator bool) rw.MenuItem {
	item := AllocItem(separator)
	rw.Init(item)
	return item
}

func NewItem() rw.MenuItem {
	return newMenuItem(false)
}

func NewSeparator() rw.MenuItem {
	return newMenuItem(true)
}