package edittext

import (
	"fmt"
	"github.com/mkch/rw"
	"github.com/mkch/rw/event"
)

// EditText is a Control where you can type text.
type EditText interface {
	rw.Control
	Text() string
	SetText(text string)
	OnChanged() *event.Hub
}

func (edt *editTextImpl) String() string {
	if edt.Wrapper().Valid() {
		return fmt.Sprintf("EditText %#X %q", edt.Wrapper().Handle(), edt.Text())
	} else {
		return "EditText <Invalid>"
	}
}
