#import <AppKit/NSView.h>
#import <AppKit/NSColor.h>
#import "RWFlippedView.h"

//Working with the View Hierarchy
//https://developer.apple.com/library/prerelease/mac/documentation/Cocoa/Conceptual/CocoaViewsGuide/WorkingWithAViewHierarchy/WorkingWithAViewHierarchy.html

@implementation RWFlippedView

- (BOOL) isFlipped {
    return YES;
}

- (void)setAcceptFirstResponder:(BOOL)aAccept {
	_acceptFirstResponsder = aAccept;
}

- (BOOL) acceptsFirstResponder {
	return _acceptFirstResponsder;
}

- (void)setBackgroundColor:(NSColor*)aColor {
	[_backgroundColor release];
	_backgroundColor = [aColor retain];
}

- (NSColor*)backgroundColor {
	if(_backgroundColor) {
		return _backgroundColor;
	} else {
		return [NSColor clearColor];
	}
}

- (void)drawRect:(NSRect)dirtyRect {
	if(_backgroundColor) {
	    [_backgroundColor set];
	    NSRectFill(dirtyRect);
	}
}

- (void)dealloc {
	[_backgroundColor release];
	[super dealloc];
}
@end
