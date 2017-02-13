package rw

import (
	"fmt"
	"github.com/mkch/rw/internal/native/darwin/app"
	"github.com/mkch/rw/internal/native/darwin/autoreleasepool"
	"github.com/mkch/rw/internal/native/darwin/deallochook"
	"github.com/mkch/rw/internal/native/darwin/post"
	"github.com/mkch/rw/native"
	"io"
)

func nativeInit() {
	app.RWApp_sharedApplication()
	app.NSApplication_setActivationPolicy(app.NSApp(), app.NSApplicationActivationPolicyRegular)
	post.Init()
	deallochook.Init(func(handle native.Handle) {
		defaultObjectTable.Remove(handle)
		// Remove the window(if is a window) from disabled window list.
		app.EnableWindow(handle, true)
	})
}

func nativeRun(initializeCallback, terminateCallback func()) {
	autoreleasepool.Run(func() {
		if initializeCallback != nil {
			initializeCallback()
		}

		app.NSApplication_run(app.NSApp())

		if terminateCallback != nil {
			terminateCallback()
		}
	})
}

func nativeExit() {
	Post(func() {
		//If you call this method from an event handler running in your main run loop, the app object exits out of the run method, thereby returning control to the main() function. If you call this method from within a modal event loop, it will exit the modal loop instead of the main event loop.
		app.NSApplication_stop(app.NSApp(), app.NSApp())
	})
}

func objectsLeaked() bool {
	return !defaultObjectTable.Empty()
}

func printLeakedObjects(w io.Writer) {
	fmt.Fprintln(w, "// BUG(?) The main menu of application is always reported leaked. Never released by Cocoa framework?")
	if !defaultObjectTable.Empty() {
		defaultObjectTable.Print("default object table", w)
	}
}
