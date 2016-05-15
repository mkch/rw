package runloop

//#include "runloop.h"
import "C"

import (
	"github.com/mkch/rw/native"
)

var (
	NSDefaultRunLoopMode       native.Handle = native.Handle(C.getNSDefaultRunLoopMode())
	NSRunLoopCommonModes                     = native.Handle(C.getNSRunLoopCommonModes())
	NSConnectionReplyMode                    = native.Handle(C.getNSConnectionReplyMode())
	NSModalPanelRunLoopMode                  = native.Handle(C.getNSModalPanelRunLoopMode())
	NSEventTrackingRunLoopMode               = native.Handle(C.getNSEventTrackingRunLoopMode())
)

func NSRunLoop_currentRunLoop() native.Handle {
	return native.Handle(C.NSRunLoop_currentRunLoop())
}

func NSRunloop_currentMode(rl native.Handle) native.Handle {
	return native.Handle(C.NSRunloop_currentMode(C.OBJC_PTR(rl)))
}
