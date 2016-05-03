#ifndef _RW_COLOR_H
#define _RW_COLOR_H

#include <CoreGraphics/CGBase.h>
#include "../types.h"

OBJC_PTR NSColor_colorWithRGB(CGFloat r, CGFloat g, CGFloat b, CGFloat a);
void NSColor_getRGBA(OBJC_PTR ptr, CGFloat* r, CGFloat *g, CGFloat* b, CGFloat* a);

#endif