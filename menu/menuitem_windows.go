package menu

import (
	"github.com/mkch/rw"
)


func AllocItem() rw.MenuItem {	
	return rw.AllocMenuItem()
}

func AllocSeparatorItem() rw.MenuItem {
	return rw.AllocSeparatorMenuItem()
}
