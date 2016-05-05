package event_test

import (
	"github.com/mkch/rw/event"
	"testing"
)

func TestZeroHook(t *testing.T) {
	var hook event.HookChain
	var sender interface{} = 1

	if hook.HasCallback() {
		t.Errorf("Empty Hook has callbacks")
	}

	if ret := hook.Call(newSimpleEvent(sender, 0)); ret != false {
		t.Errorf("Call of zero value Hook returned %v. Should be false", ret)
	}

	testHookUnhook(&hook, t)
}

func TestNewHook(t *testing.T) {
	var sender interface{} = 1
	var eventSent = newSimpleEvent(sender, 11)
	var defValue = true

	var hook = &event.HookChain{DefaultReturnValue: defValue}

	if hook.HasCallback() {
		t.Errorf("Empty Hook has callback")
	}

	if ret := hook.Call(eventSent); ret != defValue {
		t.Errorf("Wrong return value of hook.Call. Want %v, got %v", defValue, ret)
	}

	testHookUnhook(&event.HookChain{}, t)
}

func testHookUnhook(hook *event.HookChain, t *testing.T) {
	if hook.HasCallback() {
		t.Errorf("Empty Hook has callback")
	}

	var sender interface{} = 1
	var eventValue = 11
	var eventSent = newSimpleEvent(sender, eventValue)

	var expectedEventValueInCallback1 = eventSent.value
	var expectedNextHookRet = true
	var calledCallback1 bool = false
	item1 := hook.AddHook(func(event event.Event, nextHook event.Handler) bool {
			calledCallback1 = true
			if event != eventSent {
				t.Errorf("Wrong event in hook callback. Want %v, got %v", eventSent, event)
			}
			if s := event.Sender(); s != sender {
				t.Errorf("Wrong sender of Event in hook callback. Want %v, got %v", sender, s)
			}
			if v := event.(*simpleEvent).Value(); v != expectedEventValueInCallback1 {
				t.Errorf("Wrong value of simpleEvent in hook callback. Want %v, got %v", expectedEventValueInCallback1, v)
			}
			if ret := nextHook(event); ret != false {
				t.Errorf("Wrong return value of nextHook. Want %v, got %v", false, ret)
			}
			return expectedNextHookRet
		})
	if !hook.HasCallback() {
		t.Errorf("Non-empty Hook does not have callbacks")
	}
	if ret := hook.Call(eventSent); ret != expectedNextHookRet {
		t.Errorf("Wrong return value of hook.Call. Want %v, got %v", expectedNextHookRet, ret)
	}

	if !calledCallback1 {
		t.Errorf("Hook callback is not called")
	}

	var calledCallback2 bool = false
	// Add 2nd hook item
	expectedEventValueInCallback1 = 22
	var expectedEventValueInCallback2 = eventValue
	calledCallback1 = false
	item2 := hook.AddHook(func(event event.Event, nextHook event.Handler) bool {
			if calledCallback1 {
				t.Errorf("Wrong order of hook callback.")
			}
			calledCallback2 = true
			if event != eventSent {
				t.Errorf("Wrong event in hook callback. Want %v, got %v", eventSent, event)
			}
			if s := event.Sender(); s != sender {
				t.Errorf("Wrong sender of Event in hook callback. Want %v, got %v", sender, s)
			}
			if v := event.(*simpleEvent).Value(); v != expectedEventValueInCallback2 {
				t.Errorf("Wrong value of simpleEvent in hook callback. Want %v, got %v", expectedEventValueInCallback2, v)
			}
			event.(*simpleEvent).value = expectedEventValueInCallback1
			if ret := nextHook(event); ret != expectedNextHookRet {
				t.Errorf("Wrong return value of nextHook. Want %v, got %v", expectedNextHookRet, ret)
			}
			return false
		})
	if !hook.HasCallback() {
		t.Errorf("Non-empty Hook does not have callbacks")
	}
	if ret := hook.Call(eventSent); ret != false {
		t.Errorf("Wrong return value of hook.Call. Want %v, got %v", false, ret)
	}

	if !calledCallback1 {
		t.Errorf("Hook callback1 is not called")
	}

	if !calledCallback2 {
		t.Errorf("Hook callback2 is not called")
	}


	hook.RemoveHook(item2)
	if !hook.HasCallback() {
		t.Errorf("Non-empty Hook does not have callbacks")
	}
	calledCallback1, calledCallback2 = false, false
	expectedNextHookRet = false

	if ret := hook.Call(eventSent); ret != expectedNextHookRet {
		t.Errorf("Wrong return value of hook.Call after item2 unhooked. Want %v, got %v", expectedNextHookRet, ret)
	}

	if !calledCallback1 {
		t.Errorf("Callback1 not called after unhooking callback2")
	}

	if calledCallback2 {
		t.Errorf("callback2 is called after unhooking callback2")
	}

	hook.RemoveHook(item1)
	if hook.HasCallback() {
		t.Errorf("Empty Hook has callbacks")
	}
	calledCallback1 = false
	calledCallback2 = false

	if ret := hook.Call(eventSent); ret != false {
		t.Errorf("Wrong return value of hook.Call after unhooking. Want %v, got %v", false, ret)
	}

	if calledCallback1 {
		t.Errorf("Hook callback is still called after unhooking")
	}
}

func TestDefaultReturnValue(t *testing.T) {
	var sender interface{} = 1

	var hub = &event.Hub{DefaultReturnValue: true}
	if r := hub.Send(newSimpleEvent(sender, 100)); r != true {
		t.Errorf("Wrong return value. Want %v, got %v", true, r)
	}

	hub = &event.Hub{}
	if r := hub.Send(newSimpleEvent(sender, 100)); r != false {
		t.Errorf("Wrong return value. Want %v, got %v", false, r)
	}
}

func TestSimpleEventHandling(t *testing.T) {
	var sender interface{} = 1

	var hub = &event.Hub{}
	if hub.HasHandler() {
		t.Error("Empty hub has handler")
	}
	hub.SetHandler(func(event event.Event) bool {
		evt := event.(*simpleEvent)
		v := evt.Value()
		if v != 3 {
			t.Errorf("Wrong value in listener: %v. Want %v.\n", v, 3)
		}
		s := evt.Sender()
		if s != sender {
			t.Errorf("Wrong sender in listener: %v. Want %v.\n", s, sender)
		}
		return false
	})

	if !hub.HasHandler() {
		t.Error("HasHandler() on a non-empty hub returns false")
	}

	hook1 := hub.AddHook(func(event event.Event, next event.Handler) bool{
			v := (event.(*simpleEvent)).Value()
			if v != 2 {
				t.Errorf("Wrong value in hook: %v. Want %v.\n", v, 2)
			}
			(event.(*simpleEvent)).value++
			return next(event)
		})

	hook2 := hub.AddHook(func(event event.Event, next event.Handler) bool {
			v := (event.(*simpleEvent)).Value()
			if v != 1 {
				t.Errorf("Wrong value in hook: %v. Want %v.\n", v, 1)
			}
			(event.(*simpleEvent)).value = 2
			return next(event)
		})

	ret := hub.Send(newSimpleEvent(sender, 1))
	if ret != false {
		t.Errorf("Wrong return value of Send(): %v. Want %v\n", ret, false)
	}
	hub.RemoveHook(hook2)
	hub.RemoveHook(hook1)
	hub.SetHandler(nil)
	
	if hub.HasHandler() {
		t.Error("Empty hub has handler")
	}
}

func TestEventHandlingNotCallingNext(t *testing.T) {
	var hub = &event.Hub{}
	hub.SetHandler(func(event event.Event) bool {
		t.Errorf("Should not execute this listener.")
		return false
	})

	hub.AddHook(func(event event.Event, next event.Handler) bool{
			t.Errorf("Should not execute this hook.")
			return next(event)
		})

	hub.AddHook(func(event event.Event, next event.Handler) bool {
			v := (event.(*simpleEvent)).Value()
			if v != 1 {
				t.Errorf("Wrong value in hook: %v. Want %v.\n", v, 1)
			}
			(event.(*simpleEvent)).value = 2
			return true // Not calling next.
		})

	ret := hub.Send(newSimpleEvent(nil, 1))
	if ret != true {
		t.Errorf("Wrong return value of Send(): %v. Want %v\n", ret, false)
	}
}

func TestEventHandlingUnhook(t *testing.T) {
	var sender interface{} = 1

	var hub = &event.Hub{}
	hub.SetHandler(func(event event.Event) bool {
		evt := event.(*simpleEvent)
		v := evt.Value()
		if v != 2 {
			t.Errorf("Wrong value in listener: %v. Want %v.\n", v, 2)
		}
		s := evt.Sender()
		if s != sender {
			t.Errorf("Wrong sender in listener: %v. Want %v.\n", s, sender)
		}
		evt.SetValue(-1)
		return false
	})

	hook1 := hub.AddHook(func(event event.Event, next event.Handler) bool{
			t.Errorf("Should not execute this hook.")
			(event.(*simpleEvent)).SetValue(100)
			return next(event)
		})

	hub.AddHook(func(event event.Event, next event.Handler) bool {
			v := (event.(*simpleEvent)).Value()
			if v != 1 {
				t.Errorf("Wrong value in hook: %v. Want %v.\n", v, 1)
			}
			(event.(*simpleEvent)).value = 2
			hub.RemoveHook(hook1) // Test unhook in hook callback.
			return next(event)
		})

	event := newSimpleEvent(sender, 1)
	ret := hub.Send(event)
	if ret != false {
		t.Errorf("Wrong return value of Send(): %v. Want %v\n", ret, false)
	}
	value := event.Value()
	if value != -1 {
		t.Errorf("Wrong value after send: %v Wnat %v.\n",  value, -1)
	}
}

type simpleEvent struct {
	sender interface{}
	value int
}

func newSimpleEvent(sender interface{}, value int) *simpleEvent {
	return &simpleEvent{sender: sender, value: value}
}

func (event *simpleEvent) Sender() interface{} {
	return event.sender
}

func (event *simpleEvent) Value() int {
	return event.value
}

func (event *simpleEvent) SetValue(value int) {
	event.value = value
}

