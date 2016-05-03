#import <AppKit/NSScreen.h>
#import "screen.h"

OBJC_PTR NSScreen_mainScreen() {
	return [NSScreen mainScreen];
}

CGRect NSScreen_visibleFrame(OBJC_PTR ptr) {
	return [(NSScreen*)ptr visibleFrame];
}

CGRect NSScreen_frame(OBJC_PTR ptr) {
	return [(NSScreen*)ptr frame];
}