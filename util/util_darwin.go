package util

import (
	"io"
	"fmt"
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/internal/native/darwin/object"
)

func printHandle(w io.Writer, handle native.Handle) {
	fmt.Fprintf(w, "%#X(retainCount: %v)", handle, object.NSObject_retainCount(handle))
}