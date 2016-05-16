package alert

import (
	"github.com/mkch/rw"
	"github.com/mkch/rw/internal/native/windows/messagebox"
	"github.com/mkch/rw/native"
)

const (
	Abort    string = "Abort"
	Cancel          = "Cancel"
	Continue        = "Continue"
	Ignore          = "Ignore"
	No              = "No"
	OK              = "OK"
	Retry           = "Retry"
	TryAgain        = "Try Again"
	Yes             = "Yes"
)

var returnCodes = map[string]messagebox.ReturnCode{
	Abort:    messagebox.IDABORT,
	Cancel:   messagebox.IDCANCEL,
	Continue: messagebox.IDCONTINUE,
	Ignore:   messagebox.IDIGNORE,
	No:       messagebox.IDNO,
	OK:       messagebox.IDOK,
	Retry:    messagebox.IDRETRY,
	TryAgain: messagebox.IDTRYAGAIN,
	Yes:      messagebox.IDYES,
}

type alertData struct {
	caption            string
	text               string
	informativeMessage string
	t                  messagebox.Type
	buttons            [3]button
}

type button struct {
	code    messagebox.ReturnCode
	handler func()
	def     bool
}

func (a *alertData) SetStyle(style Style) Alert {
	switch style {
	case Informational:
		a.t = messagebox.MB_ICONINFORMATION
	case Warning, Critical:
		a.t = messagebox.MB_ICONWARNING
	case Error:
		a.t = messagebox.MB_ICONERROR
	}
	return a
}

func (a *alertData) SetTitle(title string) Alert {
	a.caption = title
	return a
}

func (a *alertData) SetMessage(message string) Alert {
	a.text = message
	return a
}

func (a *alertData) SetInformativeMessage(message string) Alert {
	a.informativeMessage = message
	return a
}

func (a *alertData) SetPositiveButton(title string, handler func(), asDefault bool) Alert {
	var ok bool
	if a.buttons[0].code, ok = returnCodes[title]; !ok {
		panic("Invalid button title")
	}
	if a.buttons[1].code == a.buttons[0].code || a.buttons[2].code == a.buttons[0].code {
		panic("Duplicated button")
	}
	a.buttons[0].handler = handler
	if asDefault {
		if a.buttons[1].def || a.buttons[2].def {
			panic("Default button is already set")
		}
		a.buttons[0].def = true
	}
	return a
}

func (a *alertData) SetNegativeButton(title string, handler func(), asDefault bool) Alert {
	var ok bool
	if a.buttons[1].code, ok = returnCodes[title]; !ok {
		panic("Invalid button title")
	}
	if a.buttons[0].code == a.buttons[1].code || a.buttons[2].code == a.buttons[1].code {
		panic("Duplicated button")
	}
	a.buttons[1].handler = handler
	if asDefault {
		if a.buttons[0].def || a.buttons[2].def {
			panic("Default button is already set")
		}
		a.buttons[1].def = true
	}
	return a
}

func (a *alertData) SetNeutralButton(title string, handler func(), asDefault bool) Alert {
	var ok bool
	if a.buttons[2].code, ok = returnCodes[title]; !ok {
		panic("Invalid button title")
	}
	if a.buttons[0].code == a.buttons[2].code || a.buttons[1].code == a.buttons[2].code {
		panic("Duplicated button")
	}
	a.buttons[2].handler = handler
	if asDefault {
		if a.buttons[0].def || a.buttons[1].def {
			panic("Default button is already set")
		}
		a.buttons[2].def = true
	}
	return a
}

func badButtons() {
	panic("Invalid buttons")
}

func (a *alertData) Show(parent rw.Window) {
	var defCode messagebox.ReturnCode
	if a.buttons[0].def {
		defCode = a.buttons[0].code
	} else if a.buttons[1].def {
		defCode = a.buttons[1].code
	} else if a.buttons[2].def {
		defCode = a.buttons[2].code
	}

	var code0 = a.buttons[0].code
	var code1 = a.buttons[1].code
	var code2 = a.buttons[2].code
	if code0 == messagebox.IDABORT && code1 == messagebox.IDRETRY && code2 == messagebox.IDIGNORE ||
		code0 == messagebox.IDABORT && code2 == messagebox.IDRETRY && code1 == messagebox.IDIGNORE ||
		code1 == messagebox.IDABORT && code0 == messagebox.IDRETRY && code2 == messagebox.IDIGNORE ||
		code1 == messagebox.IDABORT && code2 == messagebox.IDRETRY && code0 == messagebox.IDIGNORE ||
		code2 == messagebox.IDABORT && code0 == messagebox.IDRETRY && code1 == messagebox.IDIGNORE ||
		code2 == messagebox.IDABORT && code1 == messagebox.IDRETRY && code0 == messagebox.IDIGNORE {
		a.t |= messagebox.MB_ABORTRETRYIGNORE
		switch defCode {
		case messagebox.IDABORT:
			a.t |= messagebox.MB_DEFBUTTON1
		case messagebox.IDRETRY:
			a.t |= messagebox.MB_DEFBUTTON2
		case messagebox.IDIGNORE:
			a.t |= messagebox.MB_DEFBUTTON3
		}
	} else if code0 == messagebox.IDCANCEL && code1 == messagebox.IDTRYAGAIN && code2 == messagebox.IDCONTINUE ||
		code0 == messagebox.IDCANCEL && code2 == messagebox.IDTRYAGAIN && code1 == messagebox.IDTRYAGAIN ||
		code1 == messagebox.IDCANCEL && code0 == messagebox.IDTRYAGAIN && code2 == messagebox.IDTRYAGAIN ||
		code1 == messagebox.IDCANCEL && code2 == messagebox.IDTRYAGAIN && code0 == messagebox.IDTRYAGAIN ||
		code2 == messagebox.IDCANCEL && code0 == messagebox.IDTRYAGAIN && code1 == messagebox.IDTRYAGAIN ||
		code2 == messagebox.IDCANCEL && code1 == messagebox.IDTRYAGAIN && code0 == messagebox.IDTRYAGAIN {
		a.t |= messagebox.MB_CANCELTRYCONTINUE
		switch defCode {
		case messagebox.IDCANCEL:
			a.t |= messagebox.MB_DEFBUTTON1
		case messagebox.IDTRYAGAIN:
			a.t |= messagebox.MB_DEFBUTTON2
		case messagebox.IDCONTINUE:
			a.t |= messagebox.MB_DEFBUTTON3
		}

	} else if code0 == messagebox.IDOK && code1 == 0 && code2 == 0 ||
		code0 == 0 && code1 == messagebox.IDOK && code2 == 0 ||
		code0 == 0 && code1 == 0 && code2 == messagebox.IDOK {
		a.t |= messagebox.MB_OK
	} else if code0 == messagebox.IDOK && code1 == messagebox.IDCANCEL && code2 == 0 ||
		code0 == messagebox.IDOK && code2 == messagebox.IDCANCEL && code1 == 0 ||
		code1 == messagebox.IDOK && code0 == messagebox.IDCANCEL && code2 == 0 ||
		code1 == messagebox.IDOK && code2 == messagebox.IDCANCEL && code0 == 0 ||
		code2 == messagebox.IDOK && code0 == messagebox.IDCANCEL && code1 == 0 ||
		code2 == messagebox.IDOK && code1 == messagebox.IDCANCEL && code0 == 0 {
		a.t |= messagebox.MB_OKCANCEL
		switch defCode {
		case messagebox.IDOK:
			a.t |= messagebox.MB_DEFBUTTON1
		case messagebox.IDCANCEL:
			a.t |= messagebox.MB_DEFBUTTON2

		}
	} else if code0 == messagebox.IDRETRY && code1 == messagebox.IDCANCEL && code2 == 0 ||
		code0 == messagebox.IDRETRY && code2 == messagebox.IDCANCEL && code1 == 0 ||
		code1 == messagebox.IDRETRY && code0 == messagebox.IDCANCEL && code2 == 0 ||
		code1 == messagebox.IDRETRY && code2 == messagebox.IDCANCEL && code0 == 0 ||
		code2 == messagebox.IDRETRY && code0 == messagebox.IDCANCEL && code1 == 0 ||
		code2 == messagebox.IDRETRY && code1 == messagebox.IDCANCEL && code0 == 0 {
		a.t |= messagebox.MB_RETRYCANCEL
		switch defCode {
		case messagebox.IDRETRY:
			a.t |= messagebox.MB_DEFBUTTON1
		case messagebox.IDCANCEL:
			a.t |= messagebox.MB_DEFBUTTON2

		}
	} else if code0 == messagebox.IDYES && code1 == messagebox.IDNO && code2 == 0 ||
		code0 == messagebox.IDYES && code2 == messagebox.IDNO && code1 == 0 ||
		code1 == messagebox.IDYES && code0 == messagebox.IDNO && code2 == 0 ||
		code1 == messagebox.IDYES && code2 == messagebox.IDNO && code0 == 0 ||
		code2 == messagebox.IDYES && code0 == messagebox.IDNO && code1 == 0 ||
		code2 == messagebox.IDYES && code1 == messagebox.IDNO && code0 == 0 {
		a.t |= messagebox.MB_YESNO
		switch defCode {
		case messagebox.IDYES:
			a.t |= messagebox.MB_DEFBUTTON1
		case messagebox.IDNO:
			a.t |= messagebox.MB_DEFBUTTON2

		}
	} else if code0 == messagebox.IDYES && code1 == messagebox.IDNO && code2 == messagebox.IDCANCEL ||
		code0 == messagebox.IDYES && code2 == messagebox.IDNO && code1 == messagebox.IDCANCEL ||
		code1 == messagebox.IDYES && code0 == messagebox.IDNO && code2 == messagebox.IDCANCEL ||
		code1 == messagebox.IDYES && code2 == messagebox.IDNO && code0 == messagebox.IDCANCEL ||
		code2 == messagebox.IDYES && code0 == messagebox.IDNO && code1 == messagebox.IDCANCEL ||
		code2 == messagebox.IDYES && code1 == messagebox.IDNO && code0 == messagebox.IDCANCEL {
		a.t |= messagebox.MB_YESNOCANCEL
		switch defCode {
		case messagebox.IDYES:
			a.t |= messagebox.MB_DEFBUTTON1
		case messagebox.IDNO:
			a.t |= messagebox.MB_DEFBUTTON2
		case messagebox.IDCANCEL:
			a.t |= messagebox.MB_DEFBUTTON3

		}
	} else {
		badButtons()
	}

	var parentHandle native.Handle
	if parent != nil {
		parentHandle = parent.Wrapper().Handle()
	}
	msg := a.text
	if len(a.informativeMessage) > 0 {
		msg += ("\n\n" + a.informativeMessage)
	}
	ret := messagebox.MessageBox(parentHandle, msg, a.caption, a.t)
	switch ret {
	case a.buttons[0].code:
		if a.buttons[0].handler != nil {
			a.buttons[0].handler()
		}
	case a.buttons[1].code:
		if a.buttons[1].handler != nil {
			a.buttons[1].handler()
		}
	case a.buttons[2].code:
		if a.buttons[2].handler != nil {
			a.buttons[2].handler()
		}
	default:
		panic("Unexpected return code")
	}
}

func newAlert() Alert {
	return &alertData{}
}
