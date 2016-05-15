package post

//#include "post.h"
import "C"

import (
	"github.com/mkch/rw/internal/native/windows/nativeutil"
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/internal/stackescape"
	"github.com/mkch/rw/native"
)

//export runUnsafePostedFunc
func runUnsafePostedFunc(id uintptr) {
	stackescape.Get(stackescape.Id(id)).(func())()
}

var safePostTable = stackescape.NewSafeTable()

//export runSafePostedFunc
func runSafePostedFunc(id uintptr) {
	safePostTable.Get(stackescape.Id(id)).(func())()
}

func Post(f func()) {
	window.PostMessage(postMessageOnlyWindow, uint(C.WM_GO_SAFE_POST), uintptr(safePostTable.Add(f)), 0)
}

func UnsafePost(f func()) {
	window.PostMessage(postMessageOnlyWindow, uint(C.WM_GO_UNSAFE_POST), uintptr(stackescape.Add(f)), 0)
}

var postMessageOnlyWindow native.Handle

func Init() {
	postMessageOnlyWindow = native.Handle(C.Ptr(C.createPostMessageOnlyWindow()))
	if postMessageOnlyWindow == 0 {
		nativeutil.PanicWithLastError()
	}
}
