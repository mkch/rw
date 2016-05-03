#ifndef _RW_TEXTVIEW_H
#define _RW_TEXTVIEW_H

#include "../types.h"

void NSTextView_setSelectedRange(OBJC_PTR ptr, unsigned long loc, unsigned long len);
OBJC_PTR/*(NSArray*)*/ NSTextView_selectedRanges(OBJC_PTR ptr);

#endif