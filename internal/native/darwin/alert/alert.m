#import <AppKit/NSAlert.h>
#import "alert.h"
#import "_cgo_export.h"

const NSUInteger VarNSWarningAlertStyle = NSWarningAlertStyle;
const NSUInteger VarNSInformationalAlertStyle = NSInformationalAlertStyle;
const NSUInteger VarNSCriticalAlertStyle = NSCriticalAlertStyle;
const NSInteger VarNSAlertFirstButtonReturn = NSAlertFirstButtonReturn;

OBJC_PTR NSAlert_alloc() {
	return [NSAlert alloc];
}

OBJC_PTR NSAlert_alertWithError(OBJC_PTR error) {
	return [NSAlert alertWithError:(NSError*)error];
}

void NSAlert_layout(OBJC_PTR ptr) {
	[(NSAlert*)ptr layout];
}

NSUInteger NSAlert_alertStyle(OBJC_PTR ptr) {
	return [(NSAlert*)ptr alertStyle];
}

void NSAlert_setAlertStyle(OBJC_PTR ptr, NSUInteger style) {
	[(NSAlert*)ptr setAlertStyle:style];
}

OBJC_PTR NSAlert_accessoryView(OBJC_PTR ptr) {
	return [(NSAlert*)ptr accessoryView];
}

void NSAlert_setAccessoryView(OBJC_PTR ptr, OBJC_PTR view) {
	[(NSAlert*)ptr setAccessoryView:(NSView*)view];
}

bool NSAlert_showsHelp(OBJC_PTR ptr) {
	return [(NSAlert*)ptr showsHelp] == YES;
}

void NSAlert_setShowsHelp(OBJC_PTR ptr, bool show) {
	[(NSAlert*)ptr setShowsHelp:(show ? YES : NO)];
}

char* NSAlert_helpAnchor(OBJC_PTR ptr) {
	return (char*)[[(NSAlert*)ptr helpAnchor] UTF8String];
}

void NSAlert_setHelpAnchor(OBJC_PTR ptr, char* anchor) {
	[(NSAlert*)ptr setHelpAnchor:[NSString stringWithUTF8String:anchor]];
}

OBJC_PTR NSAlert_delegate(OBJC_PTR ptr) {
	return [(NSAlert*)ptr delegate];
}

void NSAlert_setDelegate(OBJC_PTR ptr, OBJC_PTR delegate) {
	[(NSAlert*)ptr setDelegate:(id<NSAlertDelegate>)delegate];
}

NSInteger NSAlert_runModal(OBJC_PTR ptr) {
	return [(NSAlert*)ptr runModal];
}

static void NSAlert_beginSheetModalForWindow_completionHandler_AlertHandler(UINTPTR handlerId, int returnCode) {
	// Call into go.
	goBeginSheetModalForWindowCompletionHandler(handlerId, returnCode);
}

void NSAlert_beginSheetModalForWindow_completionHandler(OBJC_PTR ptr, OBJC_PTR sheetWindow, UINTPTR handlerId) {
	// https://developer.apple.com/library/ios/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/WorkingwithBlocks/WorkingwithBlocks.html
	void (^handler)(NSModalResponse returnCode) = handlerId ? 
		^(NSModalResponse returnCode){
			NSAlert_beginSheetModalForWindow_completionHandler_AlertHandler(handlerId, returnCode);
		}
		: nil;
	[(NSAlert*)ptr beginSheetModalForWindow:(NSWindow*)sheetWindow completionHandler:handler];
}

OBJC_PTR NSAlert_suppressionButton(OBJC_PTR ptr) {
	return [(NSAlert*)ptr suppressionButton];
}

bool NSAlert_showsSuppressionButton(OBJC_PTR ptr) {
	return [(NSAlert*)ptr showsSuppressionButton] == YES;
}

void NSAlert_setShowsSuppressionButton(OBJC_PTR ptr, bool show) {
	[(NSAlert*)ptr setShowsSuppressionButton:(show ? YES : NO)];
}

char* NSAlert_informativeText(OBJC_PTR ptr) {
	return (char*)[[(NSAlert*)ptr informativeText] UTF8String];
}

void NSAlert_setInformativeText(OBJC_PTR ptr, char* text) {
	[(NSAlert*)ptr setInformativeText:[NSString stringWithUTF8String:text]];
}

char* NSAlert_messageText(OBJC_PTR ptr) {
	return (char*)[[(NSAlert*)ptr messageText] UTF8String];
}

void NSAlert_setMessageText(OBJC_PTR ptr, char* text) {
	[(NSAlert*)ptr setMessageText:[NSString stringWithUTF8String:text]];
}

OBJC_PTR NSAlert_icon(OBJC_PTR ptr) {
	return [(NSAlert*)ptr icon];
}

void NSAlert_setIcon(OBJC_PTR ptr, OBJC_PTR icon) {
	[(NSAlert*)ptr setIcon:(NSImage*)icon];
}

OBJC_PTR NSAlert_buttons(OBJC_PTR ptr) {
	return [(NSAlert*)ptr buttons];
}

OBJC_PTR NSAlert_addButtonWithTitle(OBJC_PTR ptr, char* title) {
	return [(NSAlert*)ptr addButtonWithTitle:[NSString stringWithUTF8String:title]];
}

OBJC_PTR NSAlert_window(OBJC_PTR ptr) {
	return [(NSAlert*)ptr window];
}
