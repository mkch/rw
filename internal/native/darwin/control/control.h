#ifndef _RW_CONTROL_H
#define _RW_CONTROL_H

#include "../types.h"
#include <CoreGraphics/CGGeometry.h>

OBJC_PTR NSControl_initWithFrame(OBJC_PTR ptr, CGRect frameRect);
OBJC_PTR NSControl_target(OBJC_PTR ptr);
void NSControl_setTarget(OBJC_PTR ptr, OBJC_PTR target);
void NSControl_setStringValue(OBJC_PTR ptr, char* value);
UINTPTR NSControl_action(OBJC_PTR ptr);
void NSControl_setAction(OBJC_PTR ptr, UINTPTR selAction);
char* NSControl_stringValue(OBJC_PTR ptr);
bool NSControl_isEnabled(OBJC_PTR ptr);
void NSControl_setEnabled(OBJC_PTR ptr, bool enabled);
OBJC_PTR NSControl_currentEditor(OBJC_PTR ptr);

#endif