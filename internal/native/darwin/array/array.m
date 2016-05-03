#import <Foundation/NSArray.h>
#import "array.h"

unsigned long NSArray_count(OBJC_PTR ptr) {
	return [(NSArray*)ptr count];
}

OBJC_PTR NSArray_objectAtIndex(OBJC_PTR ptr, unsigned long index) {
	return [(NSArray*)ptr objectAtIndex:index];
}