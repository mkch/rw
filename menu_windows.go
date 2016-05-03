package rw

import (
	"github.com/kevin-yuan/rw/internal/native/windows/menu"
	"github.com/kevin-yuan/rw/util/ustr"
	"github.com/kevin-yuan/rw/internal/native/windows/window"
	"github.com/kevin-yuan/rw/util"
	"github.com/kevin-yuan/rw/native"
)

// menuTable is the ObjectTable which contains Menus.
// Menu handle is a different type from Window handle.
var menuTable = util.NewObjectTable()

type menuExtra interface {
	setOpener(MenuItem)
	setWindow(Window)
	addChildItemToUI(MenuItem)
	removeChildItemFromUI(MenuItem)
	drawMenuBar()
	rootWindow() Window
	addAccelerators(Window)
	removeAccelerators(Window)
}

type menuBase struct {
	objectBase
	wrapper       util.WrapperImpl
	items         []MenuItem
	opener        MenuItem
	syncScheduled bool
	window        Window
}

// menuItemStateValue returns the current state value, used as menu.MenuItemInfo.State, of the menu item.
func menuItemStateValue(item MenuItem) uint {
	state := uint(0)
	if item.Enabled() {
		state |= menu.MFS_ENABLED
	} else {
		state |= menu.MFS_DISABLED
	}

	if item.Checked() {
		state |= menu.MFS_CHECKED
	} else {
		state |= menu.MFS_UNCHECKED
	}

	return state
}

func (m *menuBase) rootWindow() Window {
	var mm Menu = m
	for {
		if win := mm.Window(); win != nil {
			return win
		} else if opener := mm.Opener(); opener != nil {
			mm = opener.Menu()
		} else {
			return nil
		}
	}
	return nil
}

func (m *menuBase) Window() Window {
	return m.window
}

func (m *menuBase) setWindow(window Window) {
	m.window = window
}

func (m *menuBase) InsertItem(item MenuItem, i int) {
	if pos := m.findItem(item); pos != -1 {
		panic("Item is already in this menu")
	} else {
		if prevMenu := item.Menu(); prevMenu != nil {
			prevMenu.RemoveItem(item)
		}
		item.setMenu(m.Self().(Menu))

		m.items = append(m.items, nil)
		copy(m.items[i+1:], m.items[i:])
		m.items[i] = item

		if item.Visible() {
			m.addChildItemToUI(item)
		}
	}
}

func (m *menuBase) drawMenuBar() {
	if m.window != nil {
		window.DrawMenuBar(m.window.Wrapper().Handle())
	}
}

func (m *menuBase) AddItem(item MenuItem) {
	if prevMenu := item.Menu(); prevMenu == m.Self() {
		panic("Item is already in this menu")
	} else if prevMenu != nil {
		prevMenu.RemoveItem(item)
	}
	item.setMenu(m.Self().(Menu))
	m.items = append(m.items, item)
	if item.Visible() {
		m.addChildItemToUI(item)
	}
}

func (m *menuBase) displayPos(item MenuItem) int {
	var pos int
	for _, child := range m.items {
		if child == item {
			return pos
		}
		if child.Visible() {
			pos++
		}
	}
	panic("No such menu item")
}

func (m *menuBase) addChildItemToUI(item MenuItem) {
	var menuItemInfo *menu.MenuItemInfo
	if item.separator() {
		menuItemInfo = &menu.MenuItemInfo{
			Mask: menu.MIIM_FTYPE | menu.MIIM_ID,
			Type: menu.MFT_SEPARATOR,
		}
	} else {
		menuItemInfo = &menu.MenuItemInfo{
			Mask:     menu.MIIM_STRING | menu.MIIM_ID | menu.MIIM_STATE,
			TypeData: uintptr(ustr.CStringUtf16(item.displayTitle())),
			State:    menuItemStateValue(item),
		}
		if submenu := item.Submenu(); submenu != nil {
			menuItemInfo.Mask |= menu.MIIM_SUBMENU
			menuItemInfo.SubMenu = submenu.Wrapper().Handle()
		}
	}
	menuItemInfo.ID = uint(item.Wrapper().Handle())
	menu.InsertMenuItem(m.Wrapper().Handle(), uint(m.displayPos(item)), true, menuItemInfo)
	if win := m.rootWindow(); win != nil {
		item.addAccelerator(win)
		if sub := item.Submenu(); sub != nil {
			sub.addAccelerators(win)
		}
	}
	m.drawMenuBar()
}

func (m *menuBase) RemoveItem(item MenuItem) {
	if i := m.findItem(item); i == -1 {
		panic("Invalid menu item to remove, not in this menu")
	} else {
		m.items = append(m.items[:i], m.items[i+1:]...)
		m.removeChildItemFromUI(item)
		item.setMenu(nil)
	}
}

// Remove all accelerators in this menu from the root window(if any), including all submenus.
func (m *menuBase) removeAccelerators(win Window) {
	for _, item := range m.items {
		if sub := item.Submenu(); sub != nil {
			sub.removeAccelerators(win)
		}
		item.removeAccelerator(win)
	}
}

// Add all accelerators in this menu to win, including all submenus.
func (m *menuBase) addAccelerators(win Window) {
	for _, item := range m.items {
		if sub := item.Submenu(); sub != nil {
			sub.addAccelerators(win)
		}
		item.addAccelerator(win)
	}
}

func (m *menuBase) removeChildItemFromUI(item MenuItem) {
	// Do not call menu.DeleteMenu here.
	// menu.RemoveMenu does not destroy the sub menu.
	menu.RemoveMenu(m.Wrapper().Handle(), uint(item.Wrapper().Handle()), menu.MF_BYCOMMAND)
	if win := m.rootWindow(); win != nil {
		item.removeAccelerator(win)
		if sub := item.Submenu(); sub != nil {
			sub.removeAccelerators(win)
		}
	}
	m.drawMenuBar()
}

func (m *menuBase) setOpener(item MenuItem) {
	m.opener = item
}

func (m *menuBase) Opener() MenuItem {
	return m.opener
}

func (m *menuBase) Release() {
	if m.Wrapper().Valid() {
		// Remove this menu from the window if it has been set to a window.
		if m.window != nil {
			m.window.SetMenu(nil)
		}
		// Release all items.
		for _, item := range m.items {
			// Release submenus recursively.
			if submenu := item.Submenu(); submenu != nil {
				submenu.Release()
			}
			// Release the item itself.
			item.Release()
		}
		// Release menu itself.
		util.Release(m)
	}
}

type MenuHandleManager struct{}

func (h MenuHandleManager) Destroy(handle native.Handle) {
	menu.DestroyMenu(handle)
	menuTable.Remove(handle)
}

func (h MenuHandleManager) Valid(handle native.Handle) bool {
	return handle != 0
}

func (h MenuHandleManager) Table() util.ObjectTable {
	return menuTable
}

func (h MenuHandleManager) Create(util.Bundle) native.Handle {
	handle := menu.CreateMenu()
	return handle
}