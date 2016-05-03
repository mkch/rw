#import <Foundation/NSAutoreleasePool.h>
#import "autoreleasepool.h"

OBJC_PTR NSAutoreleasePool_alloc() {
	return [NSAutoreleasePool alloc];
}