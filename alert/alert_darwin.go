package alert

import (
	"github.com/mkch/rw"
	nativeAlert "github.com/mkch/rw/internal/native/darwin/alert"
	"github.com/mkch/rw/internal/native/darwin/object"
	"github.com/mkch/rw/internal/native/darwin/util"
	"sort"
)

type alertData struct {
	style              Style
	message            string
	informativeMessage string
	positive           button
	negative           button
	neutral            button
}

type button struct {
	title   string
	handler func()
	def     bool
	value   int
}

func (a *alertData) SetStyle(style Style) Alert {
	a.style = style
	return a
}

func (a *alertData) SetTitle(title string) Alert {
	// Do nothing.
	return a
}

func (a *alertData) SetMessage(message string) Alert {
	a.message = message
	return a
}

func (a *alertData) SetInformativeMessage(message string) Alert {
	a.informativeMessage = message
	return a
}

func (a *alertData) SetPositiveButton(title string, handler func(), asDefault bool) Alert {
	a.positive.title = title
	a.positive.handler = handler
	a.positive.def = asDefault
	if asDefault {
		if a.negative.def || a.neutral.def {
			panic("Default button is already set")
		}
		a.positive.value = nativeAlert.NSAlertFirstButtonReturn
		a.negative.value = nativeAlert.NSAlertFirstButtonReturn + 1
		a.neutral.value = nativeAlert.NSAlertFirstButtonReturn + 2
	}
	return a
}

func (a *alertData) SetNegativeButton(title string, handler func(), asDefault bool) Alert {
	a.negative.title = title
	a.negative.handler = handler
	a.negative.def = asDefault
	if asDefault {
		if a.positive.def || a.neutral.def {
			panic("Default button is already set")
		}
		a.positive.value = nativeAlert.NSAlertFirstButtonReturn + 1
		a.negative.value = nativeAlert.NSAlertFirstButtonReturn
		a.neutral.value = nativeAlert.NSAlertFirstButtonReturn + 2
	}
	return a
}

func (a *alertData) SetNeutralButton(title string, handler func(), asDefault bool) Alert {
	a.neutral.title = title
	a.neutral.handler = handler
	a.neutral.def = asDefault
	if asDefault {
		if a.positive.def || a.negative.def {
			panic("Default button is already set")
		}
		a.positive.value = nativeAlert.NSAlertFirstButtonReturn + 1
		a.negative.value = nativeAlert.NSAlertFirstButtonReturn + 2
		a.neutral.value = nativeAlert.NSAlertFirstButtonReturn
	}
	return a
}

type buttonsByValue []*button

func (a buttonsByValue) Len() int {
	return len(a)
}

func (a buttonsByValue) Less(i, j int) bool {
	return a[i].value < a[j].value
}

func (a buttonsByValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a *alertData) Show(parent rw.Window) {
	na := object.NSObject_autorelease(object.NSObject_init(nativeAlert.NSAlert_alloc()))
	// Messages
	nativeAlert.NSAlert_setMessageText(na, a.message)
	nativeAlert.NSAlert_setInformativeText(na, a.informativeMessage)
	// Style
	var style = nativeAlert.NSInformationalAlertStyle
	switch a.style {
	case Informational:
		style = nativeAlert.NSInformationalAlertStyle
	case Warning, Error:
		style = nativeAlert.NSWarningAlertStyle
	case Critical:
		style = nativeAlert.NSCriticalAlertStyle
	default:
		panic("Invalid alert style")
	}
	nativeAlert.NSAlert_setAlertStyle(na, style)
	// Buttons
	buttons := []*button{&a.positive, &a.negative, &a.neutral}
	sort.Sort(buttonsByValue(buttons)) // The first button is the default button.
	for _, b := range buttons {
		nativeAlert.NSAlert_addButtonWithTitle(na, b.title)
	}

	var returnCode = nativeAlert.NSAlertFirstButtonReturn
	if parent == nil {
		returnCode = nativeAlert.NSAlert_runModal(na)
	} else {
		var rsc util.RunloopShortCircuiter
		nativeAlert.NSAlert_beginSheetModalForWindow_completionHandler(na, parent.Wrapper().Handle(), func(code int) {
			returnCode = code
			rsc.Stop()
		})
		rsc.Run()
	}
	if i := sort.Search(len(buttons), func(i int) bool { return buttons[i].value >= returnCode }); i != len(buttons) {
		if buttons[i].handler != nil {
			buttons[i].handler()
		}
	}
}

func newAlert() Alert {
	return &alertData{
		style:    Informational,
		positive: button{value: nativeAlert.NSAlertFirstButtonReturn},
		negative: button{value: nativeAlert.NSAlertFirstButtonReturn + 1},
		neutral:  button{value: nativeAlert.NSAlertFirstButtonReturn + 2},
	}
}
