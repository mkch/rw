#import <Foundation/NSString.h>
#import <objc/runtime.h>
#import "object.h"

void NSObject_retain(OBJC_PTR ptr) {
	[(id)ptr retain];
}

long NSObject_retainCount(OBJC_PTR ptr) {
	return [(id)ptr retainCount];
}

OBJC_PTR NSObject_init(OBJC_PTR ptr) {
	return [(id)ptr init];
}

void NSObject_release(OBJC_PTR ptr) {
	return [(id)ptr release];
}

OBJC_PTR NSObject_autorelease(OBJC_PTR ptr) {
	return [(id)ptr autorelease];
}

char* NSObject_description(OBJC_PTR ptr) {
	return (char*)[[(id)ptr description] UTF8String];
}

void NSObject_setAssociatedObjectRetain(OBJC_PTR ptr, void* key, OBJC_PTR value) {
	objc_setAssociatedObject((id)ptr, (const void*)key, (id)value, OBJC_ASSOCIATION_RETAIN);
}

OBJC_PTR NSObject_getAssociatedObject(OBJC_PTR ptr, void* key) {
	return objc_getAssociatedObject((id)ptr, (const void*)key);
}

void NSObject_setAssociatedUnsignedLong(OBJC_PTR ptr, void* key, unsigned long value) {
	objc_setAssociatedObject((id)ptr, (const void*)key, (id)value, OBJC_ASSOCIATION_ASSIGN);
}

unsigned long NSObject_getAssociatedUnsignedLong(OBJC_PTR ptr, void* key) {
	return (unsigned long)objc_getAssociatedObject((id)ptr, (const void*)key);
}
 
void Object_SetDelegateRetain(OBJC_PTR obj, OBJC_PTR delegate) {
	[(id)obj performSelector:@selector(setDelegate:) withObject:(id)delegate];
	static char _key;
	NSObject_setAssociatedObjectRetain(obj, &_key, delegate);
}

OBJC_PTR Object_getDelegate(OBJC_PTR obj) {
	return [(id)obj performSelector:@selector(delegate)];
}

void Object_SetTargetRetain(OBJC_PTR obj, OBJC_PTR target) {
	[(id)obj performSelector:@selector(setTarget:) withObject:(id)target];
	static char _key;
	NSObject_setAssociatedObjectRetain(obj, &_key, target);
}

OBJC_PTR Object_getTarget(OBJC_PTR obj) {
	return [(id)obj performSelector:@selector(target)];
}