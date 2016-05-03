#ifndef _RW_BUTTON_H
#define _RW_BUTTON_H

#include "../types.h"

OBJC_PTR NSTextField_alloc();
bool NSTextField_isEditable(OBJC_PTR ptr);
void NSTextField_setEditable(OBJC_PTR ptr, bool value);
bool NSTextField_isSelectable(OBJC_PTR ptr);
void NSTextField_setSelectable(OBJC_PTR ptr, bool value);
OBJC_PTR NSTextField_textColor(OBJC_PTR ptr);
void NSTextField_setTextColor(OBJC_PTR ptr, OBJC_PTR color);
OBJC_PTR NSTextField_backgroundColor(OBJC_PTR ptr);
void NSTextField_setBackgroundColor(OBJC_PTR ptr, OBJC_PTR color);

OBJC_PTR RWTextFieldDelegate_init();
int RWTextFieldDelegate_multiline(OBJC_PTR ptr);
void RWTextFieldDelegate_setMultiline(OBJC_PTR ptr, bool value);


#endif