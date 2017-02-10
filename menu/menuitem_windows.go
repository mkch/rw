package menu

import (
	"github.com/mkch/rw"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
)

type handleManager struct {
	util.HandleManager
}

func (m handleManager) Create(util.Bundle) native.Handle {
	// Begins from 100 to skip system IDs, IDOK=1, IDCANCEL=2... IDCONTINUE=11, etc.
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms645505(v=vs.85).aspx
	for h := native.Handle(100); h <= 0xFFFF; h++ {
		if m.Table().Query(h) == nil {
			return h
		}
	}
	panic("Run out of menu item id")
}

func AllocItem() rw.MenuItem {
	item := rw.AllocMenuItem(func(util.Bundle) native.Handle { return 0 /*will never be called*/ })
	item.Wrapper().SetHandleManager(handleManager{item.Wrapper().HandleManager()})
	return item
}

func AllocSeparatorItem() rw.MenuItem {
	item := rw.AllocSeparatorMenuItem(func(util.Bundle) native.Handle { return 0 /*will never be called*/ })
	item.Wrapper().SetHandleManager(handleManager{item.Wrapper().HandleManager()})
	return item
}
