package menu

import(
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/internal/native/darwin/menu"
	"github.com/mkch/rw/internal/native/darwin/deallochook"
)

func createMenu(util.Bundle) native.Handle {
	// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSMenuItem_Class/#//apple_ref/occ/instp/NSMenuItem/enabled
	// NSMenu.enabled
	// "This property has no effect unless the menu in which the item will be added or is already a part of has been sent setAutoenablesItems:NO"
	handle := menu.NewMenu()
	menu.NSMenu_setAutoenablesItems(handle, false)
	return deallochook.Apply(handle)
}
