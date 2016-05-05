package edittext

import (
	"github.com/mkch/rw"
	"github.com/mkch/rw/event"
	"fmt"
)

// HandleManager is the handle manager of EditText.
type HandleManager struct {
	rw.ControlHandleManagerBase
}

var hm = &HandleManager{}

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
