package rw

import (
	"fmt"
	"os"
	"runtime"
)

var initialized bool

// Start the GUI system. Do any initialization work in initializeCallback.
// This function never return until rw.Exit() is called.
func Run(initializeCallback func()) {
	// Make the main goroutine executs in this OS thread only.
	// OS GUI call must be single threaded.
	runtime.LockOSThread()
	if !initialized {
		nativeInit()
		initialized = true
	}
	nativeRun(initializeCallback, nil)
	if Debug {
		if objectsLeaked() {
			w := os.Stderr
			fmt.Fprintln(w, "** Leaked object(s) found!")
			printLeakedObjects(w)
		}
	}
}

// Exit schedules the application to exit.
func Exit() {
	nativeExit()
}

// Debug is a flag indicates whether debug mode is active.
var Debug bool
