package event

//#include "event.h"
import "C"

import (
	"github.com/mkch/rw/native"
)

type NSEventType uint

var (
	NSLeftMouseDown         = NSEventType(C.VarNSLeftMouseDown)
	NSLeftMouseUp           = NSEventType(C.VarNSLeftMouseUp)
	NSRightMouseDown        = NSEventType(C.VarNSRightMouseDown)
	NSRightMouseUp          = NSEventType(C.VarNSRightMouseUp)
	NSMouseMoved            = NSEventType(C.VarNSMouseMoved)
	NSLeftMouseDragged      = NSEventType(C.VarNSLeftMouseDragged)
	NSRightMouseDragged     = NSEventType(C.VarNSRightMouseDragged)
	NSMouseEntered          = NSEventType(C.VarNSMouseEntered)
	NSMouseExited           = NSEventType(C.VarNSMouseExited)
	NSKeyDown               = NSEventType(C.VarNSKeyDown)
	NSKeyUp                 = NSEventType(C.VarNSKeyUp)
	NSFlagsChanged          = NSEventType(C.VarNSFlagsChanged)
	NSAppKitDefined         = NSEventType(C.VarNSAppKitDefined)
	NSSystemDefined         = NSEventType(C.VarNSSystemDefined)
	NSApplicationDefined    = NSEventType(C.VarNSApplicationDefined)
	NSPeriodic              = NSEventType(C.VarNSPeriodic)
	NSCursorUpdate          = NSEventType(C.VarNSCursorUpdate)
	NSScrollWheel           = NSEventType(C.VarNSScrollWheel)
	NSTabletPoint           = NSEventType(C.VarNSTabletPoint)
	NSTabletProximity       = NSEventType(C.VarNSTabletProximity)
	NSOtherMouseDown        = NSEventType(C.VarNSOtherMouseDown)
	NSOtherMouseUp          = NSEventType(C.VarNSOtherMouseUp)
	NSOtherMouseDragged     = NSEventType(C.VarNSOtherMouseDragged)
	NSEventTypeGesture      = NSEventType(C.VarNSEventTypeGesture)
	NSEventTypeMagnify      = NSEventType(C.VarNSEventTypeMagnify)
	NSEventTypeSwipe        = NSEventType(C.VarNSEventTypeSwipe)
	NSEventTypeRotate       = NSEventType(C.VarNSEventTypeRotate)
	NSEventTypeBeginGesture = NSEventType(C.VarNSEventTypeBeginGesture)
	NSEventTypeEndGesture   = NSEventType(C.VarNSEventTypeEndGesture)
	NSEventTypeSmartMagnify = NSEventType(C.VarNSEventTypeSmartMagnify)
	NSEventTypeQuickLook    = NSEventType(C.VarNSEventTypeQuickLook)
	NSEventTypePressure     = NSEventType(C.VarNSEventTypePressure)
)

type NSEventMask uint

var (
	NSLeftMouseDownMask      = NSEventMask(C.VarNSLeftMouseDownMask)
	NSLeftMouseUpMask        = NSEventMask(C.VarNSLeftMouseUpMask)
	NSRightMouseDownMask     = NSEventMask(C.VarNSRightMouseDownMask)
	NSRightMouseUpMask       = NSEventMask(C.VarNSRightMouseUpMask)
	NSMouseMovedMask         = NSEventMask(C.VarNSMouseMovedMask)
	NSLeftMouseDraggedMask   = NSEventMask(C.VarNSLeftMouseDraggedMask)
	NSRightMouseDraggedMask  = NSEventMask(C.VarNSRightMouseDraggedMask)
	NSMouseEnteredMask       = NSEventMask(C.VarNSMouseEnteredMask)
	NSMouseExitedMask        = NSEventMask(C.VarNSMouseExitedMask)
	NSKeyDownMask            = NSEventMask(C.VarNSKeyDownMask)
	NSKeyUpMask              = NSEventMask(C.VarNSKeyUpMask)
	NSFlagsChangedMask       = NSEventMask(C.VarNSFlagsChangedMask)
	NSAppKitDefinedMask      = NSEventMask(C.VarNSAppKitDefinedMask)
	NSSystemDefinedMask      = NSEventMask(C.VarNSSystemDefinedMask)
	NSApplicationDefinedMask = NSEventMask(C.VarNSApplicationDefinedMask)
	NSPeriodicMask           = NSEventMask(C.VarNSPeriodicMask)
	NSCursorUpdateMask       = NSEventMask(C.VarNSCursorUpdateMask)
	NSScrollWheelMask        = NSEventMask(C.VarNSScrollWheelMask)
	NSTabletPointMask        = NSEventMask(C.VarNSTabletPointMask)
	NSTabletProximityMask    = NSEventMask(C.VarNSTabletProximityMask)
	NSOtherMouseDownMask     = NSEventMask(C.VarNSOtherMouseDownMask)
	NSOtherMouseUpMask       = NSEventMask(C.VarNSOtherMouseUpMask)
	NSOtherMouseDraggedMask  = NSEventMask(C.VarNSOtherMouseDraggedMask)
	NSEventMaskGesture       = NSEventMask(C.VarNSEventMaskGesture)
	NSEventMaskMagnify       = NSEventMask(C.VarNSEventMaskMagnify)
	NSEventMaskSwipe         = NSEventMask(C.VarNSEventMaskSwipe)
	NSEventMaskRotate        = NSEventMask(C.VarNSEventMaskRotate)
	NSEventMaskBeginGesture  = NSEventMask(C.VarNSEventMaskBeginGesture)
	NSEventMaskEndGesture    = NSEventMask(C.VarNSEventMaskEndGesture)
	NSEventMaskSmartMagnify  = NSEventMask(C.VarNSEventMaskSmartMagnify)
	NSEventMaskPressure      = NSEventMask(C.VarNSEventMaskPressure)
	NSAnyEventMask           = NSEventMask(C.VarNSAnyEventMask)
)

type NSEventModifierFlags uint

var (
	NSAlphaShiftKeyMask                  = NSEventModifierFlags(C.VarNSAlphaShiftKeyMask)
	NSShiftKeyMask                       = NSEventModifierFlags(C.VarNSShiftKeyMask)
	NSControlKeyMask                     = NSEventModifierFlags(C.VarNSControlKeyMask)
	NSAlternateKeyMask                   = NSEventModifierFlags(C.VarNSAlternateKeyMask)
	NSCommandKeyMask                     = NSEventModifierFlags(C.VarNSCommandKeyMask)
	NSNumericPadKeyMask                  = NSEventModifierFlags(C.VarNSNumericPadKeyMask)
	NSHelpKeyMask                        = NSEventModifierFlags(C.VarNSHelpKeyMask)
	NSFunctionKeyMask                    = NSEventModifierFlags(C.VarNSFunctionKeyMask)
	NSDeviceIndependentModifierFlagsMask = NSEventModifierFlags(C.VarNSDeviceIndependentModifierFlagsMask)
)

func NSEvent_otherEventWithType_location_modifierFlags_timestamp_windowNumber_context_subtype_data1_data2(
	eventType NSEventType, locationX, locationY float64, modifierFlags NSEventModifierFlags, timestamp float64, windowNumber int, context native.Handle, subtype int16, data1, data2 int) native.Handle {
	return native.Handle(
		C.NSEvent_otherEventWithType_location_modifierFlags_timestamp_windowNumber_context_subtype_data1_data2(
			C.NSUInteger(eventType), C.double(locationX), C.double(locationY), C.uint(modifierFlags), C.double(timestamp), C.NSInteger(windowNumber), C.OBJC_PTR(context), C.short(subtype), C.NSInteger(data1), C.NSInteger(data2)))
}

func NSEvent_type(evt native.Handle) NSEventType {
	return NSEventType(C.NSEvent_type(C.OBJC_PTR(evt)))
}

func NSEvent_window(evt native.Handle) native.Handle {
	return native.Handle(C.NSEvent_window(C.OBJC_PTR(evt)))
}
