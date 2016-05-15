package rw

import (
	"bytes"
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/internal/native/windows/acceltable"
	"github.com/mkch/rw/internal/native/windows/menu"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/util/ustr"
	"strings"
	"unicode"
)

// menuItemTable is the ObjectTable which contains MenuItems.
// MenuItem handle(ID) is a different type from Window handle.
var menuItemTable = util.NewObjectTable()

const DefaultModifierKey = ControlKey

type menuItemExtra interface {
	setMenu(menu Menu)
	separator() bool
	displayTitle() string
	addAccelerator(Window)
	removeAccelerator(Window)
}

type menuItemBase struct {
	objectBase
	wrapper util.WrapperImpl

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
	buf := bytes.NewBufferString(util.Windows_ControlTitleWithMnemonic(strings.Replace(item.title, "\t", "    ", -1), item.mnemonic))
	if mod := item.keyModifierString(); mod != "" {
		buf.WriteRune('\t')
		buf.WriteString(mod)
		if item.accelKey != 0 {
			buf.WriteRune('+')
			buf.WriteString(strings.ToUpper(string(item.accelKey)))
		}
	} else {
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
	if prevOpener := submenu.Opener(); prevOpener != nil {
		prevOpener.SetSubmenu(nil)
	}
	item.submenu = submenu

	var submenuHandle native.Handle
	if submenu != nil {
		submenuHandle = submenu.Wrapper().Handle()
		submenu.setOpener(item.Self().(MenuItem))
	}
	if item.menu != nil {
		menu.SetMenuItemInfo(item.menu.Wrapper().Handle(), uint(item.Wrapper().Handle()), false, &menu.MenuItemInfo{
			Mask:    menu.MIIM_SUBMENU,
			SubMenu: submenuHandle,
		})
		item.menu.drawMenuBar()
	}
}

func (item *menuItemBase) hasAccelerator() bool {
	return item.accelKey != 0
}

func (item *menuItemBase) addAccelerator(win Window) {
	if item.hasAccelerator() && !item.acceleratorAdded {
		// http://stackoverflow.com/questions/23592079/why-does-createacceleratortable-not-work-without-fvirtkey
		//https://msdn.microsoft.com/en-us/library/windows/desktop/dd375731(v=vs.85).aspx
		var fVirt byte = acceltable.FVIRTKEY
		k := unicode.ToUpper(item.accelKey) // Virtual key code.
		if item.accelModMask&ControlKey != 0 {
			fVirt |= acceltable.FCONTROL
		}
		if item.accelModMask&AltKey != 0 {
			fVirt |= acceltable.FALT
		}
		if item.accelModMask&ShiftKey != 0 {
			fVirt |= acceltable.FSHIFT
		}
		win.accelTable().Add(fVirt, uint16(k), uint16(item.Wrapper().Handle()))
		item.acceleratorAdded = true
	}
}

func (item *menuItemBase) removeAccelerator(win Window) {
	if item.acceleratorAdded {
		win.accelTable().Remove(uint16(item.Wrapper().Handle()))
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
		util.Release(item)
		if item.submenu != nil {
			item.submenu.Release()
		}
	}
}

func newMenuItemTemplate(separator bool) MenuItem {
	item := &menuItemBase{}
	item.sep = separator
	item.visible = true
	item.enabled = true
	return item
}

type MenuItemHandleManager struct {
	Windows_Id native.Handle
}

func (h *MenuItemHandleManager) Destroy(handle native.Handle) {
	menuItemTable.Remove(handle)
}

func (h *MenuItemHandleManager) Valid(handle native.Handle) bool {
	return handle != 0
}

func (h *MenuItemHandleManager) Table() util.ObjectTable {
	return menuItemTable
}

func (h *MenuItemHandleManager) Create(util.Bundle) native.Handle {
	return h.Windows_Id
}

func Windows_NextMenuItemHandle(m *MenuItemHandleManager) native.Handle {
	// Begins from 100 to skip system IDs, IDOK=1, IDCANCEL=2... IDCONTINUE=11, etc.
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms645505(v=vs.85).aspx
	for h := native.Handle(100); h <= 0xFFFF; h++ {
		if m.Table().Query(h) == nil {
			return h
		}
	}
	panic("Run out of menu item id")
}

type MenuSeparatorHandleManager struct {
	MenuItemHandleManager
}
