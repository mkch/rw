#ifndef _RW_ARRAY_H
#define _RW_ARRAY_H

#include "../types.h"

unsigned long NSArray_count(OBJC_PTR ptr);
OBJC_PTR NSArray_objectAtIndex(OBJC_PTR ptr, unsigned long index);

#endif