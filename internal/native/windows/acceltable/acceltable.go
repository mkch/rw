package acceltable

//#include <windows.h>
import "C"

import (
	"github.com/kevin-yuan/rw/internal/native/windows/post"
	"github.com/kevin-yuan/rw/internal/native/windows/nativeutil"
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/internal/mem"
	"unsafe"
	"fmt"
	"bytes"
)

type ACCEL C.ACCEL

func(accel ACCEL) String() string {
	return fmt.Sprintf("{fVirt=0x%x key=0x%x(%v) cmd=0x%x}", accel.fVirt, accel.key, string(accel.key), accel.cmd)
}

const C_ACCEL_SIZE = int(unsafe.Sizeof(C.ACCEL{}))

type AccelTable struct {
	// The accelerator table handle.
	accelTableHandle C.HACCEL
	// The array used to create the accelerator table.
	accelTableEntries []C.ACCEL
	// nil value for deletion, other values for updating and adding.
	pendingChanges map[C.WORD]*C.ACCEL
	onChanged func(native.Handle)
}

func (t *AccelTable) String() string {
	buf := bytes.NewBufferString("----- AccelTable -----\n")
	buf.WriteString(fmt.Sprintf("Handle=%p\n", unsafe.Pointer(t.accelTableHandle)))
	for i, accel := range t.accelTableEntries {
		buf.WriteString(fmt.Sprintf("%v) %v\n", i, ACCEL(accel)))
	}
	buf.WriteString("----------------------\n")
	return buf.String()
}

func (t *AccelTable) Handle() native.Handle {
	return native.Handle(unsafe.Pointer(t.accelTableHandle))
}

func (t *AccelTable) Destroy() {
	if t.accelTableHandle != nil {
		if C.DestroyAcceleratorTable(t.accelTableHandle) == C.FALSE {
			nativeutil.PanicWithLastError()
		}
		t.accelTableHandle = nil
	}
	if cap(t.accelTableEntries) > 0 {
		t.accelTableEntries = t.accelTableEntries[:cap(t.accelTableEntries)] // In case `&t.accelTableEntries[0]` out of index.
		mem.Free(unsafe.Pointer(&t.accelTableEntries[0]))
	}
	t.accelTableEntries = nil
}

const (
    FALT = byte(C.FALT)
    FCONTROL = byte(C.FCONTROL)
    FNOINVERT = byte(C.FNOINVERT)
    FSHIFT = byte(C.FSHIFT)
    FVIRTKEY = byte(C.FVIRTKEY)
)

func (t *AccelTable) Add(fVirt byte, key, cmd uint16) {
	if t.pendingChanges == nil {
		// Add.
		post.Post(t.rebuildTable)
		t.pendingChanges = make(map[C.WORD]*C.ACCEL)
		t.pendingChanges[C.WORD(cmd)] = &C.ACCEL{fVirt:C.BYTE(fVirt), key:C.WORD(key), cmd:C.WORD(cmd)}
	} else if accel := t.pendingChanges[C.WORD(cmd)]; accel != nil {
		// Update existing or readd the deleted entry.
		accel.fVirt = C.BYTE(fVirt)
		accel.key = C.WORD(key)
	} else {
		// Add.
		t.pendingChanges[C.WORD(cmd)] = &C.ACCEL{fVirt:C.BYTE(fVirt), key:C.WORD(key), cmd:C.WORD(cmd)}
	}
}

func (t *AccelTable) Remove(cmd uint16) {
	if t.pendingChanges == nil {
		post.Post(t.rebuildTable)
		t.pendingChanges = make(map[C.WORD]*C.ACCEL)
	} 
	// Deletion.
	t.pendingChanges[C.WORD(cmd)] = nil
}

func (t *AccelTable) SetOnChangedListener(l func(native.Handle)) {
	t.onChanged = l
}

func (t *AccelTable) rebuildTable() {
	if len(t.pendingChanges) == 0 {
		return // Nothing changed.
	}
	changed := false
	// Process deletion.
	for cmd, accel := range t.pendingChanges {
		if accel != nil {
			continue
		}
		for i, _ := range t.accelTableEntries {
			if t.accelTableEntries[i].cmd == cmd {
				// Delete this entry.
				copy(t.accelTableEntries[i:], t.accelTableEntries[i+1:])
				t.accelTableEntries = t.accelTableEntries[:len(t.accelTableEntries)-1]
				changed = true
				break
			}
		}
	}
	//Process updating/adding
	PendingChangesLoop:
	for cmd, accel := range t.pendingChanges {
		if accel == nil {
			continue
		}
		for i, _ := range t.accelTableEntries {
			if t.accelTableEntries[i].cmd == cmd {
				// Update this entry if necessary.
				if t.accelTableEntries[i].fVirt != accel.fVirt {
					t.accelTableEntries[i].fVirt = accel.fVirt
					changed = true
				}
				if t.accelTableEntries[i].key != accel.key {
					t.accelTableEntries[i].key = accel.key
					changed = true
				}
				continue PendingChangesLoop
			}
		}
		// Not found, add new entry.
		c := cap(t.accelTableEntries)
		l := len(t.accelTableEntries)
		if t.accelTableEntries == nil { // Not allocated
			c = 16
			t.accelTableEntries = (*[(1<<30)/C_ACCEL_SIZE]C.ACCEL)(mem.Alloc(uintptr(C_ACCEL_SIZE*c)))[:0:c]
		} else if c == l { // No more capacity
			c *= 2
			t.accelTableEntries = (*[(1<<30)/C_ACCEL_SIZE]C.ACCEL)(mem.Realloc(unsafe.Pointer(&t.accelTableEntries[0]), uintptr(C_ACCEL_SIZE*c)))[:l:c]
		}
		t.accelTableEntries = append(t.accelTableEntries, *accel)
		changed = true
	}

	if !changed {
		return // Nothing changed.
	}

	if t.accelTableHandle != nil {
		if C.DestroyAcceleratorTable(t.accelTableHandle) == C.FALSE {
			nativeutil.PanicWithLastError()
		}
	}
	if entryCount := len(t.accelTableEntries); entryCount > 0 {
		t.accelTableHandle = C.CreateAcceleratorTable(&t.accelTableEntries[0], C.int(entryCount))
		if t.accelTableHandle == nil {
			nativeutil.PanicWithLastError()
		}
	} else {
		t.accelTableHandle = nil
	}

	if t.onChanged != nil {
		t.onChanged(t.Handle())
	}
	t.pendingChanges = nil
}

