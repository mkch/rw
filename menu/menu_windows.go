package menu

import (
	"github.com/mkch/rw/internal/native/windows/menu"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
)

func createMenu(util.Bundle) native.Handle {
	return menu.CreateMenu()
}
