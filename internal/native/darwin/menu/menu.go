package menu

//#include <stdlib.h>
//#include "menu.h"
import "C"

import (
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/internal/native/darwin/object"
	"github.com/kevin-yuan/rw/util/ustr"
)

func NewMenu() native.Handle {
	return native.Handle(C.NSMenu_initWithTitle(C.NSMenu_alloc(), nil))
}

func NSMenu_insertItemAtIndex(menu native.Handle, item native.Handle, index int) {
	C.NSMenu_insertItemAtIndex(C.OBJC_PTR(menu), C.OBJC_PTR(item), C.int(index))
}

func NSMenu_addItem(menu, item native.Handle) {
	C.NSMenu_addItem(C.OBJC_PTR(menu), C.OBJC_PTR(item))
}

func NSMenu_numberOfItems(menu native.Handle) int {
	return int(C.NSMenu_numberOfItems(C.OBJC_PTR(menu)))
}

func NSMenu_removeItem(menu, item native.Handle) {
	C.NSMenu_removeItem(C.OBJC_PTR(menu), C.OBJC_PTR(item))
}

func NSMenu_removeItemAtIndex(menu native.Handle, index int) {
	C.NSMenu_removeItemAtIndex(C.OBJC_PTR(menu), C.int(index))
}

func NSMenu_superMenu(menu native.Handle) native.Handle {
	return native.Handle(C.NSMenu_supermenu(C.OBJC_PTR(menu)))
}

func NSMenu_update(menu native.Handle) {
	C.NSMenu_update(C.OBJC_PTR(menu))
}

func NSMenu_itemAtIndex(menu native.Handle, index int) native.Handle {
	return native.Handle(C.NSMenu_itemAtIndex(C.OBJC_PTR(menu), C.int(index)))
}

func NSMenu_indexOfItem(menu, item native.Handle) int {
	return int(C.NSMenu_indexOfItem(C.OBJC_PTR(menu), C.OBJC_PTR(item)))
}

func NSMenu_title(menu native.Handle) string {
	return C.GoString(C.NSMenu_title(C.OBJC_PTR(menu)))
}

func NSMenu_setTitle(menu native.Handle, title string) {
	C.NSMenu_setTitle(C.OBJC_PTR(menu), (*C.char)(ustr.CStringUtf8(title)))
}

func NSMenu_autoenablesItems(menu native.Handle) bool {
	return bool(C.NSMenu_autoenablesItems(C.OBJC_PTR(menu)))
}

func NSMenu_setAutoenablesItems(menu native.Handle, v bool) {
	C.NSMenu_setAutoenablesItems(C.OBJC_PTR(menu), C.bool(v))
}

//////////////////////////////////////////////


func NewMenuItem() native.Handle {
	return native.Handle(C.NSMenuItem_initWithTitleItemNameAactionKeyEquivalent(C.NSMenuItem_alloc(), nil, 0, nil))
}


func NewSeparatorMenuItem() native.Handle {
	return object.NSObject_retain(native.Handle(C.NSMenuItem_separatorItem()))
}

func NSMenuItem_title(item native.Handle) string {
	return C.GoString(C.NSMenuItem_title(C.OBJC_PTR(item)))
}

func NSMenuItem_setTitle(item native.Handle, title string) {
	C.NSMenuItem_setTitle(C.OBJC_PTR(item), (*C.char)(ustr.CStringUtf8(title)))
}

func NSMenuItem_setSubmenu(item, menu native.Handle) {
	C.NSMenuItem_setSubmenu(C.OBJC_PTR(item), C.OBJC_PTR(menu))
}

func NSMenuItem_submenu(item native.Handle) native.Handle {
	return native.Handle(C.NSMenuItem_submenu(C.OBJC_PTR(item)))
}

func NSMenuItem_isHidden(item native.Handle) bool {
	return bool(C.NSMenuItem_isHidden(C.OBJC_PTR(item)))
}

func NSMenuItem_setHidden(item native.Handle, hidden bool) {
	C.NSMenuItem_setHidden(C.OBJC_PTR(item), C.bool(hidden))
}

func NSMenuItem_isEnabled(item native.Handle) bool {
	return bool(C.NSMenuItem_isEnabled(C.OBJC_PTR(item)))
}

func NSMenuItem_setEnabled(item native.Handle, enabled bool) {
	C.NSMenuItem_setEnabled(C.OBJC_PTR(item), C.bool(enabled))
}

func NSMenuItem_menu(item native.Handle) native.Handle {
	return native.Handle(C.NSMenuItem_menu(C.OBJC_PTR(item)))
}

func NSMenuItem_isSeparatorItem(item native.Handle) bool {
	return bool(C.NSMenuItem_isSeparatorItem(C.OBJC_PTR(item)))
}

type NSMenuItemState int
var(
	NSOffState = NSMenuItemState(C.NSMenuItem_NSOffState)
	NSOnState = NSMenuItemState(C.NSMenuItem_NSOnState)
	NSMixedState = NSMenuItemState(C.NSMenuItem_NSMixedState)
)

func NSMenuItem_state(item native.Handle) NSMenuItemState {
	return NSMenuItemState(C.NSMenuItem_state(C.OBJC_PTR(item)))
}

func NSMenuItem_setState(item native.Handle, state NSMenuItemState) {
	C.NSMenuItem_setState(C.OBJC_PTR(item), C.int(state))
}

func NSMenuItem_action(item native.Handle) uintptr {
	return uintptr(C.NSMenuItem_action(C.OBJC_PTR(item)))
}

func NSMenuItem_setAction(item native.Handle, sel uintptr) {
	C.NSMenuItem_setAction(C.OBJC_PTR(item), C.UINTPTR(sel))
}

func NSMenuItem_keyEquivalent(item native.Handle) string {
	return C.GoString(C.NSMenuItem_keyEquivalent(C.OBJC_PTR(item)))
}

// 0 key to remove.
func NSMenuItem_setKeyEquivalent(item native.Handle, key rune) {
	strKey := "" // Apple Doc: The receiver’s unmodified keyboard equivalent, or the empty string if one hasn’t been defined.
	if key != 0 {
		strKey = string(key)
	}
	C.NSMenuItem_setKeyEquivalent(C.OBJC_PTR(item), (*C.char)(ustr.CStringUtf8(strKey)))
}

type NSKeyEquivalentModifierMask uint
var (
	NSShiftKeyMask = NSKeyEquivalentModifierMask(C.NSKeyEquivalentModifierMask_NSShiftKeyMask)
	NSAlternateKeyMask = NSKeyEquivalentModifierMask(C.NSKeyEquivalentModifierMask_NSAlternateKeyMask)
	NSCommandKeyMask = NSKeyEquivalentModifierMask(C.NSKeyEquivalentModifierMask_NSCommandKeyMask)
	NSControlKeyMask = NSKeyEquivalentModifierMask(C.NSKeyEquivalentModifierMask_NSControlKeyMask)
)

func NSMenuItem_keyEquivalentModifierMask(item native.Handle) NSKeyEquivalentModifierMask {
	return NSKeyEquivalentModifierMask(C.NSMenuItem_keyEquivalentModifierMask(C.OBJC_PTR(item)))
}

func NSMenuItem_setKeyEquivalentModifierMask(item native.Handle, k NSKeyEquivalentModifierMask) {
	C.NSMenuItem_setKeyEquivalentModifierMask(C.OBJC_PTR(item), C.uint(k))
}

