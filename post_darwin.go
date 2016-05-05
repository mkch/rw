package rw

import (
	"github.com/mkch/rw/internal/native/darwin/post"
)

// Post posts a function to the GUI goroutine and runs it there.
func Post(f func()) {
	post.Post(f)
}

// Only safe when used in the UI goroutine.
func unsafePost(f func()) {
	post.UnsafePost(f)
}
