package menu

import (
"github.com/mkch/rw/native"
	"github.com/mkch/rw"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/internal/native/darwin/deallochook"
	"github.com/mkch/rw/internal/native/darwin/menu"
)


func createMenuItem(util.Bundle) native.Handle {
	return deallochook.Apply(menu.NewMenuItem())
}

func createSeparator(util.Bundle) native.Handle {
	return deallochook.Apply(menu.NewSeparatorMenuItem())
}

func AllocItem(separator bool) rw.MenuItem {
	var f func(util.Bundle) native.Handle
	if separator {
		f = createSeparator
	} else {
		f = createMenuItem
	}
	return rw.AllocMenuItem(f)
}
