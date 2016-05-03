#import <Foundation/NSString.h>
#import <AppKit/NSMenu.h>
#import <AppKit/NSCell.h>
#import "menu.h"

OBJC_PTR NSMenu_alloc() {
	return [NSMenu alloc];
}

OBJC_PTR NSMenu_initWithTitle(OBJC_PTR ptr, char* title) {
	if(title == NULL) {
		title = "";
	}
	return [(NSMenu*)ptr initWithTitle:[NSString stringWithUTF8String:title]];
}

void NSMenu_insertItemAtIndex(OBJC_PTR ptr, OBJC_PTR item, int index) {
	[(NSMenu*)ptr insertItem:(NSMenuItem*)item atIndex:index];
}

void NSMenu_addItem(OBJC_PTR ptr, OBJC_PTR item) {
	[(NSMenu*)ptr addItem:(NSMenuItem*)item];
}

int NSMenu_numberOfItems(OBJC_PTR ptr) {
	return [(NSMenu*)ptr numberOfItems];
}

void NSMenu_removeItem(OBJC_PTR ptr, OBJC_PTR item) {
	[(NSMenu*)ptr removeItem:(NSMenuItem*)item];
}

void NSMenu_removeItemAtIndex(OBJC_PTR ptr, int index) {
	[(NSMenu*)ptr removeItemAtIndex:index];
}

OBJC_PTR NSMenu_supermenu(OBJC_PTR ptr) {
	return [(NSMenu*)ptr supermenu];
}

void NSMenu_update(OBJC_PTR ptr) {
	[(NSMenu*)ptr update];
}

OBJC_PTR NSMenu_itemAtIndex(OBJC_PTR ptr, int index) {
	return [(NSMenu*)ptr itemAtIndex:index];
}

int NSMenu_indexOfItem(OBJC_PTR ptr, OBJC_PTR item) {
	return (int)[(NSMenu*)ptr indexOfItem:(NSMenuItem*)item];
}

char* NSMenu_title(OBJC_PTR ptr) {
	return (char*)[[(NSMenu*)ptr title] UTF8String];
}

void NSMenu_setTitle(OBJC_PTR ptr, char* title) {
	[(NSMenu*)ptr setTitle:[NSString stringWithUTF8String:title]];
}

bool NSMenu_autoenablesItems(OBJC_PTR ptr) {
	return [(NSMenu*)ptr autoenablesItems];
}

void NSMenu_setAutoenablesItems(OBJC_PTR ptr, bool v) {
	[(NSMenu*)ptr setAutoenablesItems:v];
}

//////////////////////////////////////////////////////////////
OBJC_PTR NSMenuItem_alloc() {
	return [NSMenuItem alloc];
}

OBJC_PTR NSMenuItem_initWithTitleItemNameAactionKeyEquivalent(OBJC_PTR ptr, char* title, UINTPTR action, char* keyEquv) {
	if(title == NULL) {
		title = "";
	}
	if(keyEquv == NULL) {
		keyEquv = "";
	}
	return [(NSMenuItem*)ptr initWithTitle:[NSString stringWithUTF8String:title] action:(SEL)(void*)action keyEquivalent:[NSString stringWithUTF8String:keyEquv]];
}

OBJC_PTR NSMenuItem_separatorItem() {
	return [NSMenuItem separatorItem];
}

char* NSMenuItem_title(OBJC_PTR ptr) {
	return (char*)[[(NSMenuItem*)ptr title] UTF8String];
}

void NSMenuItem_setTitle(OBJC_PTR ptr, char* title) {
	[(NSMenuItem*)ptr setTitle:[NSString stringWithUTF8String:title]];
}

void NSMenuItem_setSubmenu(OBJC_PTR ptr, OBJC_PTR submenu) {
	[(NSMenuItem*)ptr setSubmenu:(NSMenu*)submenu];
}

OBJC_PTR NSMenuItem_submenu(OBJC_PTR ptr) {
	return [(NSMenuItem*)ptr submenu];
}

bool NSMenuItem_isHidden(OBJC_PTR ptr) {
	return [(NSMenuItem*)ptr isHidden];
}

void NSMenuItem_setHidden(OBJC_PTR ptr, bool hidden) {
	[(NSMenuItem*)ptr setHidden:hidden];
}

bool NSMenuItem_isEnabled(OBJC_PTR ptr) {
	return [(NSMenuItem*)ptr isEnabled];
}

void NSMenuItem_setEnabled(OBJC_PTR ptr, bool enabled) {
	[(NSMenuItem*)ptr setEnabled:enabled];
}

OBJC_PTR NSMenuItem_menu(OBJC_PTR ptr) {
	return [(NSMenuItem*)ptr menu];
}

bool NSMenuItem_isSeparatorItem(OBJC_PTR ptr) {
	return [(NSMenuItem*)ptr isSeparatorItem];
}

int NSMenuItem_NSOffState = NSOffState;
int NSMenuItem_NSOnState = NSOnState;
int NSMenuItem_NSMixedState = NSMixedState;

int NSMenuItem_state(OBJC_PTR ptr) {
	return [(NSMenuItem*)ptr state];
}

void NSMenuItem_setState(OBJC_PTR ptr, int state) {
	[(NSMenuItem*)ptr setState:state];
}

UINTPTR NSMenuItem_action(OBJC_PTR ptr) {
	return (UINTPTR)(void*)[(NSMenuItem*)ptr action];
}

void NSMenuItem_setAction(OBJC_PTR ptr, UINTPTR selAction) {
	[(NSMenuItem*)ptr setAction:(SEL)(void*)selAction];
}

char* NSMenuItem_keyEquivalent(OBJC_PTR ptr) {
	return (char*)[[(NSMenuItem*)ptr keyEquivalent] UTF8String];
}

void NSMenuItem_setKeyEquivalent(OBJC_PTR ptr, char* k) {
	[(NSMenuItem*)ptr setKeyEquivalent:[NSString stringWithUTF8String:k]];
}

unsigned int NSMenuItem_keyEquivalentModifierMask(OBJC_PTR ptr) {
	return (unsigned int)[(NSMenuItem*)ptr keyEquivalentModifierMask];
}

unsigned int NSKeyEquivalentModifierMask_NSShiftKeyMask = NSShiftKeyMask;
unsigned int NSKeyEquivalentModifierMask_NSAlternateKeyMask = NSAlternateKeyMask;
unsigned int NSKeyEquivalentModifierMask_NSCommandKeyMask = NSCommandKeyMask;
unsigned int NSKeyEquivalentModifierMask_NSControlKeyMask = NSControlKeyMask;

void NSMenuItem_setKeyEquivalentModifierMask(OBJC_PTR ptr, unsigned int k) {
	[(NSMenuItem*)ptr setKeyEquivalentModifierMask:(NSUInteger)k];
}


