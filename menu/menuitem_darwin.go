package menu

import (
	"github.com/kevin-yuan/rw"
	"github.com/kevin-yuan/rw/util"
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