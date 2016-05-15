// Make it safe to pass a go pointer to C and pass back to go.
package stackescape

import (
	"fmt"
	"os"
	"sync"
)

type Id uintptr
type list map[Id]interface{}

const firstId = Id(1)
const maxId = ^Id(0)

// Table is a map of Id to interface{}.
// Table is not concurrency safe.
type Table struct {
	l          list
	nextId     Id
	idRewinded bool
}

func NewTable() *Table {
	return &Table{
		l:      make(list),
		nextId: 1,
	}
}

// Add a pointer to the escape list. An unique id of this pointer is returned.
func (tab *Table) Add(ptr interface{}) (id Id) {
	id = tab.nextId
	tab.l[id] = ptr
	if tab.idRewinded || tab.nextId == maxId {
		fmt.Fprintf(os.Stderr, "WARNNING: %#T Id rewinded!\n", tab)
		tab.idRewinded = true
		found := false
		for i := firstId; i <= maxId; i++ {
			if _, ok := tab.l[i]; !ok {
				tab.nextId = i
				found = true
				break
			}
		}
		if !found {
			panic("Out of id")
		}
	} else {
		tab.nextId++
	}
	return id
}

// Get the pointer associated to the id.
func (tab *Table) Get(id Id) interface{} {
	return tab.l[id]
}

// Remove escape entry.
func (tab *Table) Remove(id Id) {
	delete(tab.l, id)
}

// SafeTable is the concurrency safe version of Table.
// Multi goroutines can call the medhods concurrently.
type SafeTable struct {
	Table
	m sync.RWMutex
}

func (tab *SafeTable) Add(ptr interface{}) (id Id) {
	tab.m.Lock()
	defer tab.m.Unlock()
	return tab.Table.Add(ptr)
}

func (tab *SafeTable) Get(id Id) interface{} {
	tab.m.RLock()
	defer tab.m.RUnlock()
	return tab.Table.Get(id)
}

func (tab *SafeTable) Remove(id Id) {
	tab.m.Lock()
	defer tab.m.Unlock()
	tab.Table.Remove(id)
}

func NewSafeTable() *SafeTable {
	return &SafeTable{
		Table: *NewTable(),
	}
}

var defaultTable = NewTable()

// Add calls Add method of the global Table. Not concurrency safe.
func Add(ptr interface{}) (id Id) {
	return defaultTable.Add(ptr)
}

// Get calls Get method of the global Table. Not concurrency safe.
func Get(id Id) interface{} {
	return defaultTable.Get(id)
}

// Remove calls Remove method of the global Table. Not concurrency safe.
func Remove(id Id) {
	defaultTable.Remove(id)
}
