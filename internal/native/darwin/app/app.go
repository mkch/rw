package app

//#include "app.h"
import "C"

import (
	"github.com/mkch/rw/internal/native/darwin/event"
	"github.com/mkch/rw/native"
)

func NSApp() native.Handle {
	return native.Handle(C.getNSApp())
}

func NSApplication_stop(app, sender native.Handle) {
	C.NSApplication_stop(C.OBJC_PTR(app), C.OBJC_PTR(sender))
}

func NSApplication_terminate(app, sender native.Handle) {
	C.NSApplication_terminate(C.OBJC_PTR(app), C.OBJC_PTR(sender))
}

func NSApplication_setMainMenu(app, menu native.Handle) {
	C.NSApplication_setMainMenu(C.OBJC_PTR(app), C.OBJC_PTR(menu))
}

func NSApplication_mainMenu(app native.Handle) native.Handle {
	return native.Handle(C.NSApplication_mainMenu(C.OBJC_PTR(app)))
}

func NSApplication_runModalForWindow(app, win native.Handle) uintptr {
	return uintptr(C.NSApplication_runModalForWindow(C.OBJC_PTR(app), C.OBJC_PTR(win)))
}

func NSApplication_abortModal(app native.Handle) {
	C.NSApplication_abortModal(C.OBJC_PTR(app))
}

func NSApplication_stopModalWithCode(app native.Handle, code uintptr) {
	C.NSApplication_stopModalWithCode(C.OBJC_PTR(app), C.UINTPTR(code))
}

var (
	ModalResponseStop     = uintptr(C.NSApplication_NSModalResponseStop)
	ModalResponseAbort    = uintptr(C.NSApplication_NSModalResponseAbort)
	ModalResponseContinue = uintptr(C.NSApplication_NSModalResponseContinue)
)

func NSApplication_modalWindow(app native.Handle) native.Handle {
	return native.Handle(C.NSApplication_modalWindow(C.OBJC_PTR(app)))
}

func NSApplication_run(app native.Handle) {
	C.NSApplication_run(C.OBJC_PTR(app))
}

type NSApplicationActivationPolicy int

var (
	NSApplicationActivationPolicyRegular    = NSApplicationActivationPolicy(C.VarNSApplicationActivationPolicyRegular)
	NSApplicationActivationPolicyAccessory  = NSApplicationActivationPolicy(C.VarNSApplicationActivationPolicyAccessory)
	NSApplicationActivationPolicyProhibited = NSApplicationActivationPolicy(C.VarNSApplicationActivationPolicyProhibited)
)

func NSApplication_setActivationPolicy(app native.Handle, policy NSApplicationActivationPolicy) {
	C.NSApplication_setActivationPolicy(C.OBJC_PTR(app), C.long(policy))
}

func NSApplication_windows(app native.Handle) native.Handle {
	return native.Handle(C.NSApplication_windows(C.OBJC_PTR(app)))
}

func NSApplication_sendEvent(app, event native.Handle) {
	C.NSApplication_sendEvent(C.OBJC_PTR(app), C.OBJC_PTR(event))
}

func NSApplication_postEvent_atStart(app, event native.Handle, flag bool) {
	C.NSApplication_postEvent_atStart(C.OBJC_PTR(app), C.OBJC_PTR(event), C.bool(flag))
}

func NSApplication_nextEventMatchingMask_untilDate_inMode_dequeue(app native.Handle, mask event.NSEventMask, expiration, mode native.Handle, flag bool) native.Handle {
	return native.Handle(C.NSApplication_nextEventMatchingMask_untilDate_inMode_dequeue(C.OBJC_PTR(app), C.UINTPTR(mask), C.OBJC_PTR(expiration), C.OBJC_PTR(mode), C.bool(flag)))
}
