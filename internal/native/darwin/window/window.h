#ifndef _RW_WINDOW_H
#define _RW_WINDOW_H

#include "../types.h"
#include <CoreGraphics/CGGeometry.h>

OBJC_PTR NSWindow_alloc();
OBJC_PTR RWWindow_alloc();
OBJC_PTR NSWindow_initWithContentRect_styleMask_backing_defer_screen(OBJC_PTR ptr, 
	CGRect rect, UINTPTR styleMask, UINTPTR bufferingType, bool deferCreation, OBJC_PTR screen);
void NSWindow_makeKeyAndOrderFront(OBJC_PTR ptr);
char* NSWindow_title(OBJC_PTR ptr);
void NSWindow_setTitle(OBJC_PTR ptr, char* title);
void NSWindow_cascadeTopLeftFromPoint(OBJC_PTR ptr, CGPoint topLeft);
void NSWindow_setFrameTopLeftPoint(OBJC_PTR ptr, CGPoint topLeft);
OBJC_PTR NSWindow_contentView(OBJC_PTR ptr);
void NSWindow_setContentView(OBJC_PTR ptr, OBJC_PTR view);
void NSWindow_setDelegate(OBJC_PTR ptr, OBJC_PTR delegate);
OBJC_PTR NSWindow_delegate(OBJC_PTR ptr);
CGRect NSWindow_frameRectForContentRect_styleMask(CGRect windowContentRect, UINTPTR windowStyle);
CGRect NSWindow_contentRectForFrameRect_styleMask(CGRect windowFrameRect, UINTPTR windowStyle);
CGRect NSWindow_frame(OBJC_PTR ptr);
void NSWindow_setFrameDisplay(OBJC_PTR ptr, CGRect frame, bool displayViews);
bool NSWindow_makeFirstResponder(OBJC_PTR ptr, OBJC_PTR responder);
void NSWindow_close(OBJC_PTR ptr);
void NSWindow_performClose(OBJC_PTR ptr, OBJC_PTR sender);
void NSWindow_center(OBJC_PTR ptr);
void NSWindow_beginSheet_completionHandler(OBJC_PTR ptr, OBJC_PTR win, UINTPTR handlerId);
void NSWindow_endSheet_returnCode(OBJC_PTR ptr, OBJC_PTR win, int returnCode);
OBJC_PTR NSWindow_attachedSheet(OBJC_PTR ptr);
bool NSWindow_isSheet(OBJC_PTR ptr);
OBJC_PTR NSWindow_sheetParent(OBJC_PTR ptr);
extern int VarNSModalResponseStop;
extern int VarNSModalResponseAbort;
extern int VarNSModalResponseContinue;
void NSWindow_discardEventsMatchingMask_beforeEvent(OBJC_PTR ptr, unsigned long eventMask, OBJC_PTR lastEvent);
void NSWindow_sendEvent(OBJC_PTR ptr, OBJC_PTR event);
void NSWindow_display(OBJC_PTR ptr);
void NSWindow_displayIfNeeded(OBJC_PTR ptr);
bool NSWindow_isVisible(OBJC_PTR ptr);
void NSWindow_orderOut(OBJC_PTR ptr);
void NSWindow_orderBack(OBJC_PTR ptr);
void NSWindow_orderFront(OBJC_PTR ptr);
void NSWindow_orderFrontRegardless(OBJC_PTR ptr);
extern long VarNSWindowAbove;
extern long VarNSWindowBelow;
extern long VarNSWindowOut;
void NSWindow_orderWindow_relativeTo(OBJC_PTR ptr, long orderingMode, long otherWindowNumber);
long NSWindow_level(OBJC_PTR ptr);
void NSWindow_setLevel(OBJC_PTR ptr, long level);
OBJC_PTR NSWindow_firstResponder(OBJC_PTR ptr);
OBJC_PTR NSWindow_screen(OBJC_PTR ptr);

bool RWWindow_enabled(OBJC_PTR ptr);
void RWWindow_setEnabled(OBJC_PTR ptr, bool enabled);

#endif
