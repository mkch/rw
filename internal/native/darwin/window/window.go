package window

//#include "window.h"
//#include <stdlib.h>
import "C"

import (
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/util/ustr"
	"github.com/kevin-yuan/rw/internal/stackescape"
	"github.com/kevin-yuan/rw/internal/native/darwin/event"
)

func NSWindow_makeFirstResponder(w native.Handle, responder native.Handle) bool {
	return bool(C.NSWindow_makeFirstResponder(C.OBJC_PTR(w), C.OBJC_PTR(responder)))
}

func NSWindow_makeKeyAndOrderFront(w native.Handle) {
	C.NSWindow_makeKeyAndOrderFront(C.OBJC_PTR(w))
}

func NSWindow_title(w native.Handle) string {
	return C.GoString(C.NSWindow_title(C.OBJC_PTR(w)))
}

func NSWindow_setTitle(w native.Handle, title string) {
	C.NSWindow_setTitle(C.OBJC_PTR(w), (*C.char)(ustr.CStringUtf8(title)))
}

func NSWindow_contentView(w native.Handle) native.Handle {
	return native.Handle(C.NSWindow_contentView(C.OBJC_PTR(w)))
}

func NSWindow_setContentView(win, view native.Handle) {
	C.NSWindow_setContentView(C.OBJC_PTR(win), C.OBJC_PTR(view))
}

func NSWindow_frame(w native.Handle)(x, y, width, height int) {
	rect := C.NSWindow_frame(C.OBJC_PTR(w))
	return int(rect.origin.x), int(rect.origin.y), int(rect.size.width), int(rect.size.height)
}

func NSWindow_setFrame(win native.Handle, x, y, width, height int) {
	C.NSWindow_setFrameDisplay(C.OBJC_PTR(win), C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height)), true)
}

func NSWindow_center(win native.Handle) {
	C.NSWindow_center(C.OBJC_PTR(win))
}

func NSWindow_frameRectForContentRect_styleMask(contentX, contentY, contentWidth, contentHeight int, style uint)(x, y, width, height int) {
	rect := C.NSWindow_frameRectForContentRect_styleMask(C.CGRectMake(C.CGFloat(contentX), C.CGFloat(contentY), C.CGFloat(contentWidth), C.CGFloat(contentHeight)), C.UINTPTR(style))
	return int(rect.origin.x), int(rect.origin.y), int(rect.size.width), int(rect.size.height)

}

func NSWindow_contentRectForFrameRect_styleMask(frameX, frameY, frameWidth, frameHeight int, style uint)(x, y, width, height int) {
	rect := C.NSWindow_contentRectForFrameRect_styleMask(C.CGRectMake(C.CGFloat(frameX), C.CGFloat(frameY), C.CGFloat(frameWidth), C.CGFloat(frameHeight)), C.UINTPTR(style))
	return int(rect.origin.x), int(rect.origin.y), int(rect.size.width), int(rect.size.height)

}

func NSWindow_close(win native.Handle) {
	C.NSWindow_close(C.OBJC_PTR(win))
}

func NSWindow_performClose(win, sender native.Handle) {
	C.NSWindow_performClose(C.OBJC_PTR(win), C.OBJC_PTR(sender))
}

var (
	NSModalResponseStop = int(C.VarNSModalResponseStop)
	NSModalResponseAbort = int(C.VarNSModalResponseAbort)
	NSModalResponseContinue = int(C.VarNSModalResponseContinue)
)

//export goBeginSheetCompletionHandler
func goBeginSheetCompletionHandler(id C.UINTPTR, code C.int) {
	handlerId := stackescape.Id(uintptr(id))
	stackescape.Get(handlerId).(func(int))(int(code))
	stackescape.Remove(handlerId)
}

func NSWindow_beginSheet_completionHandler(w, sheet native.Handle, completionHandler func(returnCode int)) {
	var handlerId C.UINTPTR
	if completionHandler != nil {
		handlerId = C.UINTPTR(stackescape.Add(completionHandler))
	}
	C.NSWindow_beginSheet_completionHandler(C.OBJC_PTR(w), C.OBJC_PTR(sheet), handlerId)
}

func NSWindow_endSheet_returnCode(w, sheet native.Handle, returnCode int) {
	C.NSWindow_endSheet_returnCode(C.OBJC_PTR(w), C.OBJC_PTR(sheet), C.int(returnCode))
}

func NSWindow_attachedSheet(w native.Handle) native.Handle {
	return native.Handle(C.NSWindow_attachedSheet(C.OBJC_PTR(w)))
}

func NSWindow_isSheet(w native.Handle) bool {
	return bool(C.NSWindow_isSheet(C.OBJC_PTR(w)))
}

func NSWindow_sheetParent(w native.Handle) native.Handle {
	return native.Handle(C.NSWindow_sheetParent(C.OBJC_PTR(w)))
}

func NSWindow_discardEventsMatchingMask_beforeEvent(w native.Handle, eventMask event.NSEventMask, lastEvent native.Handle) {
	C.NSWindow_discardEventsMatchingMask_beforeEvent(C.OBJC_PTR(w), C.ulong(eventMask), C.OBJC_PTR(lastEvent));
}

func NSWindow_sendEvent(w, evt native.Handle) {
	C.NSWindow_sendEvent(C.OBJC_PTR(w), C.OBJC_PTR(evt))
}

func NSWindow_display(w native.Handle) {
	C.NSWindow_display((C.OBJC_PTR(w)))
}

func NSWindow_displayIfNeeded(w native.Handle) {
	C.NSWindow_displayIfNeeded(C.OBJC_PTR(w))
}

func NSWindow_isVisible(w native.Handle) bool {
	return bool(C.NSWindow_isVisible(C.OBJC_PTR(w)))
}

func NSWindow_orderOut(w native.Handle) {
	C.NSWindow_orderOut(C.OBJC_PTR(w))
}

func NSWindow_orderBack(w native.Handle) {
	C.NSWindow_orderBack(C.OBJC_PTR(w))
}

func NSWindow_orderFront(w native.Handle) {
	C.NSWindow_orderFront(C.OBJC_PTR(w))
}

func NSWindow_orderFrontRegardless(w native.Handle) {
	C.NSWindow_orderFrontRegardless(C.OBJC_PTR(w))
}

type NSWindowOrderingMode int

var (
	NSWindowAbove = NSWindowOrderingMode(C.VarNSWindowAbove)
	NSWindowBelow = NSWindowOrderingMode(C.VarNSWindowBelow)
	NSWindowOut = NSWindowOrderingMode(C.VarNSWindowOut)
)

func NSWindow_orderWindow_relativeTo(w native.Handle, mode NSWindowOrderingMode, otherWindowNumber int) {
	C.NSWindow_orderWindow_relativeTo(C.OBJC_PTR(w), C.long(mode), C.long(otherWindowNumber))
}

func NSWindow_level(w native.Handle) int {
	return int(C.NSWindow_level(C.OBJC_PTR(w)))
}

func NSWindow_setLevel(w native.Handle, level int) {
	C.NSWindow_setLevel(C.OBJC_PTR(w), C.long(level))
}

func NSWindow_firstResponder(w native.Handle) native.Handle {
	return native.Handle(C.NSWindow_firstResponder(C.OBJC_PTR(w)))
}

func NSWindow_screen(w native.Handle) native.Handle {
	return native.Handle(C.NSWindow_screen(C.OBJC_PTR(w)))
}

func initWindow(w C.OBJC_PTR, x, y, width, height int, style uint) native.Handle {
	rect := C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(width), C.CGFloat(height))
	winPtr := C.NSWindow_initWithContentRect_styleMask_backing_defer_screen(w, 
			rect, 
			C.UINTPTR(style), 
			2, // NSBackingStoreBuffered
			false, nil)
	return native.Handle(winPtr)
}

func NewWindow(x, y, width, height int, style uint) native.Handle {
	return initWindow(C.NSWindow_alloc(), x, y, width, height, style)
}

// RWWindow

func RWWindow_enabled(w native.Handle) bool {
	return bool(C.RWWindow_enabled(C.OBJC_PTR(w)))
}

func RWWindow_setEnabled(w native.Handle, enabled bool) {
	C.RWWindow_setEnabled(C.OBJC_PTR(w), C.bool(enabled))
}

func NewRWWindow(x, y, width, height int, style uint) native.Handle {
	return initWindow(C.RWWindow_alloc(), x, y, width, height, style)
}
