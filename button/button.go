package button

import (
	"fmt"
	"github.com/mkch/rw"
	"github.com/mkch/rw/event"
)

type Button interface {
	rw.Control
	SetTitle(string)
	Title() string
	// Mnemonic returns the access key(mnemonic, underlined single character) of this button. 0 for none.
	// Only used on Windows. Always returns 0 on other platforms.
	Mnemonic() rune
	// Mnemonic sets the access key(mnemonic, underlined single character). 0 for none.
	// Only used on Windows. Do nothing on other platforms.
	SetMnemonic(k rune)
	OnClick() *event.Hub
}

func (b *buttonImpl) String() string {
	if b.Wrapper().Valid() {
		return fmt.Sprintf("Button %#X %q", b.Wrapper().Handle(), b.Title())
	} else {
		return "Button <Invalid>"
	}
}
