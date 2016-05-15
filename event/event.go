package event

// Handler is the functio to handle event.
// It should return true if the event is processed, false otherwise.
type Handler func(event Event) bool

// Event represents a arbitrary event.
type Event interface {
	Sender() interface{}
}

type SimpleEvent struct {
	TheSender interface{}
}

func (evt *SimpleEvent) Sender() interface{} {
	return evt.TheSender
}

// Hub is a event hub to send and receive events.
// Hub can be created as part of other structures.
type Hub struct {
	// DefaultReturnValue is the listener return value if no listener set.
	DefaultReturnValue bool
	hookChain          HookChain
	listener           Handler
}

func (hub *Hub) defaultHookHandler(evt Event) bool {
	if hub.listener != nil {
		return hub.listener(evt)
	}
	return hub.DefaultReturnValue
}

// HasHandler returns true if an event sent to this hub will be processed by at least one Handler or Callback.
func (hub *Hub) HasHandler() bool {
	// Do not use hub.hookChain.HasCallback() here. hub.HookChian may contain the default handler.
	var hasHook = hub.hookChain.item != nil && hub.hookChain.item.next != nil
	return hasHook || hub.listener != nil
}

// Send sends an event to this hub. This event will be processed by any Handler and/or Callbacks.
// When Hub.Send is called: 1. The same behavior as HookChain.Send is performed.
// 2. Listener, if any, is called if the last called hook callback in step 1 calls nextHook or if no hook callbacks added.
func (hub *Hub) Send(event Event) bool {
	if hub.hookChain.item == nil {
		hub.hookChain.item = &hookItem{callback: hub.defaultHookHandler}
	}
	return hub.hookChain.Call(event)
}

// SetHandler sets a Handler to handle any events sent to this hub.
func (hub *Hub) SetHandler(handler Handler) {
	hub.listener = handler
}

// SetListener is a helper function of SetHandler.
// listener can be one of the 4 types:
// 1. Handler, equivalent to hub.SetHandler(listener);
// 2. func(), equivalent to hub.SetHandler(func(Event)bool{listener();return true});
// 3. func(Event), equivalent to hub.SetHandler(func(evt Event)bool{listener(evt);return true});
// 4. func() bool, equivalent to hub.SetHandler(func(Event)bool{return listener()}).
func (hub *Hub) SetListener(listener interface{}) {
	if handler, ok := listener.(Handler); ok {
		hub.SetHandler(handler)
	} else if fl, ok := listener.(func(event Event) bool); ok {
		hub.SetHandler(fl)
	} else if f, ok := listener.(func()); ok {
		hub.SetHandler(func(Event) bool { f(); return true })
	} else if fe, ok := listener.(func(Event)); ok {
		hub.SetHandler(func(evt Event) bool { fe(evt); return true })
	} else if fb, ok := listener.(func() bool); ok {
		hub.SetHandler(func(Event) bool { return fb() })
	} else {
		panic("Invalid listener. Must be one of: event.Handler, func(), func(event.Event), func()bool")
	}
}

// AddHook adds a hook entry to the hook chain. Hook callbacks has higher priority than Handlers.
// The returned HookItem value can be used to Unhook(). See HookChain.AddHook.
func (hub *Hub) AddHook(callback Callback) HookItem {
	if hub.hookChain.item == nil {
		hub.hookChain.item = &hookItem{callback: hub.defaultHookHandler}
	}
	return hub.hookChain.AddHook(callback)
}

// RemoveHook removes a hook entry from the hook chain. See HookChain.RemoveHook.
func (hub *Hub) RemoveHook(item HookItem) {
	hub.hookChain.RemoveHook(item)
}
