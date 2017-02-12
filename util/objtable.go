package util

import (
	"fmt"
	"github.com/mkch/rw/native"
	"io"
)

// ObjectTable is a map of native.Handle to WrapperHolder.
type ObjectTable interface {
	// Register adds an WrapperHolder to this table, using it's handle as the key.
	Register(obj WrapperHolder)
	// Query returns the WrapperHolder with the handle in the table.
	Query(handle native.Handle) WrapperHolder
	// Remove removes the WrapperHolder with the handle form the table, and send AfterDestroyed event to the wrapper.
	Remove(handle native.Handle)
	Empty() bool
	Print(name string, w io.Writer)
}

type objectTable struct {
	m map[native.Handle]WrapperHolder
}

// NewObjectTable creates an ObjectTable.
func NewObjectTable() ObjectTable {
	return &objectTable{make(map[native.Handle]WrapperHolder)}
}

func (table *objectTable) Register(obj WrapperHolder) {
	handle := obj.Wrapper().Handle()
	if _, exists := table.m[handle]; exists {
		panic("Duplicated handle")
	}
	table.m[handle] = obj
}

func (table *objectTable) Query(handle native.Handle) WrapperHolder {
	return table.m[handle]
}

// Remove removes the WrapperHolder with the handle form the table, and send AfterDestroyed event to the wrapper.
func (table *objectTable) Remove(handle native.Handle) {
	if obj, ok := table.m[handle]; !ok {
		return
	} else {
		delete(table.m, handle)
		obj.Wrapper().setHandle(obj.Wrapper().HandleManager().Invalid())
		afterDestroyed := obj.Wrapper().AfterDestroyed()
		if afterDestroyed.HasCallback() {
			afterDestroyed.Call(&WrapperEvent{sender: obj, recreating: obj.Wrapper().Recreating(), bundle: obj.Wrapper().recreateBundle()})
		}
		obj.Wrapper().setHandle(0)
	}
}

func (table *objectTable) Empty() bool {
	return len(table.m) == 0
}

func (table *objectTable) Print(name string, w io.Writer) {
	fmt.Fprintf(w, "----- %v -----\n", name)
	var i int
	for handle, obj := range table.m {
		fmt.Fprintf(w, "%v)\t", i)
		printHandle(w, handle)
		fmt.Fprintf(w, "\t%T\t%[1]v\n", obj)
		i++
	}
}
