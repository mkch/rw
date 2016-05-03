package event

// hookItem represents an entry in the hook chain.
// Next and Call of zero value of hookItem return false.
type hookItem struct {
	callback Handler // The callback function of this item.
	next *hookItem	// Next hook item in the chain.
}

// Next calls the next hook item and returns its return value.
// False is returned if no next item evailable.
func (i *hookItem) Next(event Event) bool {
	if i.next != nil && i.next.callback != nil{
		return i.next.callback(event)
	} else {
		return false
	}
}

// Call calls the callback of this item.
// False is returned if no callbak is available.
func (i *hookItem) Call(event Event) bool {
	if i.callback == nil  {
		return false
	}
	return i.callback(event)
}

// Callback is the hook callback function.
// nextHook is the next handler in hook chain. Call nextHook(event) to pass
// the control downwards the hook chain.
// It should return true if the event is processed, false otherwise.
type Callback func(event Event, nextHook Handler) bool

// HookItem is the identifier of a hook entry in the hook chain.
// It can be passed to HookChain.Unhook() to remove a hook entry form the hook chain.
type HookItem *hookItem

// HookChain is a chain of callbacks.
// HookChain can be created as part of other structures.
type HookChain struct {
	//DefaultReturnValue is used as the return value of Call if no hook is added.
	DefaultReturnValue bool
	item *hookItem
}

// HasCallback returns whether the hook has any callbasks.
func (h *HookChain) HasCallback() bool {
	return h.item != nil
}

// Call calls the last added callback and returns it's return value if any, or DefaultReturnValue is returned.
// The callback can process the event, and optionally modifies the event, passes event to next callback(the one
// added just before this callback), and so forth.
// If a callback does not call nextHook to pass event to next callback, no more callbacks will be called.
func (h *HookChain) Call(event Event) bool {
	if h.item == nil {
		return h.DefaultReturnValue
	}
	return h.item.Call(event)
}

// addFront adds a handler in the front, and returns the newly added hook item.
// The returned hookItem can be used to call the next callback or passed to Unhook().
func (h *HookChain) addFront(handler Handler) *hookItem {
	if handler == nil {
		panic("nil handler")
	}
	item := &hookItem{callback: handler, next: h.item}
	h.item = item
	return item
}

// AddHook adds a callback to the hook chain.
// The returned HookItem value can be used to call RemoveHook.
func (h *HookChain) AddHook(callback Callback) HookItem {
	var item *hookItem
	item = h.addFront(func (event Event) bool {
		return callback(event, func(event Event) bool{ return item.Next(event) })
	})
	return item
}

func panicWhenCalled(event Event) bool {
	panic("Handler invalid")
}

// Contains returns whether item is this hook chain.
func (h *HookChain) Contains(item HookItem) bool {
	if item == nil {
		panic ("nil hook item")
	}
	
	for p := h.item; p != nil; p = p.next {
		if p == item {
			return true
		}
	}
	return false
}

// RemoveHook removes a callback from the chain.
func (h *HookChain)RemoveHook(item HookItem) {
	if h.item == item {
		h.item = item.next
		// panic when operating item after unhook.
		item.callback = panicWhenCalled
		item.next = nil
		return
	}

	p := h.item
	for p != nil && p.next != item {
		p = p.next
	}
	if p == nil {
		panic("Invalid item to unhook")
	}

	p.next = item.next
	// panic when operating item after unhook.
	item.callback = panicWhenCalled
	item.next = nil
}


