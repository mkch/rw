#import <AppKit/NSColor.h>
#import <AppKit/NSColorSpace.h>
#import "color.h"

// https://developer.apple.com/library/mac/documentation/Cocoa/Conceptual/DrawColor/Tasks/UsingColorSpaces.html

OBJC_PTR NSColor_colorWithRGB(CGFloat r, CGFloat g, CGFloat b, CGFloat a) {
	return [NSColor colorWithCalibratedRed:r green:g blue:b alpha:a];
}

void NSColor_getRGBA(OBJC_PTR ptr, CGFloat* r, CGFloat* g, CGFloat* b, CGFloat* a) {
	return [[(NSColor*)ptr colorUsingColorSpace:[NSColorSpace genericRGBColorSpace]] getRed:r green:g blue:b alpha:a];
}
