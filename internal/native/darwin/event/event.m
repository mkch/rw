#import <AppKit/NSResponder.h>
//#import <Foundation/NSProcessInfo.h>
//#import <Foundation/NSDate.h>
#import "event.h"

unsigned long VarNSLeftMouseDown = NSLeftMouseDown;
unsigned long VarNSLeftMouseUp = NSLeftMouseUp;
unsigned long VarNSRightMouseDown = NSRightMouseDown;
unsigned long VarNSRightMouseUp = NSRightMouseUp;
unsigned long VarNSMouseMoved = NSMouseMoved;
unsigned long VarNSLeftMouseDragged = NSLeftMouseDragged;
unsigned long VarNSRightMouseDragged = NSRightMouseDragged;
unsigned long VarNSMouseEntered = NSMouseEntered;
unsigned long VarNSMouseExited = NSMouseExited;
unsigned long VarNSKeyDown = NSKeyDown;
unsigned long VarNSKeyUp = NSKeyUp;
unsigned long VarNSFlagsChanged = NSFlagsChanged;
unsigned long VarNSAppKitDefined = NSAppKitDefined;
unsigned long VarNSSystemDefined = NSSystemDefined;
unsigned long VarNSApplicationDefined = NSApplicationDefined;
unsigned long VarNSPeriodic = NSPeriodic;
unsigned long VarNSCursorUpdate = NSCursorUpdate;
unsigned long VarNSScrollWheel = NSScrollWheel;
unsigned long VarNSTabletPoint = NSTabletPoint;
unsigned long VarNSTabletProximity = NSTabletProximity;
unsigned long VarNSOtherMouseDown = NSOtherMouseDown;
unsigned long VarNSOtherMouseUp = NSOtherMouseUp;
unsigned long VarNSOtherMouseDragged = NSOtherMouseDragged;
unsigned long VarNSEventTypeGesture = NSEventTypeGesture;
unsigned long VarNSEventTypeMagnify = NSEventTypeMagnify;
unsigned long VarNSEventTypeSwipe = NSEventTypeSwipe;
unsigned long VarNSEventTypeRotate = NSEventTypeRotate;
unsigned long VarNSEventTypeBeginGesture = NSEventTypeBeginGesture;
unsigned long VarNSEventTypeEndGesture = NSEventTypeEndGesture;
unsigned long VarNSEventTypeSmartMagnify = NSEventTypeSmartMagnify;
unsigned long VarNSEventTypeQuickLook = NSEventTypeQuickLook;
unsigned long VarNSEventTypePressure = NSEventTypePressure;

unsigned long VarNSLeftMouseDownMask = NSLeftMouseDownMask;
unsigned long VarNSLeftMouseUpMask = NSLeftMouseUpMask;
unsigned long VarNSRightMouseDownMask = NSRightMouseDownMask;
unsigned long VarNSRightMouseUpMask = NSRightMouseUpMask;
unsigned long VarNSMouseMovedMask = NSMouseMovedMask;
unsigned long VarNSLeftMouseDraggedMask = NSLeftMouseDraggedMask;
unsigned long VarNSRightMouseDraggedMask = NSRightMouseDraggedMask;
unsigned long VarNSMouseEnteredMask = NSMouseEnteredMask;
unsigned long VarNSMouseExitedMask = NSMouseExitedMask;
unsigned long VarNSKeyDownMask = NSKeyDownMask;
unsigned long VarNSKeyUpMask = NSKeyUpMask;
unsigned long VarNSFlagsChangedMask = NSFlagsChangedMask;
unsigned long VarNSAppKitDefinedMask = NSAppKitDefinedMask;
unsigned long VarNSSystemDefinedMask = NSSystemDefinedMask;
unsigned long VarNSApplicationDefinedMask = NSApplicationDefinedMask;
unsigned long VarNSPeriodicMask = NSPeriodicMask;
unsigned long VarNSCursorUpdateMask = NSCursorUpdateMask;
unsigned long VarNSScrollWheelMask = NSScrollWheelMask;
unsigned long VarNSTabletPointMask = NSTabletPointMask;
unsigned long VarNSTabletProximityMask = NSTabletProximityMask;
unsigned long VarNSOtherMouseDownMask = NSOtherMouseDownMask;
unsigned long VarNSOtherMouseUpMask = NSOtherMouseUpMask;
unsigned long VarNSOtherMouseDraggedMask = NSOtherMouseDraggedMask;
unsigned long VarNSEventMaskGesture = NSEventMaskGesture;
unsigned long VarNSEventMaskMagnify = NSEventMaskMagnify;
unsigned long VarNSEventMaskSwipe = NSEventMaskSwipe;
unsigned long VarNSEventMaskRotate = NSEventMaskRotate;
unsigned long VarNSEventMaskBeginGesture = NSEventMaskBeginGesture;
unsigned long VarNSEventMaskEndGesture = NSEventMaskEndGesture;
unsigned long VarNSEventMaskSmartMagnify = NSEventMaskSmartMagnify;
unsigned long VarNSEventMaskPressure = NSEventMaskPressure;
unsigned long VarNSAnyEventMask = NSAnyEventMask;

unsigned long VarNSAlphaShiftKeyMask = NSAlphaShiftKeyMask;
unsigned long VarNSShiftKeyMask = NSShiftKeyMask;
unsigned long VarNSControlKeyMask = NSControlKeyMask;
unsigned long VarNSAlternateKeyMask = NSAlternateKeyMask;
unsigned long VarNSCommandKeyMask = NSCommandKeyMask;
unsigned long VarNSNumericPadKeyMask = NSNumericPadKeyMask;
unsigned long VarNSHelpKeyMask = NSHelpKeyMask;
unsigned long VarNSFunctionKeyMask = NSFunctionKeyMask;
unsigned long VarNSDeviceIndependentModifierFlagsMask = NSDeviceIndependentModifierFlagsMask;

OBJC_PTR NSEvent_otherEventWithType_location_modifierFlags_timestamp_windowNumber_context_subtype_data1_data2(
	NSUInteger type, double locationX, double locationY, unsigned int modifierFlags, double timestamp, NSInteger windowNumber, OBJC_PTR context, short subtype, NSInteger data1, NSInteger data2) {
	return [NSEvent otherEventWithType:type
		location:CGPointMake(locationX, locationY)
		modifierFlags:modifierFlags
		timestamp: timestamp // [[NSProcessInfo processInfo]systemUptime]
		windowNumber: windowNumber
		context: context
		subtype: subtype
		data1: data1
		data2:data2];
}

unsigned long long NSEvent_type(OBJC_PTR ptr) {
	return [(NSEvent*)ptr type];
}

OBJC_PTR NSEvent_window(OBJC_PTR ptr) {
	return [(NSEvent*)ptr window];
}

