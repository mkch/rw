#ifndef _RW_NOTIFICATION_CENTER_H
#define _RW_NOTIFICATION_CENTER_H

#include "../types.h"

OBJC_PTR NSNotificationCenter_defaultCenter();
void NSNotificationCenter_addObserver_selector_name_object(OBJC_PTR ptr, OBJC_PTR observer, PVOID sel, char* name, OBJC_PTR object);
void NSNotificationCenter_removeObserver_name_object(OBJC_PTR ptr, OBJC_PTR observer, char* name, OBJC_PTR object);

#endif
