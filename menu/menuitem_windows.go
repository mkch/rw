package menu

import (
	"github.com/mkch/rw"

)

func AllocItem(separator bool) rw.MenuItem {
	item := rw.NewMenuItemTemplate(separator)
    mgr := &rw.MenuItemHandleManager{}
    mgr.Windows_Id = rw.Windows_NextMenuItemHandle(mgr)
	item.Wrapper().SetHandleManager(mgr)
	return item
}
