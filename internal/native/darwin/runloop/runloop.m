#import <Foundation/NSRunLoop.h>
#import <Foundation/NSConnection.h>
#import <Appkit/NSApplication.h>
#import "runloop.h"

//http://blog.ibireme.com/2015/05/18/runloop/

OBJC_PTR getNSRunLoopCommonModes() {
	return NSRunLoopCommonModes;
}

OBJC_PTR getNSDefaultRunLoopMode() {
	return NSDefaultRunLoopMode;
}

OBJC_PTR getNSConnectionReplyMode() {
	return NSConnectionReplyMode;
}

OBJC_PTR getNSModalPanelRunLoopMode() {
	return NSModalPanelRunLoopMode;
}

OBJC_PTR getNSEventTrackingRunLoopMode() {
	return NSEventTrackingRunLoopMode;
}

OBJC_PTR NSRunLoop_currentRunLoop() {
	return [NSRunLoop currentRunLoop];
}

OBJC_PTR NSRunloop_currentMode(OBJC_PTR ptr) {
	return [(NSRunLoop*)ptr currentMode];
}
