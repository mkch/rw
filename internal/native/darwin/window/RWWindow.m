#import <AppKit/NSWindow.h>
#import "RWWindow.h"

@implementation RWWindow

- (void)sendEvent:(NSEvent *)event {
	if(!_disabled) {
		[super sendEvent:event];
	}
}

- (BOOL) enabled {
	return !_disabled;
}

- (void) setEnabled:(BOOL)aEnabled {
	_disabled = !aEnabled;
}

@end
