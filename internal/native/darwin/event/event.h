#ifndef _RW_EVENT_H
#define _RW_EVENT_H

#include "../types.h"

extern unsigned long VarNSLeftMouseDown;
extern unsigned long VarNSLeftMouseUp;
extern unsigned long VarNSRightMouseDown;
extern unsigned long VarNSRightMouseUp;
extern unsigned long VarNSMouseMoved;
extern unsigned long VarNSLeftMouseDragged;
extern unsigned long VarNSRightMouseDragged;
extern unsigned long VarNSMouseEntered;
extern unsigned long VarNSMouseExited;
extern unsigned long VarNSKeyDown;
extern unsigned long VarNSKeyUp;
extern unsigned long VarNSFlagsChanged;
extern unsigned long VarNSAppKitDefined;
extern unsigned long VarNSSystemDefined;
extern unsigned long VarNSApplicationDefined;
extern unsigned long VarNSPeriodic;
extern unsigned long VarNSCursorUpdate;
extern unsigned long VarNSScrollWheel;
extern unsigned long VarNSTabletPoint;
extern unsigned long VarNSTabletProximity;
extern unsigned long VarNSOtherMouseDown;
extern unsigned long VarNSOtherMouseUp;
extern unsigned long VarNSOtherMouseDragged;
extern unsigned long VarNSEventTypeGesture;
extern unsigned long VarNSEventTypeMagnify;
extern unsigned long VarNSEventTypeSwipe;
extern unsigned long VarNSEventTypeRotate;
extern unsigned long VarNSEventTypeBeginGesture;
extern unsigned long VarNSEventTypeEndGesture;
extern unsigned long VarNSEventTypeSmartMagnify;
extern unsigned long VarNSEventTypeQuickLook;
extern unsigned long VarNSEventTypePressure;

extern unsigned long VarNSLeftMouseDownMask;
extern unsigned long VarNSLeftMouseUpMask;
extern unsigned long VarNSRightMouseDownMask;
extern unsigned long VarNSRightMouseUpMask;
extern unsigned long VarNSMouseMovedMask;
extern unsigned long VarNSLeftMouseDraggedMask;
extern unsigned long VarNSRightMouseDraggedMask;
extern unsigned long VarNSMouseEnteredMask;
extern unsigned long VarNSMouseExitedMask;
extern unsigned long VarNSKeyDownMask;
extern unsigned long VarNSKeyUpMask;
extern unsigned long VarNSFlagsChangedMask;
extern unsigned long VarNSAppKitDefinedMask;
extern unsigned long VarNSSystemDefinedMask;
extern unsigned long VarNSApplicationDefinedMask;
extern unsigned long VarNSPeriodicMask;
extern unsigned long VarNSCursorUpdateMask;
extern unsigned long VarNSScrollWheelMask;
extern unsigned long VarNSTabletPointMask;
extern unsigned long VarNSTabletProximityMask;
extern unsigned long VarNSOtherMouseDownMask;
extern unsigned long VarNSOtherMouseUpMask;
extern unsigned long VarNSOtherMouseDraggedMask;
extern unsigned long VarNSEventMaskGesture;
extern unsigned long VarNSEventMaskMagnify;
extern unsigned long VarNSEventMaskSwipe;
extern unsigned long VarNSEventMaskRotate;
extern unsigned long VarNSEventMaskBeginGesture;
extern unsigned long VarNSEventMaskEndGesture;
extern unsigned long VarNSEventMaskSmartMagnify;
extern unsigned long VarNSEventMaskPressure;
extern unsigned long VarNSAnyEventMask;

extern unsigned long VarNSAlphaShiftKeyMask;
extern unsigned long VarNSShiftKeyMask;
extern unsigned long VarNSControlKeyMask;
extern unsigned long VarNSAlternateKeyMask;
extern unsigned long VarNSCommandKeyMask;
extern unsigned long VarNSNumericPadKeyMask;
extern unsigned long VarNSHelpKeyMask;
extern unsigned long VarNSFunctionKeyMask;
extern unsigned long VarNSDeviceIndependentModifierFlagsMask;

OBJC_PTR NSEvent_otherEventWithType_location_modifierFlags_timestamp_windowNumber_context_subtype_data1_data2(
	NSUInteger type, double locationX, double locationY, unsigned int modifierFlags, double timestamp, NSInteger windowNumber, OBJC_PTR context, short subtype, NSInteger data1, NSInteger data2);
unsigned long long NSEvent_type(OBJC_PTR ptr);
OBJC_PTR NSEvent_window(OBJC_PTR ptr);

#endif
