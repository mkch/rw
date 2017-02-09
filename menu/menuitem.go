package menu

import (
	"github.com/mkch/rw"
)


func NewItem() rw.MenuItem {
	item := AllocItem(false)
	rw.Init(item)
	return item
}

func NewSeparator() rw.MenuItem {
	item := AllocItem(true)
	rw.Init(item)
	return item
}
