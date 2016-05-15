package util

import (
	"fmt"
	"github.com/mkch/rw/internal/native/darwin/object"
	"github.com/mkch/rw/native"
	"io"
)

func printHandle(w io.Writer, handle native.Handle) {
	fmt.Fprintf(w, "%#X(retainCount: %v)", handle, object.NSObject_retainCount(handle))
}
