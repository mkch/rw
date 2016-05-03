#import <AppKit/NSWindow.h>
#import "window.h"
#import "RWWindow.h"
#import "_cgo_export.h"

OBJC_PTR NSWindow_alloc() {
	return [NSWindow alloc];
}

OBJC_PTR RWWindow_alloc() {
	return [RWWindow alloc];
}

OBJC_PTR NSWindow_initWithContentRect_styleMask_backing_defer_screen(OBJC_PTR ptr,
		CGRect rect, 
		UINTPTR styleMask, UINTPTR bufferingType, bool deferCreation, OBJC_PTR screen) {
	return [(NSWindow*)ptr initWithContentRect:rect
                                     styleMask:styleMask
                                       backing:bufferingType
                                         defer:deferCreation
                                        screen:(NSScreen*)screen];
}

char* NSWindow_title(OBJC_PTR ptr) {
	return (char*)[[(NSWindow*)ptr title] UTF8String];
}
void NSWindow_setTitle(OBJC_PTR ptr, char* title) {
	[(NSWindow*)ptr setTitle:[NSString stringWithUTF8String:title]];
}

void NSWindow_makeKeyAndOrderFront(OBJC_PTR ptr) {
	[(NSWindow*)ptr makeKeyAndOrderFront:nil];
}

void NSWindow_cascadeTopLeftFromPoint(OBJC_PTR ptr, CGPoint topLeft) {
	[(NSWindow*)ptr cascadeTopLeftFromPoint:topLeft];
}

void NSWindow_setFrameTopLeftPoint(OBJC_PTR ptr, CGPoint topLeft) {
	[(NSWindow*)ptr setFrameTopLeftPoint:topLeft];
}

OBJC_PTR NSWindow_contentView(OBJC_PTR ptr) {
	return [(NSWindow*)ptr contentView];
}

void NSWindow_setContentView(OBJC_PTR ptr, OBJC_PTR view) {
	[(NSWindow*)ptr setContentView:view];
}

void NSWindow_setDelegate(OBJC_PTR ptr, OBJC_PTR delegate) {
	[(NSWindow*)ptr setDelegate:(id)delegate];
	
}

OBJC_PTR NSWindow_delegate(OBJC_PTR ptr) {
	return [(NSWindow*)ptr delegate];
}

CGRect NSWindow_frameRectForContentRect_styleMask(CGRect windowContentRect, UINTPTR windowStyle) {
	return [NSWindow frameRectForContentRect:windowContentRect styleMask:windowStyle];
}

CGRect NSWindow_contentRectForFrameRect_styleMask(CGRect windowFrameRect, UINTPTR windowStyle) {
	return [NSWindow contentRectForFrameRect:windowFrameRect styleMask:windowStyle];
}

CGRect NSWindow_frame(OBJC_PTR ptr) {
	return [(NSWindow*)ptr frame];
}

bool NSWindow_makeFirstResponder(OBJC_PTR ptr, OBJC_PTR responder) {
	return [(NSWindow*)ptr makeFirstResponder:(NSResponder*)responder];
}

void NSWindow_setFrameDisplay(OBJC_PTR ptr, CGRect frame, bool displayViews) {
	[(NSWindow*)ptr setFrame:frame display:displayViews];
}

void NSWindow_close(OBJC_PTR ptr) {
	[(NSWindow*)ptr close];
}

void NSWindow_performClose(OBJC_PTR ptr, OBJC_PTR sender) {
	[(NSWindow*)ptr performClose:(id)sender];
}

void NSWindow_center(OBJC_PTR ptr) {
	[(NSWindow*)ptr center];
}

int VarNSModalResponseStop = NSModalResponseStop;
int VarNSModalResponseAbort = NSModalResponseAbort;
int VarNSModalResponseContinue = NSModalResponseContinue;

static void NSWindow_beginSheet_completionHandler_SheetHandler(UINTPTR handlerId, int returnCode) {
	// Call into go.
	goBeginSheetCompletionHandler(handlerId, returnCode);
}

void NSWindow_beginSheet_completionHandler(OBJC_PTR ptr, OBJC_PTR win, UINTPTR handlerId) {
	// https://developer.apple.com/library/ios/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/WorkingwithBlocks/WorkingwithBlocks.html
	void (^handler)(NSModalResponse returnCode) = handlerId ? 
		^(NSModalResponse returnCode){
			NSWindow_beginSheet_completionHandler_SheetHandler(handlerId, returnCode);
		}
		: nil;
	[(NSWindow*)ptr beginSheet:(NSWindow*)win completionHandler:handler];
}

void NSWindow_endSheet_returnCode(OBJC_PTR ptr, OBJC_PTR win, int returnCode) {
	[(NSWindow*)ptr endSheet:(NSWindow*)win returnCode:returnCode];
}

OBJC_PTR NSWindow_attachedSheet(OBJC_PTR ptr) {
	return [(NSWindow*)ptr attachedSheet];
}

bool NSWindow_isSheet(OBJC_PTR ptr) {
	return [(NSWindow*)ptr isSheet];
}

OBJC_PTR NSWindow_sheetParent(OBJC_PTR ptr) {
	return [(NSWindow*)ptr sheetParent];
}

void NSWindow_discardEventsMatchingMask_beforeEvent(OBJC_PTR ptr, unsigned long eventMask, OBJC_PTR lastEvent) {
	[(NSWindow*)ptr discardEventsMatchingMask:eventMask beforeEvent:(NSEvent*)lastEvent];
}

void NSWindow_sendEvent(OBJC_PTR ptr, OBJC_PTR event) {
	[(NSWindow*)ptr sendEvent:(NSEvent*)event];
}

void NSWindow_display(OBJC_PTR ptr) {
	[(NSWindow*)ptr display];
}

void NSWindow_displayIfNeeded(OBJC_PTR ptr) {
	[(NSWindow*)ptr displayIfNeeded];
}

bool NSWindow_isVisible(OBJC_PTR ptr) {
	return [(NSWindow*)ptr isVisible];
}

void NSWindow_orderOut(OBJC_PTR ptr) {
	[(NSWindow*)ptr orderOut:(NSWindow*)ptr];
}

void NSWindow_orderBack(OBJC_PTR ptr) {
	[(NSWindow*)ptr orderBack:(NSWindow*)ptr];
}

void NSWindow_orderFront(OBJC_PTR ptr) {
	[(NSWindow*)ptr orderFront:(NSWindow*)ptr];
}

void NSWindow_orderFrontRegardless(OBJC_PTR ptr) {
	[(NSWindow*)ptr orderFrontRegardless];
}

long VarNSWindowAbove = NSWindowAbove;
long VarNSWindowBelow = NSWindowBelow;
long VarNSWindowOut = NSWindowOut;

void NSWindow_orderWindow_relativeTo(OBJC_PTR ptr, long orderingMode, long otherWindowNumber) {
	[(NSWindow*)ptr orderWindow:orderingMode relativeTo:otherWindowNumber];
}

long NSWindow_level(OBJC_PTR ptr) {
	return [(NSWindow*)ptr level];
}

void NSWindow_setLevel(OBJC_PTR ptr, long level) {
	[(NSWindow*)ptr setLevel: level];
}

OBJC_PTR NSWindow_firstResponder(OBJC_PTR ptr) {
	return [(NSWindow*)ptr firstResponder];
}

OBJC_PTR NSWindow_screen(OBJC_PTR ptr) {
	return [(NSWindow*)ptr screen];
}

// RWWindow

bool RWWindow_enabled(OBJC_PTR ptr) {
	return [(RWWindow*)ptr enabled];
}

void RWWindow_setEnabled(OBJC_PTR ptr, bool enabled) {
	[(RWWindow*)ptr setEnabled:enabled];
}
