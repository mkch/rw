package menu

import (
	"github.com/mkch/rw"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
)

func nextMenuItemId(m util.HandleManager) native.Handle {
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
	var f func(util.Bundle) native.Handle
	item := rw.AllocMenuItem(func(b util.Bundle) native.Handle { return f(b) })
	f = func(util.Bundle) native.Handle {
		return nextMenuItemId(item.Wrapper().HandleManager())
	}
	return item
}

func AllocSeparatorItem() rw.MenuItem {
	var f func(util.Bundle) native.Handle
	item := rw.AllocSeparatorMenuItem(func(b util.Bundle) native.Handle { return f(b) })
	f = func(util.Bundle) native.Handle {
		return nextMenuItemId(item.Wrapper().HandleManager())
	}
	return item
}
