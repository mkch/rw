package util

import (
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/native"
)

type Bundle map[string]interface{}

type HandleManager interface {
	// Create creates a new native object. The Bundle is provided by Recreate, nil if isn't called by Recreate.
	Create(Bundle) native.Handle
	// Destroy destroies a native object created by Create.
	Destroy(handle native.Handle)
	// Valid returns whether a valid native object is wrapped.
	Valid(handle native.Handle) bool
	// Table returns the ObjectTable in which this object is.
	Table() ObjectTable
}

type WrapperEvent struct {
	sender     interface{}
	recreating bool
}

func (evt *WrapperEvent) Sender() interface{} {
	return evt.sender
}

// Recreating returns true if this object is being recreated.
func (evt *WrapperEvent) Recreating() bool {
	return evt.recreating
}

// Wrapper wraps a native object.
type Wrapper interface {
	// Handle returns the wrapped native object. Panics if Valid() == false.
	Handle() native.Handle
	// Valid returns whether a valid native object is wrapped.
	Valid() bool
	HandleManager() HandleManager
	SetHandleManager(m HandleManager)
	// The type of event is *WrapperEvent.
	AfterRegistered() *event.HookChain
	// The type of event is *WrapperEvent.
	AfterDestroyed() *event.HookChain
	// setHandle sets the wrapped object.
	setHandle(handle native.Handle)
	// Recreating returns true if this object is being recreated.
	Recreating() bool
	// setRecreating sets whether this object is being recreated.
	setRecreating(r bool)
}

// WrapperImpl implements Wrapper interface.
type WrapperImpl struct {
	handle          native.Handle
	isRecreating    bool
	hm              HandleManager
	afterRegistered event.HookChain
	afterDestroyed  event.HookChain
}

func (w *WrapperImpl) Handle() native.Handle {
	if !w.HandleManager().Valid(w.handle) {
		panic(&InvalidObjectError{})
	}
	return w.handle
}

func (w *WrapperImpl) Valid() bool {
	return w.HandleManager().Valid(w.handle)
}

func (w *WrapperImpl) setHandle(handle native.Handle) {
	w.handle = handle
}

func (w *WrapperImpl) Recreating() bool {
	return w.isRecreating
}

func (w *WrapperImpl) setRecreating(r bool) {
	w.isRecreating = r
}

func (w *WrapperImpl) HandleManager() HandleManager {
	return w.hm
}

func (w *WrapperImpl) SetHandleManager(hm HandleManager) {
	w.hm = hm
}

func (w *WrapperImpl) AfterRegistered() *event.HookChain {
	return &w.afterRegistered
}

func (w *WrapperImpl) AfterDestroyed() *event.HookChain {
	return &w.afterDestroyed
}

type WrapperHolder interface {
	Wrapper() Wrapper
}

// Init initializes a Wrapper.
func Init(w WrapperHolder) {
	w.Wrapper().setHandle(w.Wrapper().HandleManager().Create(nil))
	register(w)
}

// InitWithHandle initializes a Wrapper with an existing handle.
func InitWithHandle(w WrapperHolder, handle native.Handle) {
	w.Wrapper().setHandle(handle)
	register(w)
}

// register register Wrapper to it's WrapperTable
func register(w WrapperHolder) {
	w.Wrapper().HandleManager().Table().Register(w)
	afterRegistered := w.Wrapper().AfterRegistered()
	if afterRegistered.HasCallback() {
		afterRegistered.Call(&WrapperEvent{sender: w, recreating: w.Wrapper().Recreating()})
	}
}

// Recreate destroys the existing native object and create a new one.
// The bundle is passed to HandleManager.Create.
func Recreate(w WrapperHolder, b Bundle) {
	wrapper := w.Wrapper()
	handleManager := wrapper.HandleManager()

	wrapper.setRecreating(true)
	defer wrapper.setRecreating(false)
	handleManager.Destroy(wrapper.Handle())
	wrapper.setHandle(handleManager.Create(b))
	register(w)
}

func Release(w WrapperHolder) {
	if w.Wrapper().Valid() {
		w.Wrapper().HandleManager().Destroy(w.Wrapper().Handle())
	}
}
