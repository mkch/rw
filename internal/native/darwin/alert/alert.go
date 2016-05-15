package alert

//#include "alert.h"
import "C"

import (
	"github.com/mkch/rw/internal/stackescape"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util/ustr"
)

type NSAlertStyle uint

var (
	NSWarningAlertStyle       = NSAlertStyle(C.VarNSWarningAlertStyle)
	NSInformationalAlertStyle = NSAlertStyle(C.VarNSInformationalAlertStyle)
	NSCriticalAlertStyle      = NSAlertStyle(C.VarNSCriticalAlertStyle)
)

func NSAlert_alloc() native.Handle {
	return native.Handle(C.NSAlert_alloc())
}

func NSAlert_alertWithError(a native.Handle) native.Handle {
	return native.Handle(C.NSAlert_alertWithError(C.OBJC_PTR(a)))
}

func NSAlert_layout(a native.Handle) {
	C.NSAlert_layout(C.OBJC_PTR(a))
}

func NSAlert_alertStyle(a native.Handle) NSAlertStyle {
	return NSAlertStyle(C.NSAlert_alertStyle(C.OBJC_PTR(a)))
}

func NSAlert_setAlertStyle(a native.Handle, style NSAlertStyle) {
	C.NSAlert_setAlertStyle(C.OBJC_PTR(a), C.NSUInteger(style))
}

func NSAlert_accessoryView(a native.Handle) native.Handle {
	return native.Handle(C.NSAlert_accessoryView(C.OBJC_PTR(a)))
}

func NSAlert_setAccessoryView(a native.Handle, view native.Handle) {
	C.NSAlert_setAccessoryView(C.OBJC_PTR(a), C.OBJC_PTR(view))
}

func NSAlert_showsHelp(a native.Handle) bool {
	return bool(C.NSAlert_showsHelp(C.OBJC_PTR(a)))
}

func NSAlert_setShowsHelp(a native.Handle, show bool) {
	C.NSAlert_setShowsHelp(C.OBJC_PTR(a), C.bool(show))
}

func NSAlert_helpAnchor(a native.Handle) string {
	return C.GoString(C.NSAlert_helpAnchor(C.OBJC_PTR(a)))
}

func NSAlert_setHelpAnchor(a native.Handle, anchor string) {
	C.NSAlert_setHelpAnchor(C.OBJC_PTR(a), (*C.char)(ustr.CStringUtf8(anchor)))
}

func NSAlert_delegate(a native.Handle) native.Handle {
	return native.Handle(C.NSAlert_delegate(C.OBJC_PTR(a)))
}

func NSAlert_setDelegate(a, delegate native.Handle) {
	C.NSAlert_setDelegate(C.OBJC_PTR(a), C.OBJC_PTR(delegate))
}

func NSAlert_runModal(a native.Handle) int {
	return int(C.NSAlert_runModal(C.OBJC_PTR(a)))
}

//export goBeginSheetModalForWindowCompletionHandler
func goBeginSheetModalForWindowCompletionHandler(id C.UINTPTR, returnCode C.int) {
	handlerId := stackescape.Id(uintptr(id))
	stackescape.Get(handlerId).(func(int))(int(returnCode))
	stackescape.Remove(handlerId)
}

func NSAlert_beginSheetModalForWindow_completionHandler(a native.Handle, sheetWindow native.Handle, completionHandler func(returnCode int)) {
	var handlerId C.UINTPTR
	if completionHandler != nil {
		handlerId = C.UINTPTR(stackescape.Add(completionHandler))
	}
	C.NSAlert_beginSheetModalForWindow_completionHandler(C.OBJC_PTR(a), C.OBJC_PTR(sheetWindow), handlerId)
}

func NSAlert_suppressionButton(a native.Handle) native.Handle {
	return native.Handle(C.NSAlert_suppressionButton(C.OBJC_PTR(a)))
}

func NSAlert_showsSuppressionButton(a native.Handle) bool {
	return bool(C.NSAlert_showsSuppressionButton(C.OBJC_PTR(a)))
}

func NSAlert_setShowsSuppressionButton(a native.Handle, show bool) {
	C.NSAlert_setShowsSuppressionButton(C.OBJC_PTR(a), C.bool(show))
}

func NSAlert_informativeText(a native.Handle) string {
	return C.GoString(C.NSAlert_informativeText(C.OBJC_PTR(a)))
}

func NSAlert_setInformativeText(a native.Handle, text string) {
	C.NSAlert_setInformativeText(C.OBJC_PTR(a), (*C.char)(ustr.CStringUtf8(text)))
}

func NSAlert_messageText(a native.Handle) string {
	return C.GoString(C.NSAlert_messageText(C.OBJC_PTR(a)))
}

func NSAlert_setMessageText(a native.Handle, text string) {
	C.NSAlert_setMessageText(C.OBJC_PTR(a), (*C.char)(ustr.CStringUtf8(text)))
}

func NSAlert_icon(a native.Handle) native.Handle {
	return native.Handle(C.NSAlert_icon(C.OBJC_PTR(a)))
}

func NSAlert_setIcon(a native.Handle, icon native.Handle) {
	C.NSAlert_setIcon(C.OBJC_PTR(a), C.OBJC_PTR(icon))
}

func NSAlert_buttons(a native.Handle) native.Handle {
	return native.Handle(C.NSAlert_buttons(C.OBJC_PTR(a)))
}

func NSAlert_addButtonWithTitle(a native.Handle, title string) {
	C.NSAlert_addButtonWithTitle(C.OBJC_PTR(a), (*C.char)(ustr.CStringUtf8(title)))
}

func NSAlert_window(a native.Handle) native.Handle {
	return native.Handle(C.NSAlert_window(C.OBJC_PTR(a)))
}
