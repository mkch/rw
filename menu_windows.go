package rw

import (
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/internal/native/windows/menu"
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/util/ustr"
)

// menuTable is the ObjectTable which contains Menus.
// Menu handle is a different type from Window handle.
var menuTable = util.NewObjectTable()

type menuExtra interface {
	setOpener(MenuItem)
	setWindow(Window)
	uiPos(item MenuItem) int
	uiItem(i int) MenuItem
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
			if mm == nil {
				return nil
			}
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
	if window != nil && m.opener != nil {
		m.opener.SetSubmenu(nil)
	}
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
	m.InsertItem(item, len(m.items))
}

func (m *menuBase) uiItem(i int) MenuItem {
	if i < 0 {
		panic("Invalid index")
	}
	var pos = -1
	for _, item := range m.items {
		if item.Visible() {
			pos++
			if pos == i {
				return item
			}
		}
	}
	panic("No such menu item")
}

func (m *menuBase) uiPos(item MenuItem) int {
	var pos = -1
	for _, child := range m.items {
		if child.Visible() {
			pos++
			if child == item {
				return pos
			}
		}
	}
	panic("No such menu item")
}

func (m *menuBase) addChildItemToUI(item MenuItem) {
	var menuItemInfo *menu.MenuItemInfo
	if item.separator() {
		menuItemInfo = &menu.MenuItemInfo{
			Mask: menu.MIIM_FTYPE,
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
	if id := item.Id(); id != 0 {
		menuItemInfo.Mask |= menu.MIIM_ID
		menuItemInfo.ID = uint(id)
	}
	menu.InsertMenuItem(m.Wrapper().Handle(), uint(m.uiPos(item)), true, menuItemInfo)
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
		if item.Visible() {
			m.removeChildItemFromUI(item)
		}
		item.setMenu(nil)
		m.items = append(m.items[:i], m.items[i+1:]...)
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
	if win := m.rootWindow(); win != nil {
		item.removeAccelerator(win)
		if sub := item.Submenu(); sub != nil {
			sub.removeAccelerators(win)
		}
	}
	// Do not call menu.DeleteMenu here.
	// menu.RemoveMenu does not destroy the sub menu.
	menu.RemoveMenu(m.Wrapper().Handle(), uint(m.uiPos(item)), menu.MF_BYPOSITION)
	m.drawMenuBar()
}

func (m *menuBase) setOpener(item MenuItem) {
	m.opener = item
}

func (m *menuBase) Opener() MenuItem {
	return m.opener
}

func (m *menuBase) afterRegistered(event event.Event, nextHook event.Handler) bool {
	menu.SetMenuInfo(m.wrapper.Handle(), &menu.MenuInfo{
		Mask:  menu.MIM_STYLE,
		Style: menu.MNS_NOTIFYBYPOS,
	})
	return nextHook(event)
}

func (m *menuBase) Release() {
	if m.Wrapper().Valid() {
		// Remove this menu from the window if it has been set to a window.
		if m.window != nil {
			m.window.SetMenu(nil)
		}
		// MenuItem.Release() calls Menu.RemoveItem(), which modifies m.items.
		items := make([]MenuItem, len(m.items))
		copy(items, m.items)
		// Release all items.
		for _, item := range items {
			// Release the item itself, and it's submenu.
			item.Release()
		}
		// Release menu itself.
		util.Release(m)
	}
}

type menuHandleManager func(util.Bundle) native.Handle

func (h menuHandleManager) Destroy(handle native.Handle) {
	menu.DestroyMenu(handle)
	menuTable.Remove(handle)
}

func (h menuHandleManager) Valid(handle native.Handle) bool {
	return handle != h.Invalid()
}

func (h menuHandleManager) Invalid() native.Handle {
	return 0
}

func (h menuHandleManager) Table() *util.ObjectTable {
	return menuTable
}

func (h menuHandleManager) Create(b util.Bundle) native.Handle {
	return h(b)
}

func initMenu(m *menuBase, createHandleFunc func(util.Bundle) native.Handle) *menuBase {
	m.wrapper.SetHandleManager(menuHandleManager(createHandleFunc))
	m.wrapper.AfterRegistered().AddHook(m.afterRegistered)
	return m
}

func allocMenu(createHandleFunc func(util.Bundle) native.Handle) Menu {
	return initMenu(&menuBase{}, createHandleFunc)
}
