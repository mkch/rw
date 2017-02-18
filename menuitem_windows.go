package rw

import (
	"bytes"
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/internal/native/windows/menu"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/util/ustr"
	"strings"
)

const DefaultModifierKey = ControlKey

type menuItemExtra interface {
	setMenu(menu Menu)
	separator() bool
	displayTitle() string
	addAccelerator(Window)
	removeAccelerator(Window)
	Id() uint16
	setId(id uint16)
}

type menuItemBase struct {
	objectBase
	wrapper util.WrapperImpl

	table   *util.ObjectTable
	inTable bool

	sep              bool // Whether is a menu separator
	title            string
	enabled          bool
	checked          bool
	visible          bool
	menu             Menu
	id               uint16 // Non-zero if accelerator is active.
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
	if item.visible && item.menu != nil {
		menu.SetMenuItemInfo(item.menu.Wrapper().Handle(), uint(item.menu.uiPos(item.Self().(MenuItem))), true, &menu.MenuItemInfo{
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
	if item.menu != nil {
		if visible {
			item.visible = visible
			item.menu.addChildItemToUI(item.Self().(MenuItem))
			item.addAccelerator(item.menu.rootWindow())
		} else {
			item.removeAccelerator(item.menu.rootWindow())
			item.menu.removeChildItemFromUI(item.Self().(MenuItem))
			item.visible = visible
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
	if item.visible && item.menu != nil {
		menu.SetMenuItemInfo(item.menu.Wrapper().Handle(), uint(item.menu.uiPos(item.Self().(MenuItem))), true, &menu.MenuItemInfo{
			Mask:  menu.MIIM_STATE,
			State: menuItemStateValue(item),
		})
		item.menu.drawMenuBar()
	}
}

func (item *menuItemBase) Id() uint16 {
	return item.id
}

// item.menu must be non-nil.
func (item *menuItemBase) setId(id uint16) {
	item.checkValidaty()
	if item.id == id {
		return
	}
	item.id = id
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
	}

	if item.visible && item.menu != nil {
		menu.SetMenuItemInfo(item.menu.Wrapper().Handle(), uint(item.menu.uiPos(item.Self().(MenuItem))), true, &menu.MenuItemInfo{
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

func (item *menuItemBase) addAccelerator(win Window) {
	if item.visible && item.hasAccelerator() && !item.acceleratorAdded {
		win.addMenuItemAccelerator(item.Self().(MenuItem))
		item.acceleratorAdded = true
	}
}

func (item *menuItemBase) removeAccelerator(win Window) {
	if item.acceleratorAdded {
		win.removeMenuItemAccelerator(item.id)
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
	if item.visible && item.menu != nil {
		if win := item.menu.rootWindow(); win != nil {
			item.removeAccelerator(win)
		}
	}
	item.accelKey, item.accelModMask = key, mod
	item.syncDisplayTitleToUI()
	if item.visible && item.menu != nil {
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

// HandleManager part:

func (item *menuItemBase) Destroy(handle native.Handle) {
	item.table.Remove(handle)
}

func (item *menuItemBase) Valid(handle native.Handle) bool {
	return handle == 0
}

func (item *menuItemBase) Invalid() native.Handle {
	return ^native.Handle(0)
}

func (item *menuItemBase) Table() *util.ObjectTable {
	return item.table
}

func (item *menuItemBase) Create(b util.Bundle) native.Handle {
	// MenuItem does not wrap a handle.
	// 0 is the dummy valid handle for all MenuItem.
	return 0
}

// ObjectTableStorage part:
// Every MenuItem has its own table in which the only valid object is the MenuItem itself.

func (item *menuItemBase) Get(key native.Handle) (value util.WrapperHolder, exists bool) {
	if item.inTable && key == 0 {
		return item.Self().(MenuItem), true
	} else {
		return nil, false
	}
}

func (item *menuItemBase) Set(key native.Handle, value util.WrapperHolder) {
	if key != 0 || item.Self().(MenuItem) != value {
		panic("Invalid arguments")
	}
	item.inTable = true
}

func (item *menuItemBase) Del(key native.Handle) {
	if key == 0 && item.inTable {
		item.inTable = false
	}
}

func (item *menuItemBase) Len() int {
	return 1
}

func (item *menuItemBase) ForEach(f func(native.Handle, util.WrapperHolder)) {
	f(item.wrapper.Handle(), item.Self().(MenuItem))
}

func initMenuItemBase(item *menuItemBase) *menuItemBase {
	item.table = util.NewObjectTableWithStorage(item)
	item.wrapper.SetHandleManager(item)
	return item
}

func allocMenuItem() MenuItem {
	return initMenuItemBase(&menuItemBase{sep: false, visible: true, enabled: true})
}

func allocSeparatorMenuItem() MenuItem {
	return initMenuItemBase(&menuItemBase{sep: true, visible: true, enabled: true})
}

func AllocMenuItem() MenuItem {
	return allocMenuItem()
}

func AllocSeparatorMenuItem() MenuItem {
	return allocSeparatorMenuItem()
}
