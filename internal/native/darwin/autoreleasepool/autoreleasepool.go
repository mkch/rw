package autoreleasepool

//#include "autoreleasepool.h"
import "C"

import (
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/internal/native/darwin/object"
)

func NSAutoreleasePool_alloc() native.Handle {
	return native.Handle(C.NSAutoreleasePool_alloc())
}

// Run calls f in a Cocoa @autoreleasepool block (see https://developer.apple.com/library/mac/documentation/Cocoa/Reference/Foundation/Classes/NSAutoreleasePool_Class/).
//  // go code
//  autoreleasepool.Run(f)
// is equivalent to
//  // objective-c code
//  @autoreleasepool {
//    f();	
//  }
func Run(f func()) {
	pool := object.NSObject_init(NSAutoreleasePool_alloc())
	defer object.NSObject_release(pool)
	f()
}