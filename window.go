package rw

import (
	"fmt"
	"github.com/kevin-yuan/rw/event"
	"github.com/kevin-yuan/rw/util"
)

// Window is the base interface for all windows.
// A Window is released when it is closed. Window releases it's Content and Menu when it is released.
type Window interface {
	Object
	// ShowActive makes the window visible and brings it to front.
	ShowActive()
	// Close closes the window.
	Close()
	// Visible returns whether the window is visible.
	Visible() bool
	// SetVisible sets the visibility of the window.
	SetVisible(v bool)
	// Content returns the Container that occupys whole content area of the window.
	// Content of a window resizes automatically when the window resizes.
	Content() Container
	// SetContent sets the content of this window.
	// The previous content is not released.
	SetContent(content Container)
	// SetFrame sets the position and size of the window, in screen coordinates.
	SetFrame(frame Rect)
	// Frame returns the position and size of the window, in screen coordinates.
	Frame() Rect
	// OnColse is an event hub where an event is sent before the window is closed.
	OnClose() *event.Hub
	// Menu returns the menu(bar) of this window. Does nothing on Mac OS X platform.
	Menu() Menu
	// Set the menu(bar) of this window. Only makes sense on Windows platform.
	SetMenu(menu Menu)
	// Title returns the title text of this window.
	Title() string
	// SetTitle sets the title text of this window.
	SetTitle(title string)
	Enabled() bool
	SetEnabled(enabled bool)
	// CenterToScreen centers the window to the screen.
	CenterToScreen()
	// Show the window as a modal dialog.
	// This method will not return until CloseModal is called. The return value is the result provided by CloseModal.
	ShowModal(parent Window) interface{}
	// CloseModal closes the modal dialog, and let ShowModal return result. Panics if ShowModal is not called on this window.
	// This method closes all dialogs created by this dialog, with return value nil, before closes this dialog itself.
	CloseModal(result interface{})

	util.WrapperHolder

	WindowPlatformSpecific

	windowExtra
}

func (w *windowBase) Content() Container {
	return w.content
}

func (w *windowBase) Wrapper() util.Wrapper {
	return &w.wrapper
}

func (w *windowBase) Release() {
	util.Release(w)
}

func (w *windowBase) String() string {
	if w.Wrapper().Valid() {
		return fmt.Sprintf("Window %#X %q", w.Wrapper().Handle(), w.Title())
	} else {
		return "Window <Invalid>"
	}
}

// NewWindowTemplate creates a template of Window.
// Use corresponding NewXxx functions(NewWindow in package rw/Window for example) to create objects of concrete Window.
func NewWindowTemplate() Window {
	return newWindowTemplate()
}
