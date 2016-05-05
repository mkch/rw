package util

import (
	"io"
	"fmt"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/internal/native/darwin/object"
)

func printHandle(w io.Writer, handle native.Handle) {
	fmt.Fprintf(w, "%#X(retainCount: %v)", handle, object.NSObject_retainCount(handle))
}