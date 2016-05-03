#ifndef _RW_APP_H
#define _RW_APP_H

#include "../types.h"

OBJC_PTR getNSApp();
void NSApplication_terminate(OBJC_PTR ptr, OBJC_PTR sender);
void NSApplication_stop(OBJC_PTR ptr, OBJC_PTR sender);
void NSApplication_setMainMenu(OBJC_PTR ptr, OBJC_PTR menu);
OBJC_PTR NSApplication_mainMenu(OBJC_PTR ptr);
UINTPTR NSApplication_runModalForWindow(OBJC_PTR ptr, OBJC_PTR win);
extern UINTPTR NSApplication_NSModalResponseStop, NSApplication_NSModalResponseAbort, NSApplication_NSModalResponseContinue;
void NSApplication_stopModalWithCode(OBJC_PTR ptr, UINTPTR code);
void NSApplication_abortModal(OBJC_PTR ptr);
OBJC_PTR NSApplication_modalWindow(OBJC_PTR ptr);

extern long VarNSApplicationActivationPolicyRegular;
extern long VarNSApplicationActivationPolicyAccessory;
extern long VarNSApplicationActivationPolicyProhibited;

void NSApplication_setActivationPolicy(OBJC_PTR ptr, long ploicy);
void NSApplication_run(OBJC_PTR ptr);
OBJC_PTR NSApplication_windows(OBJC_PTR ptr);

void NSApplication_sendEvent(OBJC_PTR ptr, OBJC_PTR event);
void NSApplication_postEvent_atStart(OBJC_PTR ptr, OBJC_PTR event, bool flag);
OBJC_PTR NSApplication_nextEventMatchingMask_untilDate_inMode_dequeue(OBJC_PTR ptr, UINTPTR mask, OBJC_PTR expiration, OBJC_PTR mode, bool flag);

// RWApp

OBJC_PTR RWApp_sharedApplication();
void RWApp_superSendEvent(OBJC_PTR ptr, OBJC_PTR event);

#endif
