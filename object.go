package rw

import (
	"fmt"
	"github.com/kevin-yuan/rw/util"
)

// In Darwin: The defaultObjectTable which contains all NSObject* wrappers.
// In Windows: The defaultObjectTable which contains all HWND wrappers
var defaultObjectTable = util.NewObjectTable()

// Object is the base rw interface.
type Object interface {
	// Release destroies the object and free everything it uses.
	// Releasing an already released object does nothing.
	Release()
	// Self returns the value passed to function Init, or nil if Init is not called on this object.
	// We want something like "this" pointer in C++/Java.
	Self() Object
	// String returns the string representation of this object.
	String() string
	setSelf(self Object)
}

// objectBase is the common part of all Objects.
type objectBase struct {
	self Object
}

func (obj *objectBase) Self() Object {
	return obj.self
}

func (obj *objectBase) setSelf(self Object) {
	obj.self = self
}

func (obj *objectBase) String() string {
	return fmt.Sprintf("Object %p", obj)
}

// Init initializes an object.
// Objects must be initialized before using. Objects created by NewXxx() is already initialized and there is
// no need to call Init again on them.
func Init(obj Object) {
	obj.setSelf(obj)
	if wrapperHolder, ok := obj.(util.WrapperHolder); ok {
		util.Init(wrapperHolder)
	}
}
