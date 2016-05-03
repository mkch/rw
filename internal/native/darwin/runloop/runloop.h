#ifndef _RW_RUNLOOP_H
#define _RW_RUNLOOP_H

#include "../types.h"

OBJC_PTR getNSRunLoopCommonModes();

OBJC_PTR getNSDefaultRunLoopMode();
OBJC_PTR getNSConnectionReplyMode();
OBJC_PTR getNSModalPanelRunLoopMode();
OBJC_PTR getNSEventTrackingRunLoopMode();

OBJC_PTR NSRunLoop_currentRunLoop();
OBJC_PTR NSRunloop_currentMode(OBJC_PTR ptr);

#endif
