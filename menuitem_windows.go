package rw

import (
	"bytes"
	"fmt"
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/internal/native/windows/acceltable"
	"github.com/mkch/rw/internal/native/windows/menu"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/util/ustr"
	"strings"
	"unicode"
)

// defaultMenuItemTable is the ObjectTable which contains MenuItems that do not attached to a window.
// MenuItem handle(ID) is a different type from Window handle.
var defaultMenuItemTable = util.NewObjectTable()

const DefaultModifierKey = ControlKey

type menuItemExtra interface {
	setMenu(menu Menu)
	separator() bool
	displayTitle() string
	addAccelerator(Window)
	removeAccelerator(Window)
	setTable(table *util.ObjectTable)
	setTableWithIndex(table *util.ObjectTable, index int)
}

type menuItemBase struct {
	objectBase
	wrapper       util.WrapperImpl
	handleManager menuItemHandleManager

	sep              bool // Whether is a menu separator
	title            string
	enabled          bool
	checked          bool
	visible          bool
	menu             Menu
	submenu          Menu
	onClick          event.Hub
	accelKey         rune        // The basic accelerator key.
	accelModMask     ModifierKey // The accelerator key modifiers.
	mnemonic         rune        // Underlined letter.
	acceleratorAdded bool
}

func (item *menuItemBase) keyModifierString() string {
	var keys = make([]string, 0, 3)
	if item.accelModMask&ControlKey != 0 {
		keys = append(keys, "Ctrl")
	}
	if item.accelModMask&AltKey != 0 {
		keys = append(keys, "Alt")
	}
	if item.accelModMask&ShiftKey != 0 {
		keys = append(keys, "Shift")
	}
	return strings.Join(keys, "+")
}

func (item *menuItemBase) displayTitle() string {
	buf := &bytes.Buffer{}
	if Debug {
		fmt.Fprintf(buf, "[0x%x] ", item.Wrapper().Handle())
	}
	buf.WriteString(util.Windows_ControlTitleWithMnemonic(strings.Replace(item.title, "\t", "    ", -1), item.mnemonic))
	if mod := item.keyModifierString(); mod != "" {
		buf.WriteRune('\t')
		buf.WriteString(mod)
		if item.accelKey != 0 {
			buf.WriteRune('+')
			buf.WriteString(strings.ToUpper(string(item.accelKey)))
		}
	} else if item.accelKey != 0 {
		buf.WriteRune('\t')
		buf.WriteString(strings.ToUpper(string(item.accelKey)))
	}
	return buf.String()
}

func (item *menuItemBase) separator() bool {
	return item.sep
}

func (item *menuItemBase) checkValidaty() {
	item.Wrapper().Handle() // Check validaty
}

func (item *menuItemBase) Title() string {
	item.checkValidaty()
	return item.title
}

func (item *menuItemBase) SetTitle(title string) {
	item.checkValidaty()
	if item.title == title {
		return
	}
	item.title = title
	item.syncDisplayTitleToUI()
}

func (item *menuItemBase) syncDisplayTitleToUI() {
	if item.menu != nil {
		menu.SetMenuItemInfo(item.menu.Wrapper().Handle(), uint(item.Wrapper().Handle()), false, &menu.MenuItemInfo{
			Mask:     menu.MIIM_STRING,
			TypeData: uintptr(ustr.CStringUtf16(item.displayTitle())),
		})
		item.menu.drawMenuBar()
	}
}

func (item *menuItemBase) OnClick() *event.Hub {
	return &item.onClick
}

func (item *menuItemBase) Visible() bool {
	item.checkValidaty()
	return item.visible
}

func (item *menuItemBase) SetVisible(visible bool) {
	item.checkValidaty()
	if item.visible == visible {
		return
	}
	item.visible = visible
	if item.menu != nil {
		if item.visible {
			item.menu.addChildItemToUI(item.Self().(MenuItem))
		} else {
			item.menu.removeChildItemFromUI(item.Self().(MenuItem))
		}
	}
}

func (item *menuItemBase) Enabled() bool {
	item.checkValidaty()
	return item.enabled
}

func (item *menuItemBase) SetEnabled(enabled bool) {
	item.checkValidaty()
	if item.enabled == enabled {
		return
	}
	item.enabled = enabled
	item.syncStateToUI()
}

func (item *menuItemBase) Checked() bool {
	item.checkValidaty()
	return item.checked
}

func (item *menuItemBase) SetChecked(checked bool) {
	item.checkValidaty()
	if item.checked == checked {
		return
	}
	item.checked = checked
	item.syncStateToUI()
}

func (item *menuItemBase) syncStateToUI() {
	if item.menu != nil {
		menu.SetMenuItemInfo(item.menu.Wrapper().Handle(), uint(item.Wrapper().Handle()), false, &menu.MenuItemInfo{
			Mask:  menu.MIIM_STATE,
			State: menuItemStateValue(item),
		})
		item.menu.drawMenuBar()
	}
}

func setMenuItemId(menuHandle native.Handle, pos int, id native.Handle) {
	if pos < 0 {
		panic("invalid pos")
	}
	menu.SetMenuItemInfo(menuHandle, uint(pos), true, &menu.MenuItemInfo{
		Mask: menu.MIIM_ID,
		ID:   uint(id),
	})
}

func (item *menuItemBase) Menu() Menu {
	item.checkValidaty()
	return item.menu
}

func (item *menuItemBase) setMenu(menu Menu) {
	item.checkValidaty()
	item.menu = menu
}

func (item *menuItemBase) Submenu() Menu {
	item.checkValidaty()
	return item.submenu
}

func (item *menuItemBase) SetSubmenu(submenu Menu) {
	item.checkValidaty()
	if item.submenu == submenu {
		return
	}

	if item.submenu != nil {
		item.submenu.setOpener(nil)
		item.submenu.setItemsTable(defaultMenuItemTable)
	}

	var submenuHandle native.Handle
	if submenu != nil {
		if prevOpener := submenu.Opener(); prevOpener != nil {
			prevOpener.SetSubmenu(nil)
		} else if window := submenu.Window(); window != nil {
			window.SetMenu(nil)
		}

		submenuHandle = submenu.Wrapper().Handle()
		submenu.setOpener(item.Self().(MenuItem))
		submenu.setItemsTable(item.handleManager.Table())
	}

	if item.menu != nil {
		menu.SetMenuItemInfo(item.menu.Wrapper().Handle(), uint(item.Wrapper().Handle()), false, &menu.MenuItemInfo{
			Mask:    menu.MIIM_SUBMENU,
			SubMenu: submenuHandle,
		})
		item.menu.drawMenuBar()
	}

	item.submenu = submenu
}

func (item *menuItemBase) hasAccelerator() bool {
	return item.accelKey != 0
}

func addAcceleratorToWindow(win Window, accelKey rune, accelModMask ModifierKey, id uint16) {
	// http://stackoverflow.com/questions/23592079/why-does-createacceleratortable-not-work-without-fvirtkey
	//https://msdn.microsoft.com/en-us/library/windows/desktop/dd375731(v=vs.85).aspx
	var fVirt byte = acceltable.FVIRTKEY
	k := unicode.ToUpper(accelKey) // Virtual key code.
	if accelModMask&ControlKey != 0 {
		fVirt |= acceltable.FCONTROL
	}
	if accelModMask&AltKey != 0 {
		fVirt |= acceltable.FALT
	}
	if accelModMask&ShiftKey != 0 {
		fVirt |= acceltable.FSHIFT
	}
	win.accelTable().Add(fVirt, uint16(k), id)
}

func (item *menuItemBase) addAccelerator(win Window) {
	if item.hasAccelerator() && !item.acceleratorAdded {
		addAcceleratorToWindow(win, item.accelKey, item.accelModMask, uint16(item.Wrapper().Handle()))
		item.acceleratorAdded = true
	}
}

func removeAcceleratorFromWindow(win Window, id uint16) {
	win.accelTable().Remove(id)
}

func (item *menuItemBase) removeAccelerator(win Window) {
	if item.acceleratorAdded {
		removeAcceleratorFromWindow(win, uint16(item.Wrapper().Handle()))
		item.acceleratorAdded = false
	}
}

func (item *menuItemBase) KeyboardShortcut() (mod ModifierKey, key rune) {
	return item.accelModMask, item.accelKey
}

func (item *menuItemBase) SetKeyboardShortcut(mod ModifierKey, key rune) {
	if key == 0 && mod != 0 {
		panic("Shortcut key must not be 0")
	}
	if item.accelKey == key && item.accelModMask == mod {
		return
	}
	if item.menu != nil {
		if win := item.menu.rootWindow(); win != nil {
			item.removeAccelerator(win)
		}
	}
	item.accelKey, item.accelModMask = key, mod
	item.syncDisplayTitleToUI()
	if item.menu != nil {
		if win := item.menu.rootWindow(); win != nil {
			item.addAccelerator(win)
		}
	}
}

func (item *menuItemBase) Mnemonic() rune {
	return item.mnemonic
}

func (item *menuItemBase) SetMnemonic(k rune) {
	if k == item.mnemonic {
		return
	}
	item.mnemonic = k
	item.syncDisplayTitleToUI()
}

func (item *menuItemBase) Release() {
	if item.Wrapper().Valid() {
		if item.menu != nil && item.menu.Wrapper().Valid() {
			item.menu.RemoveItem(item.Self().(MenuItem))
		}
		if item.submenu != nil {
			item.submenu.Release()
		}
		util.Release(item)
	}
}

// setTable sets the table of this item. This item will be recreated.
func (item *menuItemBase) setTable(table *util.ObjectTable) {
	if item.menu != nil {
		item.setTableWithIndex(table, item.menu.findItem(item.Self().(MenuItem)))
	} else {
		if table != item.handleManager.Table() {
			util.Recreate(item, util.Bundle{menuItem_KEY_TABLE: table, menuItem_KEY_OLD_ID: item.Wrapper().Handle()})
		}
		if Debug {
			// For displayed menu item id.
			// Post: Changing menu item string by id in the same message cicle of changing menu item id
			// fails for menu bar items.
			Post(item.syncDisplayTitleToUI)
		}
	}
}

// setTableWithIndex sets the table of this item. This item must be in a menu.
// index is the pos of this item in the menu. We can calculate this by iterating
// the items of the menu, however if the pos is known calling this function is more
// efficent.
func (item *menuItemBase) setTableWithIndex(table *util.ObjectTable, index int) {
	if index < 0 {
		panic("Invalid index")
	}
	if table != item.handleManager.Table() {
		util.Recreate(item, util.Bundle{menuItem_KEY_TABLE: table, menuItem_KEY_OLD_ID: item.Wrapper().Handle(), menuItem_KEY_INDEX: index})
	}
	if Debug {
		// For displayed menu item id.
		// Post: Changing menu item string by id in the same message cicle of changing menu item id
		// fails for menu bar items.
		Post(item.syncDisplayTitleToUI)
	}
}

const menuItem_KEY_TABLE = "rw:menuitem-table"
const menuItem_KEY_INDEX = "rw:menuitem-index"
const menuItem_KEY_OLD_ID = "rw:menuitem-old_id"

func (item *menuItemBase) afterDestroyed(event event.Event, nextHook event.Handler) bool {
	we := event.(*util.WrapperEvent)
	if we.Recreating() {
		if b := we.Bundle(); b != nil {
			if table, ok := b[menuItem_KEY_TABLE].(*util.ObjectTable); ok {
				item.handleManager.table = table
			}
		}
	}
	return nextHook(event)
}

func (item *menuItemBase) afterRegistered(event event.Event, nextHook event.Handler) bool {
	we := event.(*util.WrapperEvent)
	if we.Recreating() {
		if b := we.Bundle(); b != nil {
			if index, ok := b[menuItem_KEY_INDEX].(int); ok {
				id := item.Wrapper().Handle()
				if oldId, ok := b[menuItem_KEY_OLD_ID].(native.Handle); !ok || oldId != id {
					// item.menu can't be nil if oldId is present.
					setMenuItemId(item.menu.Wrapper().Handle(), index, id)
					if item.hasAccelerator() && item.acceleratorAdded {
						// item.menu.rootWindow() can't be nil, because the accelerator has been added.
						win := item.menu.rootWindow()
						removeAcceleratorFromWindow(win, uint16(oldId))
						addAcceleratorToWindow(win, item.accelKey, item.accelModMask, uint16(id))
					}
				}
			}
		}
	}
	return nextHook(event)
}

type menuItemHandleManager struct {
	create func(util.Bundle) native.Handle
	table  *util.ObjectTable
}

func (h *menuItemHandleManager) Destroy(handle native.Handle) {
	h.table.Remove(handle)
}

func (h *menuItemHandleManager) Valid(handle native.Handle) bool {
	return handle != h.Invalid()
}

func (h *menuItemHandleManager) Invalid() native.Handle {
	return 0
}

func (h *menuItemHandleManager) Table() *util.ObjectTable {
	return h.table
}

func (h *menuItemHandleManager) Create(b util.Bundle) native.Handle {
	return h.create(b)
}

func initMenuItemBase(item *menuItemBase, createHandleFunc func(util.Bundle) native.Handle) *menuItemBase {
	item.wrapper.AfterDestroyed().AddHook(item.afterDestroyed)
	item.wrapper.AfterRegistered().AddHook(item.afterRegistered)
	item.handleManager.table = defaultMenuItemTable
	item.handleManager.create = createHandleFunc
	item.wrapper.SetHandleManager(&item.handleManager)
	return item
}

func allocMenuItemImp(createHandleFunc func(util.Bundle) native.Handle, sep bool) MenuItem {
	return initMenuItemBase(&menuItemBase{sep: sep, visible: true, enabled: true}, createHandleFunc)
}

func allocMenuItem(createHandleFunc func(util.Bundle) native.Handle) MenuItem {
	return allocMenuItemImp(createHandleFunc, false)
}

func allocSeparatorMenuItem(createHandleFunc func(util.Bundle) native.Handle) MenuItem {
	return allocMenuItemImp(createHandleFunc, true)
}
