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
	// This method can do more than `return  handle != Invalid()`, e.g. rang checking.
	Valid(handle native.Handle) bool
	// Invalid returns a invalid object.
	Invalid() native.Handle
	// Table returns the ObjectTable in which this object is.
	Table() *ObjectTable
}

type WrapperEvent struct {
	sender     interface{}
	recreating bool
	bundle     Bundle
}

func (evt *WrapperEvent) Sender() interface{} {
	return evt.sender
}

// Recreating returns true if this object is being recreated.
func (evt *WrapperEvent) Recreating() bool {
	return evt.recreating
}

func (evt *WrapperEvent) Bundle() Bundle {
	return evt.bundle
}

// Wrapper wraps a native object.
type Wrapper interface {
	// Handle returns the wrapped native object. Panics if Valid() == false.
	Handle() native.Handle
	// Valid returns whether a valid native object is wrapped.
	Valid() bool
	HandleManager() HandleManager
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
	// recreateBundle returns the Bundle passed in Recreate.
	recreateBundle() Bundle
	setRecreateBundle(Bundle)
}

// WrapperImpl implements Wrapper interface.
type WrapperImpl struct {
	handle          native.Handle
	isRecreating    bool
	recreatArg      Bundle
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

func (w *WrapperImpl) recreateBundle() Bundle {
	return w.recreatArg
}
func (w *WrapperImpl) setRecreateBundle(b Bundle) {
	w.recreatArg = b
}

type WrapperHolder interface {
	Wrapper() Wrapper
}

// Init initializes a Wrapper.
func Init(w WrapperHolder) {
	w.Wrapper().setHandle(w.Wrapper().HandleManager().Create(nil))
	register(w, nil)
}

// InitWithHandle initializes a Wrapper with an existing handle.
func InitWithHandle(w WrapperHolder, handle native.Handle) {
	w.Wrapper().setHandle(handle)
	register(w, nil)
}

// register register Wrapper to it's WrapperTable
func register(w WrapperHolder, b Bundle) {
	wrapper := w.Wrapper()
	wrapper.HandleManager().Table().Register(w)
	afterRegistered := wrapper.AfterRegistered()
	if afterRegistered.HasCallback() {
		afterRegistered.Call(&WrapperEvent{sender: w, recreating: wrapper.Recreating(), bundle: b})
	}
}

func destroy(w WrapperHolder) {
	wrapper := w.Wrapper()
	wrapper.HandleManager().Destroy(wrapper.Handle())
	// Do not call wrapper.AfterDestroyed().Call() here.
	// HandleManager().Destroy() should call ObjectTable.Remove(),
	// where it is called.
}

// Recreate destroys the existing native object and create a new one.
// The bundle is passed to HandleManager.Create.
func Recreate(w WrapperHolder, b Bundle) {
	wrapper := w.Wrapper()
	handleManager := wrapper.HandleManager()

	wrapper.setRecreating(true)
	wrapper.setRecreateBundle(b)
	defer func() {
		wrapper.setRecreating(false)
		wrapper.setRecreateBundle(nil)
	}()
	destroy(w)
	wrapper.setHandle(handleManager.Create(b))
	register(w, b)
}

func Release(w WrapperHolder) {
	if w.Wrapper().Valid() {
		destroy(w)
	}
}
