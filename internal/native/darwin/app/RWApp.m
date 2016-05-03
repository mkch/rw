#import <AppKit/NSApplication.h>
#import "RWApp.h"

@implementation RWApp

- (void)superSendEvent:(NSEvent*)event {
	[super sendEvent:event];
}

- (void)sendEvent:(NSEvent *)event {
	// Call into go.
	goAppSendEvent(self, event);
}

@end