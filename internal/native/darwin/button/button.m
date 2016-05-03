#import <AppKit/NSButton.h>
#import "button.h"


OBJC_PTR NSButton_alloc() {
	return [NSButton alloc];
}

void NSButton_setButtonType(OBJC_PTR ptr, UINTPTR buttonType) {
	[(NSButton*)ptr setButtonType:buttonType];
}

void NSButton_setBezelStyle(OBJC_PTR ptr, UINTPTR bezelStyle) {
	[(NSButton*)ptr setBezelStyle:bezelStyle];
}

char* NSButton_title(OBJC_PTR ptr) {
	return (char*)[((NSButton*)ptr).title UTF8String];
}

void NSButton_setTitle(OBJC_PTR ptr, char* title) {
	[(NSButton*)ptr setTitle:[NSString stringWithUTF8String:title]];
}

