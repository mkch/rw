#import <AppKit/NSView.h>
#import "RWFlippedView.h"
#import "view.h"

OBJC_PTR NSView_initWithFrame(OBJC_PTR ptr, CGRect rect) {
	return [(NSView*)ptr initWithFrame:rect];
}

void NSView_addSubview(OBJC_PTR ptr, OBJC_PTR subView) {
	[(NSView*)ptr addSubview:subView];
}

void NSView_removeFromSuperview(OBJC_PTR ptr) {
	[(NSView*)ptr removeFromSuperview];
}

OBJC_PTR NSView_superview(OBJC_PTR ptr) {
	return [(NSView*)ptr superview];
}

OBJC_PTR NSView_window(OBJC_PTR ptr) {
	return [(NSView*)ptr window];
}

CGRect NSView_frame(OBJC_PTR ptr) {
	return [(NSView*)ptr frame];
}

void NSView_setFrameSize(OBJC_PTR ptr, CGSize size) {
	[(NSView*)ptr setFrameSize:size];
}

void NSView_setFrameOrigin(OBJC_PTR ptr, CGPoint origin) {
	[(NSView*)ptr setFrameOrigin:origin];
}

void NSView_setPostsFrameChangedNotifications(OBJC_PTR ptr, bool post) {
	[(NSView*)ptr setPostsFrameChangedNotifications:(post ? YES : NO)];
}

char* NSView_NSViewFrameDidChangeNotification() {
	return (char*)[NSViewFrameDidChangeNotification UTF8String];
}

bool NSView_isHidden(OBJC_PTR ptr) {
	return [(NSView*)ptr isHidden] == YES;
}

void NSView_setHidden(OBJC_PTR ptr, bool hidden) {
	[(NSView*)ptr setHidden:(hidden ? YES : NO)];
}

OBJC_PTR NSView_subviews(OBJC_PTR ptr) {
	return [(NSView*)ptr subviews];
}

void NSView_display(OBJC_PTR ptr) {
	[(NSView*)ptr display];
}

void NSView_displayIfNeeded(OBJC_PTR ptr) {
	[(NSView*)ptr displayIfNeeded];
}

bool NSView_needsDisplay(OBJC_PTR ptr) {
	return [(NSView*)ptr needsDisplay];
}

void NSView_setNeedsDisplay(OBJC_PTR ptr, bool needs) {
	[(NSView*)ptr setNeedsDisplay:needs];
}

OBJC_PTR NSView_nextKeyView(OBJC_PTR ptr) {
	return [(NSView*)ptr nextKeyView];
}

void NSView_setNextKeyView(OBJC_PTR ptr, OBJC_PTR keyView) {
	[(NSView*)ptr setNextKeyView:(NSView*)keyView];
}

OBJC_PTR NSView_previousKeyView(OBJC_PTR ptr) {
	return [(NSView*)ptr previousKeyView];
}



// RWFlippedView 

OBJC_PTR RWFlippedView_alloc() {
	return [RWFlippedView alloc];
}

OBJC_PTR RWFlippedView_backgroundColor(OBJC_PTR ptr) {
	return [(RWFlippedView*)ptr backgroundColor];
}

void RWFlippedView_setBackgroundColor(OBJC_PTR ptr, OBJC_PTR color) {
	[(RWFlippedView*)ptr setBackgroundColor:(NSColor*)color];
}

void RWFlippedView_setAcceptFirstResponder(OBJC_PTR ptr, bool accept) {
	[(RWFlippedView*)ptr setAcceptFirstResponder:accept];
}
