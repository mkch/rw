#import <AppKit/NSApplication.h>
#import <AppKit/NSMenu.h>
#import "app.h"
#import "RWApp.h"


// https://developer.apple.com/library/mac/documentation/Cocoa/Conceptual/Multithreading/RunLoopManagement/RunLoopManagement.html
// https://developer.apple.com/library/mac/documentation/Cocoa/Conceptual/WinPanel/Concepts/UsingModalWindows.html

OBJC_PTR getNSApp() {
	return NSApp;
}

void NSApplication_terminate(OBJC_PTR ptr, OBJC_PTR sender) {
	[(NSApplication*)ptr terminate:(id)sender];
}

void NSApplication_stop(OBJC_PTR ptr, OBJC_PTR sender) {
	[(NSApplication*)ptr stop:(id)sender];
	NSEvent* event =[NSEvent otherEventWithType:NSApplicationDefined
		location:NSMakePoint(0,0)
		modifierFlags:0
		timestamp:0
		windowNumber:0
		context:nil
		subtype:0
		data1:0
		data2:0];
	[(NSApplication*)ptr postEvent:event atStart:NO];
}

void NSApplication_setMainMenu(OBJC_PTR ptr, OBJC_PTR menu) {
	[(NSApplication*)ptr setMainMenu:(NSMenu*)menu];
}

OBJC_PTR NSApplication_mainMenu(OBJC_PTR ptr) {
	return [(NSApplication*)ptr mainMenu];
}

UINTPTR NSApplication_runModalForWindow(OBJC_PTR ptr, OBJC_PTR win) {
	return (UINTPTR)[(NSApplication*)ptr runModalForWindow:(NSWindow*)win];
}

UINTPTR NSApplication_NSModalResponseStop = (UINTPTR)NSModalResponseStop;
UINTPTR NSApplication_NSModalResponseAbort = (UINTPTR)NSModalResponseAbort;
UINTPTR NSApplication_NSModalResponseContinue = (UINTPTR)NSModalResponseContinue;

void NSApplication_stopModalWithCode(OBJC_PTR ptr, UINTPTR code) {
	[(NSApplication*)ptr stopModalWithCode:(NSInteger)code];
}

void NSApplication_abortModal(OBJC_PTR ptr) {
	[(NSApplication*)ptr abortModal];
}

OBJC_PTR NSApplication_modalWindow(OBJC_PTR ptr) {
	return [(NSApplication*)ptr modalWindow];
}

long VarNSApplicationActivationPolicyRegular = NSApplicationActivationPolicyRegular;
long VarNSApplicationActivationPolicyAccessory = NSApplicationActivationPolicyAccessory;
long VarNSApplicationActivationPolicyProhibited = NSApplicationActivationPolicyProhibited;

void NSApplication_setActivationPolicy(OBJC_PTR ptr, long ploicy) {
	[(NSApplication*)ptr setActivationPolicy:ploicy];
}

void NSApplication_run(OBJC_PTR ptr) {
	[(NSApplication*)ptr run];
}

OBJC_PTR NSApplication_windows(OBJC_PTR ptr) {
	return [(NSApplication*)ptr windows];
}

void NSApplication_sendEvent(OBJC_PTR ptr, OBJC_PTR event) {
	[(NSApplication*)ptr sendEvent:(NSEvent*)event];
}

void NSApplication_postEvent_atStart(OBJC_PTR ptr, OBJC_PTR event, bool flag) {
	[(NSApplication*)ptr postEvent:(NSEvent*)event atStart:flag?YES:NO];
}

OBJC_PTR NSApplication_nextEventMatchingMask_untilDate_inMode_dequeue(OBJC_PTR ptr, UINTPTR mask, OBJC_PTR expiration, OBJC_PTR mode, bool flag) {
	return [(NSApplication*)ptr nextEventMatchingMask:(NSUInteger)mask
		untilDate: expiration
		inMode: mode
		dequeue: flag ? YES : NO];
}

// RWApp

OBJC_PTR RWApp_sharedApplication() {
	return [RWApp sharedApplication];
}

void RWApp_superSendEvent(OBJC_PTR ptr, OBJC_PTR event) {
	[(RWApp*)ptr superSendEvent:(NSEvent*)event];
}