#import <AppKit/NSTextView.h>
#import "textview.h"

void NSTextView_setSelectedRange(OBJC_PTR ptr, unsigned long loc, unsigned long len) {
	[(NSTextView*)ptr setSelectedRange:NSMakeRange(loc, len)];
}

OBJC_PTR/*(NSArray*)*/ NSTextView_selectedRanges(OBJC_PTR ptr) {
	return [(NSTextView*)ptr selectedRanges];
}