package rw

import (
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/internal/native/darwin/deallochook"
	"github.com/mkch/rw/internal/native/darwin/dynamicinvocation"
	"github.com/mkch/rw/internal/native/darwin/menu"
	"github.com/mkch/rw/internal/native/darwin/object"
	"github.com/mkch/rw/internal/native/darwin/runtime"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"unicode"
	"unicode/utf8"
)

const DefaultModifierKey = CommandKey

type menuItemExtra interface {
	separator() bool
}

type menuItemBase struct {
	objcBase
	onClick event.Hub
}

func (item *menuItemBase) OnClick() *event.Hub {
	item.ensureTargetAction()
	return &item.onClick
}

func (item *menuItemBase) ensureTargetAction() {
	handle := item.Wrapper().Handle()
	if object.Target(handle) != 0 {
		return
	}
	d := dynamicinvocation.RWDynamicInvocation_initWithMethodsCallback([]string{
		"targetAction:", "v@:@",
	}, func(selector string, args native.Handle) {
		switch selector {
		case "targetAction:":
			item.onClick.Send(&simpleEvent{sender: item.Self()})
		}
	})

	object.SetTargetRetain(handle, d)
	object.NSObject_release(d)
	menu.NSMenuItem_setAction(handle, runtime.RegisterSelector("targetAction:"))
}

func (item *menuItemBase) separator() bool {
	return menu.NSMenuItem_isSeparatorItem(item.Wrapper().Handle())
}

func (item *menuItemBase) Title() string {
	return menu.NSMenuItem_title(item.Wrapper().Handle())
}

func (item *menuItemBase) SetTitle(title string) {
	menu.NSMenuItem_setTitle(item.Wrapper().Handle(), title)
	if sub := menu.NSMenuItem_submenu(item.Wrapper().Handle()); sub != 0 {
		menu.NSMenu_setTitle(sub, title)
	}
}

func (item *menuItemBase) Visible() bool {
	return !menu.NSMenuItem_isHidden(item.Wrapper().Handle())
}

func (item *menuItemBase) SetVisible(visible bool) {
	menu.NSMenuItem_setHidden(item.Wrapper().Handle(), !visible)
}

func (item *menuItemBase) Enabled() bool {
	return menu.NSMenuItem_isEnabled(item.Wrapper().Handle())
}

func (item *menuItemBase) SetEnabled(enabled bool) {
	menu.NSMenuItem_setEnabled(item.Wrapper().Handle(), enabled)
}

func (item *menuItemBase) Checked() bool {
	return menu.NSMenuItem_state(item.Wrapper().Handle()) == menu.NSOnState
}

func (item *menuItemBase) SetChecked(checked bool) {
	state := menu.NSOffState
	if checked {
		state = menu.NSOnState
	}
	menu.NSMenuItem_setState(item.Wrapper().Handle(), state)
}

func (item *menuItemBase) Menu() Menu {
	if m := menu.NSMenuItem_menu(item.Wrapper().Handle()); m != 0 {
		return defaultObjectTable.Query(m).(Menu)
	}
	return nil
}

func (item *menuItemBase) Submenu() Menu {
	if sub := menu.NSMenuItem_submenu(item.Wrapper().Handle()); sub != 0 {
		return defaultObjectTable.Query(sub).(Menu)
	}
	return nil
}

func (item *menuItemBase) SetSubmenu(sub Menu) {
	oldSubmenu := item.Submenu()
	if oldSubmenu == sub {
		return // Nothing changed.
	}

	var itemHandle = item.Wrapper().Handle()
	var subHandle native.Handle

	if sub != nil {
		subHandle = sub.Wrapper().Handle()
		if opener := sub.Opener(); opener != nil {
			object.NSObject_retain(sub.Wrapper().Handle())
			opener.SetSubmenu(nil)
		}

		menu.NSMenu_setTitle(subHandle, menu.NSMenuItem_title(itemHandle))
		sub.setOpener(item.Self().(MenuItem))
	}

	if oldSubmenu != nil {
		object.NSObject_retain(oldSubmenu.Wrapper().Handle())
		oldSubmenu.setOpener(nil)
	}

	menu.NSMenuItem_setSubmenu(itemHandle, subHandle)
	object.NSObject_release(subHandle)
}

func (item *menuItemBase) KeyboardShortcut() (mod ModifierKey, key rune) {
	key, _ = utf8.DecodeRuneInString(menu.NSMenuItem_keyEquivalent(item.Wrapper().Handle()))
	if key == utf8.RuneError {
		key = 0
	}
	// We use lower case letter for key equivalent in Mac OS X, but upper case letter in rw API.
	key = unicode.ToUpper(key)
	m := menu.NSMenuItem_keyEquivalentModifierMask(item.Wrapper().Handle())
	if m&menu.NSShiftKeyMask != 0 {
		mod |= ShiftKey
	}
	if m&menu.NSAlternateKeyMask != 0 {
		mod |= AltKey
	}
	if m&menu.NSCommandKeyMask != 0 {
		mod |= CommandKey
	}
	if m&menu.NSControlKeyMask != 0 {
		mod |= ControlKey
	}
	return
}

func (item *menuItemBase) SetKeyboardShortcut(mod ModifierKey, key rune) {
	if key == 0 && mod != 0 {
		panic("Shortcut key must not be 0")
	}
	key = unicode.ToLower(key)
	var m menu.NSKeyEquivalentModifierMask
	if mod&ShiftKey != 0 {
		m |= menu.NSShiftKeyMask
	}
	if mod&AltKey != 0 {
		m |= menu.NSAlternateKeyMask
	}
	if mod&CommandKey != 0 {
		m |= menu.NSCommandKeyMask
	}
	if mod&ControlKey != 0 {
		m |= menu.NSControlKeyMask
	}

	handle := item.Wrapper().Handle()
	menu.NSMenuItem_setKeyEquivalent(handle, key)
	menu.NSMenuItem_setKeyEquivalentModifierMask(handle, m)
}

func (item *menuItemBase) Mnemonic() rune {
	return 0
}

func (item *menuItemBase) SetMnemonic(k rune) {
	// Do nothing.
}

func newMenuItemTemplate(separator bool) MenuItem {
	item := &menuItemBase{}
	return item
}

type MenuItemHandleManager struct {
	objcHandleManagerBase
}

func (h *MenuItemHandleManager) Create(util.Bundle) native.Handle {
	return deallochook.Apply(menu.NewMenuItem())
}

type MenuSeparatorHandleManager struct {
	objcHandleManagerBase
}

func (h *MenuSeparatorHandleManager) Create(util.Bundle) native.Handle {
	return deallochook.Apply(menu.NewSeparatorMenuItem())
}
