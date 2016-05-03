#import <AppKit/NSControl.h>
#import "control.h"

OBJC_PTR NSControl_initWithFrame(OBJC_PTR ptr, CGRect frameRect) {
	return [(NSControl*)ptr initWithFrame:frameRect];
}

OBJC_PTR NSControl_target(OBJC_PTR ptr) {
	return [(NSControl*)ptr target];
}

void NSControl_setTarget(OBJC_PTR ptr, OBJC_PTR target) {
	[(NSControl*)ptr setTarget:target];
}

UINTPTR NSControl_action(OBJC_PTR ptr) {
	return (UINTPTR)(void*)[(NSControl*)ptr action];
}

void NSControl_setAction(OBJC_PTR ptr, UINTPTR selAction) {
	[(NSControl*)ptr setAction:(SEL)(void*)selAction];
}

void NSControl_setStringValue(OBJC_PTR ptr, char* value) {
	[(NSControl*)ptr setStringValue:[NSString stringWithUTF8String:value]];
}

char* NSControl_stringValue(OBJC_PTR ptr) {
	return (char*)[[(NSControl*)ptr stringValue] UTF8String];
}

bool NSControl_isEnabled(OBJC_PTR ptr) {
	return [(NSControl*)ptr isEnabled];
}

void NSControl_setEnabled(OBJC_PTR ptr, bool enabled) {
	[(NSControl*)ptr setEnabled:enabled];
}

OBJC_PTR NSControl_currentEditor(OBJC_PTR ptr) {
	return [(NSControl*)ptr currentEditor];
}