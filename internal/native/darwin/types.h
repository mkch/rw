#ifndef _RW_TYPES_DARWIN_H
#define _RW_TYPES_DARWIN_H

#include <stdint.h>
#include <stdbool.h>
// For CGFloat
#include <CoreGraphics/CoreGraphics.h>
// For NSInteger NSUInteger
#include <objc/NSObjCRuntime.h>

typedef void* OBJC_PTR;
typedef void* PVOID;
typedef uintptr_t UINTPTR;

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/Foundation/Miscellaneous/Foundation_DataTypes/#//apple_ref/c/tdef/NSUInteger
// typedef unsigned long NSUInteger;
// Discussion
// When building 32-bit applications, NSUInteger is a 32-bit unsigned integer. A 64-bit application treats NSUInteger as a 64-bit unsigned integer

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/Foundation/Miscellaneous/Foundation_DataTypes/#//apple_ref/c/tdef/NSInteger
// typedef long NSInteger;
// When building 32-bit applications, NSInteger is a 32-bit integer. A 64-bit application treats NSInteger as a 64-bit integer.

// https://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGGeometry/#//apple_ref/doc/constant_group/CGFloat_Informational_Macros
// CGFLOAT_IS_DOUBLE
// Indicates whether CGFloat is defined as a float or double type.
//
// Note from Kevin Yuan: CGFLOAT_IS_DOUBLE is 1 on my 64bit OS X.

#endif

