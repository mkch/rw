package util

import (
	"fmt"
	"github.com/mkch/rw/native"
	"io"
)

// ObjectTableStorage is the abstraction of key-value(native.Handle to WrapperHolder) map implementations.
type ObjectTableStorage interface {
	// Get returns the associated value with the key.
	Get(key native.Handle) (value WrapperHolder, exists bool)
	// Set associates a value with the key. Overwrites the existing value associated with the same key.
	Set(key native.Handle, value WrapperHolder)
	// Del remove the associated value of key. Do nothing if there is no value associated with the key.
	Del(key native.Handle)
	// Len returns the size(number of values) of the table.
	Len() int
	// ForEach calls f with every key-value pair in the table.
	ForEach(f func(native.Handle, WrapperHolder))
}

// mapObjectTableStorage is a ObjectTableStorage implementation using a map.
type mapObjectTableStorage map[native.Handle]WrapperHolder

func (m mapObjectTableStorage) Get(key native.Handle) (value WrapperHolder, exists bool) {
	value, exists = m[key]
	return
}

func (m mapObjectTableStorage) Set(key native.Handle, value WrapperHolder) {
	m[key] = value
}

func (m mapObjectTableStorage) Del(key native.Handle) {
	delete(m, key)
}

func (m mapObjectTableStorage) Len() int {
	return len(m)
}

func (m mapObjectTableStorage) ForEach(f func(native.Handle, WrapperHolder)) {
	for k, v := range m {
		f(k, v)
	}
}

// ObjectTable is a map of native.Handle to WrapperHolder.
type ObjectTable struct {
	m ObjectTableStorage
}

// NewObjectTableWithStorage creates an ObjectTable using storage
// as it's storage implementation.
func NewObjectTableWithStorage(storage ObjectTableStorage) *ObjectTable {
	return &ObjectTable{storage}
}

// NewObjectTable creates an ObjectTable using default storage implementation.
func NewObjectTable() *ObjectTable {
	return &ObjectTable{make(mapObjectTableStorage)}
}

// Register adds an WrapperHolder to this table, using it's handle as the key.
func (table *ObjectTable) Register(obj WrapperHolder) {
	handle := obj.Wrapper().Handle()
	if _, exists := table.m.Get(handle); exists {
		panic("Duplicated handle")
	}
	table.m.Set(handle, obj)
}

// Query returns the WrapperHolder with the handle in the table.
func (table *ObjectTable) Query(handle native.Handle) (obj WrapperHolder) {
	obj, _ = table.m.Get(handle)
	return
}

// Remove removes the WrapperHolder with the handle form the table, and send AfterDestroyed event to the wrapper.
func (table *ObjectTable) Remove(handle native.Handle) {
	if obj, exists := table.m.Get(handle); !exists {
		return
	} else {
		table.m.Del(handle)
		obj.Wrapper().setHandle(obj.Wrapper().HandleManager().Invalid())
		afterDestroyed := obj.Wrapper().AfterDestroyed()
		if afterDestroyed.HasCallback() {
			afterDestroyed.Call(&WrapperEvent{sender: obj, recreating: obj.Wrapper().Recreating(), bundle: obj.Wrapper().recreateBundle()})
		}
		obj.Wrapper().setHandle(0)
	}
}

func (table *ObjectTable) Empty() bool {
	return table.m.Len() == 0
}

func (table *ObjectTable) Print(name string, w io.Writer) {
	fmt.Fprintf(w, "----- %v -----\n", name)
	var i int
	table.m.ForEach(func(handle native.Handle, obj WrapperHolder) {
		fmt.Fprintf(w, "%v)\t", i)
		printHandle(w, handle)
		fmt.Fprintf(w, "\t%T\t%[1]v\n", obj)
		i++
	})
}
