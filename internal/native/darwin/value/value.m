#import <Foundation/Foundation.h>
#import "value.h"

void NSValue_rangeValue(OBJC_PTR ptr, unsigned long* loc, unsigned long* length) {
	NSRange range = [(NSValue*)ptr rangeValue];
	*loc = range.location;
	*length = range.length;
}