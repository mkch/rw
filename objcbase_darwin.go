package rw

import (
	"github.com/mkch/rw/util"

	"github.com/mkch/rw/internal/native/darwin/object"
	"github.com/mkch/rw/native"
)

// objcBase implements the common operations of NSObject
type objcBase struct {
	objectBase
	wrapper util.WrapperImpl
}

func (w *objcBase) Wrapper() util.Wrapper {
	return &w.wrapper
}

func (w *objcBase) Release() {
	util.Release(w)
}

// objcHandleManager is the common building block of all concrete HandleManagers.
// Create method calls objcHandleManager itself.
// Converting a `func(util.Bundle) native.Handle` to objcHandleManager makes a useable objcHandleManager.
type objcHandleManager func(util.Bundle) native.Handle

func (m objcHandleManager) Create(b util.Bundle) native.Handle {
	return m(b)
}

func (m objcHandleManager) Destroy(handle native.Handle) {
	object.NSObject_release(handle)
}

func (m objcHandleManager) Valid(handle native.Handle) bool {
	return handle != 0
}

func (m objcHandleManager) Table() util.ObjectTable {
	return defaultObjectTable
}
