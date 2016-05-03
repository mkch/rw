package app

//#include "../types.h"
//#include "app.h"
import "C"

import (
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/internal/native/darwin/event"
)

// disabledWindows contains all "disabled" windows.
// Cocoa does not provided a way to disable a window. The trick we use
// is to put all windows that should be disabled in a map and do not send
// input events to them. See also goAppSendEvent below.
var disabledWindows = make(map[native.Handle] bool)

// EnableWindow enables or disables a window.
// Cocoa does not provided a way to disable a window. See disabledWindows above
// and goAppSendEvent below for details of the trick.
func EnableWindow(w native.Handle, enable bool) {
	if enable {
		delete(disabledWindows, w)
	} else {
		disabledWindows[w] = true
	}
}

// WindowEnabled returns whether a window is enabled.
// Cocoa does not provided a way to disable a window. See disabledWindows above
// and goAppSendEvent below for details of the trick.
func WindowEnabled(w native.Handle) bool {
	return disabledWindows[w]
}


//export goAppSendEvent
func goAppSendEvent(theApp, evt C.OBJC_PTR) {
	switch event.NSEvent_type(native.Handle(evt)) {
	case event.NSPeriodic, event.NSAppKitDefined, event.NSSystemDefined, event.NSApplicationDefined:
	default: // Input envnts.
		if disabledWindows[event.NSEvent_window(native.Handle(evt))] {
			return // Do not send input events to "disabled" windows. See also EnableWindow above.
		}
	}
	C.RWApp_superSendEvent(C.OBJC_PTR(theApp), C.OBJC_PTR(evt))
}

func RWApp_sharedApplication() native.Handle {
	return native.Handle(C.RWApp_sharedApplication())
}

func RWApp_superSendEvent(a, event native.Handle) {
	C.RWApp_superSendEvent(C.OBJC_PTR(a), C.OBJC_PTR(event))
}