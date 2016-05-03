#ifndef _RW_SCREEN_H
#define _RW_SCREEN_H

#include "../types.h"
#include <CoreGraphics/CGGeometry.h>

OBJC_PTR NSScreen_mainScreen();
CGRect NSScreen_visibleFrame(OBJC_PTR ptr);
CGRect NSScreen_frame(OBJC_PTR ptr);

#endif
