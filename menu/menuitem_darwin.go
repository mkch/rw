package menu

import (
	"github.com/mkch/rw"
	"github.com/mkch/rw/util"
)

func AllocItem(separator bool) rw.MenuItem {
	item := rw.NewMenuItemTemplate(separator)
	var mgr util.HandleManager
	if separator {
		mgr = separatorHM
	} else {
		mgr = itemtemHM
	}
	item.Wrapper().SetHandleManager(mgr)
	return item
}