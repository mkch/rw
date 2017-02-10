package menu

import (
	"github.com/mkch/rw"
)

func NewItem() rw.MenuItem {
	item := AllocItem()
	rw.Init(item)
	return item
}

func NewSeparator() rw.MenuItem {
	item := AllocSeparatorItem()
	rw.Init(item)
	return item
}
