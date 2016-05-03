#ifndef _RW_VIEW_H
#define _RW_VIEW_H

#include "../types.h"
#include <CoreGraphics/CGGeometry.h>

OBJC_PTR NSView_initWithFrame(OBJC_PTR ptr, CGRect rect);
void NSView_addSubview(OBJC_PTR ptr, OBJC_PTR subView);
void NSView_removeFromSuperview(OBJC_PTR ptr);
OBJC_PTR NSView_superview(OBJC_PTR ptr);
OBJC_PTR NSView_window(OBJC_PTR ptr);
CGRect NSView_frame(OBJC_PTR ptr);
void NSView_setFrameSize(OBJC_PTR ptr, CGSize size);
void NSView_setFrameOrigin(OBJC_PTR ptr, CGPoint origin);
void NSView_setPostsFrameChangedNotifications(OBJC_PTR ptr, bool post);
// Get the NSViewFrameDidChangeNotification constant.
char* NSView_NSViewFrameDidChangeNotification();
bool NSView_isHidden(OBJC_PTR ptr);
void NSView_setHidden(OBJC_PTR ptr, bool hidden);
OBJC_PTR NSView_subviews(OBJC_PTR ptr);
void NSView_display(OBJC_PTR ptr);
void NSView_displayIfNeeded(OBJC_PTR ptr);
bool NSView_needsDisplay(OBJC_PTR ptr);
void NSView_setNeedsDisplay(OBJC_PTR ptr, bool needs);
OBJC_PTR NSView_nextKeyView(OBJC_PTR ptr);
void NSView_setNextKeyView(OBJC_PTR ptr, OBJC_PTR keyView);
OBJC_PTR NSView_previousKeyView(OBJC_PTR ptr);


OBJC_PTR RWFlippedView_alloc();
OBJC_PTR RWFlippedView_backgroundColor(OBJC_PTR ptr);
void RWFlippedView_setBackgroundColor(OBJC_PTR ptr, OBJC_PTR color);
void RWFlippedView_setAcceptFirstResponder(OBJC_PTR ptr, bool accept);

#endif