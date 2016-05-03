package rw

import (
	"github.com/kevin-yuan/rw/event"
	"github.com/kevin-yuan/rw/internal/native/darwin/app"
	"github.com/kevin-yuan/rw/internal/native/darwin/date"
	"github.com/kevin-yuan/rw/internal/native/darwin/runloop"
	"github.com/kevin-yuan/rw/internal/native/darwin/dynamicinvocation"
	nativeEvent "github.com/kevin-yuan/rw/internal/native/darwin/event"
	"github.com/kevin-yuan/rw/internal/native/darwin/notification"
	"github.com/kevin-yuan/rw/internal/native/darwin/object"
	"github.com/kevin-yuan/rw/internal/native/darwin/screen"
	"github.com/kevin-yuan/rw/internal/native/darwin/view"
	"github.com/kevin-yuan/rw/internal/native/darwin/window"
	"github.com/kevin-yuan/rw/util"
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/internal/native/darwin/windowstyle"
	s "github.com/kevin-yuan/rw/internal/windowstyle"
	"github.com/kevin-yuan/rw/internal/native/darwin/deallochook"
	"unsafe"
)

type windowExtra interface {
	recalculateTabOrder()
}

type WindowPlatformSpecific interface{}

type windowBase struct {
	objcBase
	onClose                 event.Hub
	content                 Container
	modalResult             interface{}
	y                       int
	recalcTabOrderScheduled bool
	inModal bool // Whether this window is in modal(application modal or window modal).
}

func flipWindowY(y, height int, _screen native.Handle) int {
	_, screenVisibleY, _, screenVisibleHeight := screen.NSScreen_visibleFrame(_screen)
	return screenVisibleHeight - height - y + screenVisibleY
}

func (w *windowBase) recalculateTabOrder() {
	if w.recalcTabOrderScheduled {
		return
	}
	unsafePost(func() {
		head, tails := recalculateTabOrder(w.content, 0, nil)
		for _, tail := range tails {
			view.NSView_setNextKeyView(tail.Wrapper().Handle(), head)
		}
		w.recalcTabOrderScheduled = false
	})
	w.recalcTabOrderScheduled = true
}

func (w *windowBase) OnClose() *event.Hub {
	w.ensureDelegate()
	return &w.onClose
}

func (w *windowBase) Enabled() bool {
	return app.WindowEnabled(w.Wrapper().Handle())
}

func (w *windowBase) SetEnabled(enabled bool) {
	app.EnableWindow(w.Wrapper().Handle(), enabled)
}

func (w *windowBase) Visible() bool {
	return window.NSWindow_isVisible(w.Wrapper().Handle())
}

func (w *windowBase) SetVisible(v bool) {
	handle := w.Wrapper().Handle()
	if v {
		window.NSWindow_orderFront(handle)
		_, y, _, height := window.NSWindow_frame(handle)
		if screen := window.NSWindow_screen(handle); screen != 0 {
			w.y = flipWindowY(y, height, screen)
		}
	} else {
		_, y, _, height := window.NSWindow_frame(handle)
		if screen := window.NSWindow_screen(handle); screen != 0 {
			w.y = flipWindowY(y, height, screen)
		}
		window.NSWindow_orderOut(handle)
	}
}

func (w *windowBase) CenterToScreen() {
	window.NSWindow_center(w.Wrapper().Handle())
}

func (w *windowBase) Title() string {
	return window.NSWindow_title(w.Wrapper().Handle())
}

func (w *windowBase) SetTitle(title string) {
	window.NSWindow_setTitle(w.Wrapper().Handle(), title)
}

func (w *windowBase) Menu() Menu {
	return nil
}

func (w *windowBase) SetMenu(menu Menu) {
	// Nothing to do
}

func (w *windowBase) ShowActive() {
	handle := w.Wrapper().Handle()
	window.NSWindow_makeKeyAndOrderFront(handle)
	_, y, _, height := window.NSWindow_frame(handle)
	if screen := window.NSWindow_screen(handle); screen != 0 {
		w.y = flipWindowY(y, height, screen)
	}
}

func (w *windowBase) Close() {
	if w.inModal {
		w.CloseModal(nil)
	} else {
		window.NSWindow_close(w.Wrapper().Handle())
	}
}


func (w *windowBase) Frame() Rect {
	handle := w.Wrapper().Handle()
	x, y, width, height := window.NSWindow_frame(handle)
	if screen := window.NSWindow_screen(handle); screen != 0 {
		y = flipWindowY(y, height, screen)
		w.y = y
	} else {
		y = w.y
	}
	return Rect{x, y, width, height}
}

func (w *windowBase) SetFrame(frame Rect) {
	handle := w.Wrapper().Handle()
	if screen := window.NSWindow_screen(handle); screen != 0 {
		frame.Y = flipWindowY(frame.Y, frame.Height, screen)
		w.y = frame.Y
	} else {
		frame.Y = w.y
	}
	window.NSWindow_setFrame(handle, frame.X, frame.Y, frame.Width, frame.Height)
}

func (w *windowBase) afterRegistered(event event.Event, nextHook event.Handler) bool {
	handle := w.Wrapper().Handle()
	// Create the content.
	content := newWindowContent()
	w.content = content
	contentHandle := content.Wrapper().Handle()
	// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSWindow_Class/#//apple_ref/occ/instp/NSWindow/contentView
	// @property(strong) __kindof NSView *contentView
	// "The window retains the new content view and owns it thereafter...Setting this property causes the old content view to be released..."
	window.NSWindow_setContentView(handle, contentHandle)
	object.NSObject_release(contentHandle)
	return nextHook(event)
}

func (w *windowBase) SetContent(content Container) {
	if w.content == content {
		return
	}
	// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSWindow_Class/#//apple_ref/occ/instp/NSWindow/contentView
	// The window retains the new content view and owns it thereafter...
	// Setting this property causes the old content view to be released; if you plan to reuse it, be sure to retain it before changing the property value
	object.NSObject_retain(w.content.Wrapper().Handle())
	window.NSWindow_setContentView(w.Wrapper().Handle(), content.Wrapper().Handle())
	w.content = content
}

func (w *windowBase) ShowModal(parent Window) interface{} {
	w.inModal = true
	defer func() {w.inModal = false}()

	handle := w.Wrapper().Handle()
	if parent == nil { // Application modal.
		app.NSApplication_runModalForWindow(app.NSApp(), handle)
	} else { // Window modal. Use sheet.
		var completed = false
		window.NSWindow_beginSheet_completionHandler(parent.Wrapper().Handle(), handle,
			func(returnCode int) {
				// Set flag.
				completed = true
				// Send an fake(empty) event to wake up the event loop below.
				//
				// Do not save this fake event in a variable and test whether the next event equals this event in the loop below.
				// NSEvent conforms to NSCopying, may be copied. Don't make pointer comparison.
				app.NSApplication_postEvent_atStart(app.NSApp(), nativeEvent.NSEvent_otherEventWithType_location_modifierFlags_timestamp_windowNumber_context_subtype_data1_data2(
					nativeEvent.NSApplicationDefined, // type
					0, 0, // location
					0,    // flags
					0,    // time
					0,    // windowNumber
					0,    // context
					0,    //subtype
					0, 0, // data1 & data2
				), false)
			})

		// Run a short circuit event loop here to wait for sheet completion. Sheet completion is asynchronous.
		for !completed {
			nextEvent := app.NSApplication_nextEventMatchingMask_untilDate_inMode_dequeue(app.NSApp(),
				nativeEvent.NSAnyEventMask, //mask
				date.NSDate_distantFuture(),  // expiration.
				runloop.NSDefaultRunLoopMode,   // mode.
				true, // flag
			)
			if nextEvent == 0 {
				break
			}
			app.NSApplication_sendEvent(app.NSApp(), nextEvent)
		}
		window.NSWindow_close(handle)
	}
	return w.modalResult
}

func (w *windowBase) endDialog(result interface{}) {
	handle := w.Wrapper().Handle()
	w.modalResult = result
	if !window.NSWindow_isSheet(handle) { // Application modal.
		if app.NSApplication_modalWindow(app.NSApp()) != handle {
			panic("It is not allowed to close a modal window other than the top one")
		}
		app.NSApplication_abortModal(app.NSApp())
		window.NSWindow_close(handle)
	} else { // Window modal.
		// sheetParent: The window to which the sheet is attached. (read-only)
		window.NSWindow_endSheet_returnCode(window.NSWindow_sheetParent(handle), handle, window.NSModalResponseStop)
	}

}

func (w *windowBase) ensureDelegate() {
	handle := w.Wrapper().Handle()
	if object.Delegate(handle) != 0 {
		return
	}
	d := dynamicinvocation.RWDynamicInvocation_initWithMethodsCallback([]string{
		"windowDidBecomeMain:", "v@:@",
		"windowWillClose:", "v@:@",
		"windowShouldClose:", "c@:@",
	}, func(selector string, args native.Handle) {
		switch selector {
		case "windowDidBecomeMain:":
			// if w.activatedEvent != nil && w.activatedEvent.HasHandler() {
			// 	w.activatedEvent.Send(&event{sender:defaultObjectTable.query(getDelegateSenderArgument(args))})
			// }
		case "windowWillClose:":
			// handle := w.Wrapper().Handle()
			// if app.NSApplication_modalWindow(app.NSApp()) == handle {
			// 	// stopModal doesn't work outside of one of the event callbacks of the modal loop.
			// 	// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSApplication_Class/index.html#//apple_ref/occ/instm/NSApplication/stopModal
			// 	app.NSApplication_abortModal(app.NSApp())
			// }
			if w.onClose.HasHandler() {
				w.onClose.Send(&simpleEvent{sender: defaultObjectTable.Query(getDelegateSenderArgument(args))})
			}
		case "windowShouldClose:":
			var yes int8 = 1
			// if w.shouldCloseEvent != nil && w.shouldCloseEvent.HasHandler() {
			// 	sender := *(*native.Handle)(dynamicinvocation.RWInvocationArguments_getArgumentAtIndex(args, 0));
			// 	if !w.shouldCloseEvent.Send(&event{sender:defaultObjectTable.query(sender)}) {
			// 		yes = 0
			// 	}
			// }
			dynamicinvocation.RWInvocationArguments_setReturnValue(args, unsafe.Pointer(&yes))
		}
	})
	object.SetDelegateRetain(handle, d)
	object.NSObject_release(d)
}

func getDelegateSenderArgument(args native.Handle) native.Handle {
	return notification.NSNotification_object(*(*native.Handle)(dynamicinvocation.RWInvocationArguments_getArgumentAtIndex(args, 0)))
}

func (w *windowBase) CloseModal(result interface{}) {
	// if w.shouldCloseEvent != nil && w.shouldCloseEvent.HasHandler() &&
	// 	w.shouldCloseEvent.Send(&event{sender: w.this()}) == false {
	// 	return
	// }
	w.endDialog(result)
}

func newWindowTemplate() Window {
	w := &windowBase{}
	w.Wrapper().AfterRegistered().AddHook(w.afterRegistered)
	return w
}

type WindowHandleManager struct {
	objcHandleManagerBase
}

func (m *WindowHandleManager) Create(util.Bundle) native.Handle {
	style := windowstyle.WindowStyle(&s.WindowStyleFeatures{
		HasBorder:         true,
		HasTitle:          true,
		HasCloseButton:    true,
		HasMinimizeButton: true,
		HasMaximizeButton: true,
		Resizable:         true,
	})
	w := window.NewRWWindow(0, 0, 300, 200, style)
	window.NSWindow_center(w)
	return deallochook.Apply(w)
}