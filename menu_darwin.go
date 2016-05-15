package rw

import (
	"github.com/mkch/rw/internal/native/darwin/app"
	"github.com/mkch/rw/internal/native/darwin/deallochook"
	"github.com/mkch/rw/internal/native/darwin/menu"
	"github.com/mkch/rw/internal/native/darwin/object"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
)

// OSX_SetMainMenu sets the main menu of this app. Available on Mac OS X only.
func OSX_SetMainMenu(mainMenu Menu) {
	a := app.NSApp()
	old := app.NSApplication_mainMenu(a)
	if (mainMenu == nil && old == 0) || (mainMenu != nil && old == mainMenu.Wrapper().Handle()) {
		return
	}
	object.NSObject_retain(old) // The old main menu can be used again. Retain a nil does nothing.

	var m native.Handle
	if mainMenu != nil {
		m = object.NSObject_autorelease(mainMenu.Wrapper().Handle())
	}
	app.NSApplication_setMainMenu(a, m)
}

func OSX_MainManu() (result Menu) {
	m := app.NSApplication_mainMenu(app.NSApp())
	if m == 0 {
		return nil
	}
	if result = defaultObjectTable.Query(m).(Menu); result == nil {
		// result = NewMenuTemplate()
		// result.Wrapper().SetHandleManager(&handlemanagers.ObjcHandl)
		// util.InitWithHandle(result, m)
	}
	return
}

type menuExtra interface {
	setOpener(MenuItem)
}

type menuBase struct {
	objcBase
	items  []MenuItem
	opener MenuItem
}

func (m *menuBase) Window() Window {
	return nil
}

func (m *menuBase) InsertItem(item MenuItem, i int) {
	if m.findItem(item) != -1 {
		panic("Item is already in this menu")
	}
	itemHandle := item.Wrapper().Handle()
	if oldMenu := item.Menu(); oldMenu != nil {
		oldMenu.RemoveItem(item) // Remove from old menu.
	}
	menu.NSMenu_insertItemAtIndex(m.Wrapper().Handle(), object.NSObject_autorelease(itemHandle), i)

	m.items = append(m.items, nil)
	copy(m.items[i+1:], m.items[i:])
	m.items[i] = item
}

func (m *menuBase) AddItem(item MenuItem) {
	if m.findItem(item) != -1 {
		panic("Item is already in this menu")
	}
	if oldMenu := item.Menu(); oldMenu != nil {
		oldMenu.RemoveItem(item) // Remove from old menu.
	}
	itemHandle := item.Wrapper().Handle()
	menu.NSMenu_addItem(m.Wrapper().Handle(), object.NSObject_autorelease(itemHandle))

	m.items = append(m.items, item)
}

func (m *menuBase) RemoveItem(item MenuItem) {
	if i := m.findItem(item); i == -1 {
		panic("Invalid menu item to remove, not in this menu")
	} else {
		menuHandle := m.Wrapper().Handle()
		itemHandle := item.Wrapper().Handle()
		object.NSObject_retain(itemHandle)
		menu.NSMenu_removeItem(menuHandle, itemHandle)
		m.items = append(m.items[:i], m.items[i+1:]...)
	}
}

func (m *menuBase) setOpener(opener MenuItem) {
	m.opener = opener
}

func (m *menuBase) Opener() MenuItem {
	return m.opener
}

type MenuHandleManager struct {
	objcHandleManagerBase
}

func (h MenuHandleManager) Create(util.Bundle) native.Handle {
	// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSMenuItem_Class/#//apple_ref/occ/instp/NSMenuItem/enabled
	// NSMenu.enabled
	// "This property has no effect unless the menu in which the item will be added or is already a part of has been sent setAutoenablesItems:NO"
	handle := menu.NewMenu()
	menu.NSMenu_setAutoenablesItems(handle, false)
	return deallochook.Apply(handle)
}
