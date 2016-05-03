#ifndef _RW_BUTTON_H
#define _RW_BUTTON_H

#include "../types.h"

typedef void(*FN_BUTTON_ON_CLICK)();

OBJC_PTR NSButton_alloc();
void NSButton_setButtonType(OBJC_PTR ptr, UINTPTR buttonType);
void NSButton_setBezelStyle(OBJC_PTR ptr, UINTPTR bezelStyle);
char* NSButton_title(OBJC_PTR);
void NSButton_setTitle(OBJC_PTR ptr, char* title);


#endif