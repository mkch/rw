package util

import (
	"github.com/mkch/rw/native"
	"io"
	"fmt"
)

// ObjectTable is a map of native.Handle to WrapperHolder.
type ObjectTable struct {
	m map[native.Handle] WrapperHolder
}

// NewObjectTable creates an ObjectTable.
func NewObjectTable() ObjectTable {
	return ObjectTable{ make(map[native.Handle]WrapperHolder) }
}

// Register adds an WrapperHolder to this table, using it's handle as the key.
func (table ObjectTable)Register(obj WrapperHolder) {
	handle := obj.Wrapper().Handle()
	if _, exists := table.m[handle]; exists {
		panic("Duplicated handle")
	}
	table.m[handle] = obj
}

// Query returns the WrapperHolder with the handle in the table.
func (table ObjectTable)Query(handle native.Handle) WrapperHolder {
	return table.m[handle]
}

// Remove removes the WrapperHolder with the handle form the table, and send AfterDestroyed event to the wrapper.
func  (table ObjectTable)Remove(handle native.Handle) {
	if obj, ok := table.m[handle]; !ok {
		return
	} else {
		delete(table.m, handle)
		obj.Wrapper().setHandle(0)
		afterDestroyed := obj.Wrapper().AfterDestroyed()
		if afterDestroyed.HasCallback() {
			afterDestroyed.Call(&WrapperEvent{sender: obj, recreating: obj.Wrapper().Recreating()})
		}
		obj.Wrapper().setHandle(0)
	}
}

func (table ObjectTable) Empty() bool {
	return len(table.m) == 0
}

func (table ObjectTable) Print(name string, w io.Writer) {
	fmt.Fprintf(w, "----- %v -----\n", name)
	var i int
	for handle, obj := range table.m {
		fmt.Fprintf(w, "%v)\t", i)
		printHandle(w, handle)
		fmt.Fprintf(w, "\t%T\t%[1]v\n", obj)
		i++
	}
}