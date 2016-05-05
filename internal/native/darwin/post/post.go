package post

//#include "post.h"
//extern void safePostCallback(UINTPTR);
//extern void unsafePostCallback(UINTPTR);
import "C"

import (
	"github.com/mkch/rw/internal/stackescape"
)

func Init() {
	C.initPostOnMainThread(C.FN_POST_CALLBACK(C.safePostCallback), C.FN_POST_CALLBACK(C.unsafePostCallback))
}

var safePostTable = stackescape.NewSafeTable();

//export safePostCallback
func safePostCallback(userData uintptr) {
	id := stackescape.Id(userData)
	f := safePostTable.Get(id).(func())
	f()
	safePostTable.Remove(id)
}

//export unsafePostCallback
func unsafePostCallback(userData uintptr) {
	id := stackescape.Id(userData)
	f := stackescape.Get(id).(func())
	f()
	stackescape.Remove(id)
}

func Post(f func()) {
	C.postOnMainThread(C.UINTPTR(safePostTable.Add(f)), true)
}

func UnsafePost(f func()) {
	C.postOnMainThread(C.UINTPTR(stackescape.Add(f)), false)
}