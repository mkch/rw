// Package alert implements standard alert/message dialog of the operating system.
package alert

import (
	"github.com/mkch/rw"
)

type Style int

const (
	// Informational alert is used to inform the user about a current or impending event.
	Informational Style = iota
	// Warning alert is used to warn the user about a current or impending event.
	Warning
	// Error alert is used to report the user about an error.
	Error
	// Critical alert is used to report the user about an critical event.
	Critical
)

type Alert interface {
	// Show shows the alert dialog. The dialog is window modal if parent is not nil, or application modal.
	Show(parent rw.Window)
	// SetStyle sets the style of the alert dialog.
	SetStyle(style Style) Alert
	// SetTitle sets the window title of the alert dialog. The title is ignored on Mac OS X.
	SetTitle(title string) Alert
	// SetMessage sets the message to display in the alert dialog.
	SetMessage(message string) Alert
	// SetInformativeMessage optionally sets the informative(secondary) message to display in the alert dialog.
	SetInformativeMessage(message string) Alert
	// SetPositiveButton sets the title and a onclick handler(can be nil) of the positive button. If the parameter asDefault
	// is true, this button will be the default button.
	SetPositiveButton(title string, handler func(), asDefault bool) Alert
	// SetNegativeButton sets the title and a onclick handler(can be nil) of the negative button. If the parameter asDefault
	// is true, this button will be the default button.
	SetNegativeButton(title string, handler func(), asDefault bool) Alert
	// SetNeutralButton sets the title and a onclick handler(can be nil) of the neutral button. If the parameter asDefault
	// is true, this button will be the default button.
	SetNeutralButton(title string, handler func(), asDefault bool) Alert
}

// New returns a new Alert.
func New() Alert {
	return newAlert()
}

// Show shows an application modal alert dialog.
//  alert.Show(message)
// is equivalent to
//	alert.New().SetMessage(message).Show(nil)
func Show(message string) {
	New().SetMessage(message).Show(nil)
}
