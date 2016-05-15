#pragma once

#include "../types.h"

extern const NSUInteger VarNSWarningAlertStyle;
extern const NSUInteger VarNSInformationalAlertStyle;
extern const NSUInteger VarNSCriticalAlertStyle;

OBJC_PTR NSAlert_alloc();
OBJC_PTR NSAlert_alertWithError(OBJC_PTR error);
void NSAlert_layout(OBJC_PTR ptr);
NSUInteger NSAlert_alertStyle(OBJC_PTR ptr);
void NSAlert_setAlertStyle(OBJC_PTR ptr, NSUInteger style);
OBJC_PTR NSAlert_accessoryView(OBJC_PTR ptr);
void NSAlert_setAccessoryView(OBJC_PTR ptr, OBJC_PTR view);
bool NSAlert_showsHelp(OBJC_PTR ptr);
void NSAlert_setShowsHelp(OBJC_PTR ptr, bool show);
char* NSAlert_helpAnchor(OBJC_PTR ptr);
void NSAlert_setHelpAnchor(OBJC_PTR ptr, char* anchor);
OBJC_PTR NSAlert_delegate(OBJC_PTR ptr);
void NSAlert_setDelegate(OBJC_PTR ptr, OBJC_PTR delegate);
NSInteger NSAlert_runModal(OBJC_PTR ptr);
void NSAlert_beginSheetModalForWindow_completionHandler(OBJC_PTR ptr, OBJC_PTR sheetWindow, UINTPTR handlerId);
OBJC_PTR NSAlert_suppressionButton(OBJC_PTR ptr);
bool NSAlert_showsSuppressionButton(OBJC_PTR ptr);
void NSAlert_setShowsSuppressionButton(OBJC_PTR ptr, bool show);
char* NSAlert_informativeText(OBJC_PTR ptr);
void NSAlert_setInformativeText(OBJC_PTR ptr, char* text);
char* NSAlert_messageText(OBJC_PTR ptr);
void NSAlert_setMessageText(OBJC_PTR ptr, char* text);
OBJC_PTR NSAlert_icon(OBJC_PTR ptr);
void NSAlert_setIcon(OBJC_PTR ptr, OBJC_PTR icon);
OBJC_PTR NSAlert_buttons(OBJC_PTR ptr);
OBJC_PTR NSAlert_addButtonWithTitle(OBJC_PTR ptr, char* title);
OBJC_PTR NSAlert_window(OBJC_PTR ptr);