#import <AppKit/NSWindow.h>
#import "windowstyle.h"

UINTPTR getWindowStyle(WindowStyleFeatures* features) {
	UINTPTR style = NSBorderlessWindowMask;
	if(features->hasTitle) {
		style |= NSTitledWindowMask;
	}
	if(features->hasCloseButton) {
		style |= NSClosableWindowMask;
	}
	if(features->hasMinimizeButton) {
		style |= NSMiniaturizableWindowMask;
	}
	if(features->resizable || features->hasMaximizeButton) {
		style |= NSResizableWindowMask;
	}
	return style;
}