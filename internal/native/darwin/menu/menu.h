#ifndef _RW_MENU_H
#define _RW_MENU_H

#include "../types.h"

OBJC_PTR NSMenu_alloc();
OBJC_PTR NSMenu_initWithTitle(OBJC_PTR ptr, char* title);
void NSMenu_insertItemAtIndex(OBJC_PTR ptr, OBJC_PTR item, int index);
void NSMenu_addItem(OBJC_PTR ptr, OBJC_PTR item);
int NSMenu_numberOfItems(OBJC_PTR ptr);
void NSMenu_removeItem(OBJC_PTR ptr, OBJC_PTR item);
void NSMenu_removeItemAtIndex(OBJC_PTR ptr, int index);
int NSMenu_indexOfItem(OBJC_PTR ptr, OBJC_PTR item);
OBJC_PTR NSMenu_supermenu(OBJC_PTR ptr);
void NSMenu_update(OBJC_PTR ptr);
OBJC_PTR NSMenu_itemAtIndex(OBJC_PTR ptr, int index);
char* NSMenu_title(OBJC_PTR ptr);
void NSMenu_setTitle(OBJC_PTR ptr, char* title);
bool NSMenu_autoenablesItems(OBJC_PTR ptr);
void NSMenu_setAutoenablesItems(OBJC_PTR ptr, bool v);
//////////////////////////////////////////////////////
OBJC_PTR NSMenuItem_alloc();
OBJC_PTR NSMenuItem_initWithTitleItemNameAactionKeyEquivalent(OBJC_PTR ptr, char* title, UINTPTR action, char* keyEquv);
OBJC_PTR NSMenuItem_separatorItem();
char* NSMenuItem_title(OBJC_PTR ptr);
void NSMenuItem_setTitle(OBJC_PTR ptr, char* title);
void NSMenuItem_setSubmenu(OBJC_PTR ptr, OBJC_PTR submenu);
OBJC_PTR NSMenuItem_submenu(OBJC_PTR ptr);
bool NSMenuItem_isHidden(OBJC_PTR ptr);
void NSMenuItem_setHidden(OBJC_PTR ptr, bool hidden);
bool NSMenuItem_isEnabled(OBJC_PTR ptr);
void NSMenuItem_setEnabled(OBJC_PTR ptr, bool enabled);
OBJC_PTR NSMenuItem_menu(OBJC_PTR ptr);
bool NSMenuItem_isSeparatorItem(OBJC_PTR ptr);
extern int NSMenuItem_NSOffState, NSMenuItem_NSOnState, NSMenuItem_NSMixedState;
int NSMenuItem_state(OBJC_PTR ptr);
void NSMenuItem_setState(OBJC_PTR ptr, int state);
UINTPTR NSMenuItem_action(OBJC_PTR ptr);
void NSMenuItem_setAction(OBJC_PTR ptr, UINTPTR selAction);
char* NSMenuItem_keyEquivalent(OBJC_PTR ptr);
void NSMenuItem_setKeyEquivalent(OBJC_PTR ptr, char* k);
extern unsigned int NSKeyEquivalentModifierMask_NSShiftKeyMask, NSKeyEquivalentModifierMask_NSAlternateKeyMask, NSKeyEquivalentModifierMask_NSCommandKeyMask, NSKeyEquivalentModifierMask_NSControlKeyMask;
unsigned int NSMenuItem_keyEquivalentModifierMask(OBJC_PTR ptr);
void NSMenuItem_setKeyEquivalentModifierMask(OBJC_PTR, unsigned int k);

#endif
