package rw

import (
	"fmt"
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/util"
)

// Menu is application or window menu.
// If a Menu is set to a Window(using Window.SetMenu on Windows platform) and that Window is released, the Window automatically releases the Menu, frees
// resources and memory occupied by the Menu. An orphan Menu must be released by the application code(calling Menu.Release). Otherwise, resource/memory
// leak may occur. Releasing a Menu will remove itself from the Window(on Windows platform) it is added, frees itself, and releases it's items and submenus.
type Menu interface {
	Object
	// Items returns all the items added in this menu.
	Items() []MenuItem
	// AddItem adds an item to this menu. Panics if item is already in this menu.
	AddItem(item MenuItem)
	// RemoveItem removes an item from this menu. Panics if item is not an item of this menu.
	RemoveItem(item MenuItem)
	// ItemCount returns the count of items in this menu.
	ItemCount() int
	// Item returns the ith item.
	Item(i int) MenuItem
	// InsertItem inserts an item to this menu at index i.
	InsertItem(item MenuItem, i int)
	// Opener returns the item that opens this menu. Nil is returned if this menu is not a sub menu.
	Opener() MenuItem
	// Window returns the window where this menu is set to. Only makes sense on Windows platform.
	// See Window.SetMenu.
	Window() Window

	util.WrapperHolder

	menuExtra
}

func (m *menuBase) Wrapper() util.Wrapper {
	return &m.wrapper
}

func (m *menuBase) Items() []MenuItem {
	return append(([]MenuItem)(nil), m.items...)
}

func (m *menuBase) Item(i int) MenuItem {
	return m.items[i]
}

func (m *menuBase) findItem(item MenuItem) int {
	for i, mi := range m.items {
		if item == mi {
			return i
		}
	}
	return -1
}

func (m *menuBase) ItemCount() int {
	return len(m.items)
}

func (m *menuBase) String() string {
	if m.Wrapper().Valid() {
		itemCount := m.ItemCount()
		var items string
		if itemCount > 0 {
			items = ": " + m.Item(0).Title()
		}
		if itemCount > 1 {
			items = items + " ..."
		}
		return fmt.Sprintf("Menu %#X, %v items%v", m.Wrapper().Handle(), m.ItemCount(), items)
	} else {
		return "Menu <Invalid>"
	}
}

func (item *menuItemBase) String() string {
	if item.Wrapper().Valid() {
		if item.separator() {
			return "MenuItem Separator"
		} else {
			var sub string
			if item.Submenu() != nil {
				sub = " <has submenu>"
			}
			return fmt.Sprintf("MenuItem %#X, %q%v", item.Wrapper().Handle(), item.Title(), sub)
		}
	} else {
		return "MenuItem <Invalid>"
	}
}

// NewMenuTemplate creates a template of Menu.
// Use NewMenu function in package rw/menu to create objects of Menu.
func NewMenuTemplate() Menu {
	return &menuBase{}
}

type ModifierKey uint

const (
	// Ctrl key.
	ControlKey = ModifierKey(1 << iota)
	// Alt(option) key.
	AltKey
	// Shift key.
	ShiftKey
	// Command(⌘) key on Mac OS X.
	CommandKey
)

type MenuItem interface {
	Object
	// Title returns the title of this menu item.
	Title() string
	// SetTitle sets the title of this menu.
	SetTitle(title string)
	// Visible returns whether this menu item is visible.
	Visible() bool
	// SetVisible sets whether this menu item is visible.
	SetVisible(v bool)
	// Menu returns which menu this menu item is in. Nil is returned if this menu item is not in any menu.
	Menu() Menu
	// Submenu returns the sub menu that this menu item opens.
	Submenu() Menu
	// SetSubmenu sets the sub menu of this menu item.
	SetSubmenu(menu Menu)
	// Enabled returns whether this menu item is enabled.
	Enabled() bool
	// SetEnabled sets whether this menu item is enabled.
	SetEnabled(enabled bool)
	// Checked returns whether this menu item is checked.
	Checked() bool
	// SetChecked sets whether this menu item is checked.
	SetChecked(Checked bool)
	// OnClick returns an event hub where an event is sent before the menu item is clicked.
	OnClick() *event.Hub
	// KeyboardShortcut returns the key combination that trigger this menu item.
	KeyboardShortcut() (mod ModifierKey, key rune)
	// SetKeyboardShortcut sets the key combination that trigger this menu item.
	// Use 0 as key for no key, and 0 modifier for no modifier.
	SetKeyboardShortcut(mod ModifierKey, key rune)
	// Mnemonic returns the access key(mnemonic, underlined single character) of this menu item. 0 for none.
	// Only used on Windows. Always returns 0 on other platforms.
	Mnemonic() rune
	// Mnemonic sets the access key(mnemonic, underlined single character). 0 for none.
	// Only used on Windows. Do nothing on other platforms.
	SetMnemonic(k rune)

	util.WrapperHolder

	menuItemExtra
}

func (item *menuItemBase) Wrapper() util.Wrapper {
	return &item.wrapper
}

// NewMenuItemTempate creates a template of MenuItem.
// Use NewItem function in package rw/menu to create objects of MenuItem.
func NewMenuItemTemplate(separator bool) MenuItem {
	return newMenuItemTemplate(separator)
}
