#import <AppKit/NSTextField.h>
#import <AppKit/NSColor.h>
#import "textfield.h"
#import "RWTextFieldDelegate.h"


OBJC_PTR NSTextField_alloc() {
	return [NSTextField alloc];
}

bool NSTextField_isEditable(OBJC_PTR ptr) {
	return [(NSTextField*)ptr isEditable];
}

void NSTextField_setEditable(OBJC_PTR ptr, bool value) {
	[(NSTextField*)ptr setEditable:value];
}

bool NSTextField_isSelectable(OBJC_PTR ptr) {
	return [(NSTextField*)ptr isSelectable];
}

void NSTextField_setSelectable(OBJC_PTR ptr, bool value) {
	[(NSTextField*)ptr setSelectable:value];
}

OBJC_PTR NSTextField_textColor(OBJC_PTR ptr) {
	return [(NSTextField*)ptr textColor];
}

void NSTextField_setTextColor(OBJC_PTR ptr, OBJC_PTR color) {
	[(NSTextField*)ptr setTextColor:(NSColor*)color];
}

OBJC_PTR NSTextField_backgroundColor(OBJC_PTR ptr) {
	return [(NSTextField*)ptr backgroundColor];
}

void NSTextField_setBackgroundColor(OBJC_PTR ptr, OBJC_PTR color) {
	[(NSTextField*)ptr setBackgroundColor:(NSColor*)color];
}

////////////////////////////////////////////

OBJC_PTR RWTextFieldDelegate_init() {
	return [[RWTextFieldDelegate alloc] init];
}

int RWTextFieldDelegate_multiline(OBJC_PTR ptr) {
	return [(RWTextFieldDelegate*)ptr multiline];
}

void RWTextFieldDelegate_setMultiline(OBJC_PTR ptr, bool value) {
	[(RWTextFieldDelegate*)ptr setMultiline:value];
}
