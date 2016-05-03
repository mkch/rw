#ifndef _RW_OBJECT_H
#define _RW_OBJECT_H

#include "../types.h"

void NSObject_retain(OBJC_PTR);
long NSObject_retainCount(OBJC_PTR);
OBJC_PTR NSObject_init(OBJC_PTR);
void NSObject_release(OBJC_PTR);
OBJC_PTR NSObject_autorelease(OBJC_PTR);
char* NSObject_description(OBJC_PTR);

// Sets an associated value for a given object using a given key with OBJC_ASSOCIATION_RETAIN association policy.
void NSObject_setAssociatedObjectRetain(OBJC_PTR ptr, void* key, OBJC_PTR value);
OBJC_PTR NSObject_getAssociatedObject(OBJC_PTR ptr, void* key);

// Sets an associated value for a given object using a given key with OBJC_ASSOCIATION_ASSIGN association policy.
void NSObject_setAssociatedUnsignedLong(OBJC_PTR ptr, void* key, unsigned long value);
unsigned long NSObject_getAssociatedUnsignedLong(OBJC_PTR ptr, void* key);

// Set delegate to obj  and let obj own delegate.
void Object_SetDelegateRetain(OBJC_PTR obj, OBJC_PTR delegate);
OBJC_PTR Object_getDelegate(OBJC_PTR obj);

void Object_SetTargetRetain(OBJC_PTR obj, OBJC_PTR target);
OBJC_PTR Object_getTarget(OBJC_PTR obj);

#endif
