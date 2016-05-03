package rw

import (
	"github.com/kevin-yuan/rw/util"

	"github.com/kevin-yuan/rw/internal/native/darwin/object"
	"github.com/kevin-yuan/rw/native"
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

// objcHandleManagerBase is the common building block of all concrete HandleManagers.
// Create method should be added to make a HandleManager.
type objcHandleManagerBase struct{}

func (m objcHandleManagerBase) Destroy(handle native.Handle) {
	object.NSObject_release(handle)
}

func (m objcHandleManagerBase) Valid(handle native.Handle) bool {
	return handle != 0
}

func (m objcHandleManagerBase) Table() util.ObjectTable {
	return defaultObjectTable
}
