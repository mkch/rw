package util

import (
	"github.com/mkch/rw/internal/native/darwin/app"
	"github.com/mkch/rw/internal/native/darwin/date"
	"github.com/mkch/rw/internal/native/darwin/event"
	"github.com/mkch/rw/internal/native/darwin/runloop"
)

type RunloopShortCircuiter struct {
	complated bool
}

// Run runs a short circuit event loop here to wait for Stop.
func (r *RunloopShortCircuiter) Run() {
	for !r.complated {
		nextEvent := app.NSApplication_nextEventMatchingMask_untilDate_inMode_dequeue(app.NSApp(),
			event.NSAnyEventMask,         // mask
			date.NSDate_distantFuture(),  // expiration.
			runloop.NSDefaultRunLoopMode, // mode.
			true, // flag
		)
		if nextEvent == 0 {
			break
		}
		app.NSApplication_sendEvent(app.NSApp(), nextEvent)
	}
	r.complated = false
}

// Stop stops the event loop short circuit and make Run return.
func (r *RunloopShortCircuiter) Stop() {
	r.complated = true
	// Send an fake(empty) event to wake up the event loop below.
	//
	// Do not save this fake event in a variable and test whether the next event equals this event in Run.
	// NSEvent conforms to NSCopying, may be copied. Don't make pointer comparison.
	app.NSApplication_postEvent_atStart(app.NSApp(), event.NSEvent_otherEventWithType_location_modifierFlags_timestamp_windowNumber_context_subtype_data1_data2(
		event.NSApplicationDefined, // type
		0, 0, // location
		0,    // flags
		0,    // time
		0,    // windowNumber
		0,    // context
		0,    // subtype
		0, 0, // data1 & data2
	), false)
}
