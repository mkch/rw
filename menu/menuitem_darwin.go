package menu

import (
	"github.com/mkch/rw"
	"github.com/mkch/rw/internal/native/darwin/deallochook"
	"github.com/mkch/rw/internal/native/darwin/menu"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
)

func createMenuItem(util.Bundle) native.Handle {
	return deallochook.Apply(menu.NewMenuItem())
}

func createSeparator(util.Bundle) native.Handle {
	return deallochook.Apply(menu.NewSeparatorMenuItem())
}

func AllocItem() rw.MenuItem {
	return rw.AllocMenuItem(createMenuItem)
}

func AllocSeparatorItem() rw.MenuItem {
	return rw.AllocSeparatorMenuItem(createSeparator)
}
