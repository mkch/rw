package menu

import (
	"github.com/mkch/rw"
	"github.com/mkch/rw/event"
	"runtime"
)

// Builder is a utility for building rw.Menu.
type Builder struct {
	menu        rw.Menu
	itemBuilder *ItemBuilder
	creator     func() rw.Menu
	itemCreator func() rw.MenuItem
}

// NewBuilder call NewBuilderCreators(nil, nil)
func NewBuilder() *Builder {
	return NewBuilderCreators(nil, nil)
}

// NewBuilderCreators creates a Builder. The creator and itemCreator will be used to create new menu and new menu item during building.
// The arguments can be nil to use New and NewMenuItem.
func NewBuilderCreators(creator func() rw.Menu, itemCreator func() rw.MenuItem) *Builder {
	b := &Builder{
		creator:     creator,
		itemCreator: itemCreator,
	}
	runtime.SetFinalizer(b, func(builder *Builder) {
		if builder.menu != nil {
			builder.menu.Release()
		}
	})
	return b
}

func (mb *Builder) ensureMenu() rw.Menu {
	if mb.menu == nil {
		if mb.creator == nil {
			mb.menu = New()
		} else {
			mb.menu = mb.creator()
		}
	}
	return mb.menu
}

func (mb *Builder) newItem() rw.MenuItem {
	if mb.itemCreator == nil {
		return NewItem()
	} else {
		return mb.itemCreator()
	}
}

// AddItem adds a menu item to the menu being built, and returns the Builder itself.
func (mb *Builder) AddItem(item rw.MenuItem) *Builder {
	mb.ensureMenu().AddItem(item)
	return mb
}

// BeginItem adds a menu item with title to the menu being built.
// Use the returned returned ItemBuilder to build the newly added menu item.
func (mb *Builder) BeginItem(title string) *ItemBuilder {
	item := mb.newItem()
	item.SetTitle(title)
	mb.ensureMenu().AddItem(item)
	return &ItemBuilder{
		item:        item,
		creator:     mb.itemCreator,
		menuCreator: mb.creator,
		menuBuilder: mb,
	}
}

// Build returns the menu being built, and sets the current builder to it's initial state.
func (mb *Builder) Build() rw.Menu {
	menu := mb.ensureMenu()
	mb.menu = nil
	mb.itemBuilder = nil
	return menu
}

// End ends the bulding of sub menu and returns the original ItemBuilder whose BeginSubmenu has been called to return this builder.
// If this builder is not returned by ItemBuilder.BeginSubmenu, nil is returned.
// Use the returned builder to continue building the menu item.
func (mb *Builder) End() *ItemBuilder {
	return mb.itemBuilder
}

func (mb *Builder) AddSeparator() *Builder {
	mb.ensureMenu().AddItem(NewSeparator())
	return mb
}

// Builder is a utility for building rw.MenuItem.
type ItemBuilder struct {
	item        rw.MenuItem
	menuBuilder *Builder
	creator     func() rw.MenuItem
	menuCreator func() rw.Menu
}

// NewItemBuilder creates a ItemBuilder. The creator and menureator will be used to create new menu item and new menu during building.
// The arguments can be nil to use NewItem and New.
func NewItemBuilder(creator func() rw.MenuItem, menuCreator func() rw.Menu) *ItemBuilder {
	b := &ItemBuilder{
		creator:     creator,
		menuCreator: menuCreator,
	}
	runtime.SetFinalizer(b, func(builder *ItemBuilder) {
		if builder.item != nil {
			builder.item.Release()
		}
	})
	return b
}

func (mib *ItemBuilder) ensureItem() rw.MenuItem {
	if mib.item == nil {
		if mib.creator == nil {
			mib.item = NewItem()
		} else {
			mib.item = mib.creator()
		}
	}
	return mib.item
}

func (mib *ItemBuilder) newMenu() rw.Menu {
	if mib.menuCreator == nil {
		return New()
	} else {
		return mib.menuCreator()
	}
}

func (mib *ItemBuilder) SetTitle(title string) *ItemBuilder {
	mib.ensureItem().SetTitle(title)
	return mib
}

func (mib *ItemBuilder) SetVisible(visible bool) *ItemBuilder {
	mib.ensureItem().SetVisible(visible)
	return mib
}

func (mib *ItemBuilder) SetEnabled(enabled bool) *ItemBuilder {
	mib.ensureItem().SetEnabled(enabled)
	return mib
}

func (mib *ItemBuilder) SetChecked(checked bool) *ItemBuilder {
	mib.ensureItem().SetChecked(checked)
	return mib
}

func (mib *ItemBuilder) SetKeyboardShortcut(mod rw.ModifierKey, key rune) *ItemBuilder {
	mib.ensureItem().SetKeyboardShortcut(mod, key)
	return mib
}

func (mib *ItemBuilder) SetMnemonic(key rune) *ItemBuilder {
	mib.ensureItem().SetMnemonic(key)
	return mib
}

// SetSubmenu sets the sub menu of the item being built, and returns the ItemBuilder.
func (mib *ItemBuilder) SetSubmenu(submenu rw.Menu) *ItemBuilder {
	mib.ensureItem().SetSubmenu(submenu)
	return mib
}

// BeginSubmenu sets the sub menu of the menu item being bulilt.
// Use the returned Builder to build the newly added sub menu.
func (mib *ItemBuilder) BeginSubmenu() *Builder {
	menu := mib.newMenu()
	mib.ensureItem().SetSubmenu(menu)
	return &Builder{
		menu:        menu,
		itemBuilder: mib,
		creator:     mib.menuCreator,
		itemCreator: mib.creator,
	}
}

func (mib *ItemBuilder) SetOnClickListener(listener interface{}) *ItemBuilder {
	mib.ensureItem().OnClick().SetListener(listener)
	return mib
}

func (mib *ItemBuilder) SetOnClickHandler(handler event.Handler) *ItemBuilder {
	mib.ensureItem().OnClick().SetHandler(handler)
	return mib
}

// Build returns the menu item being built, and sets the current builder to it's initial state.
func (mib *ItemBuilder) Build() rw.MenuItem {
	item := mib.item
	mib.item = nil
	return item
}

// End ends the bulding of menu item and returns the original Builder whose BeginItem has been called to return this builder.
// If this builder is not returned by MeuBuilder.BeginItem, nil is returned
// Use the returned builder to continue building the menu item.
func (mib *ItemBuilder) End() *Builder {
	return mib.menuBuilder
}
