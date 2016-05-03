package button

import (
	"github.com/kevin-yuan/rw"
	"github.com/kevin-yuan/rw/util"
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/internal/native/windows/nativeutil/ustrings"
	"github.com/kevin-yuan/rw/internal/native/windows/window"
	"github.com/kevin-yuan/rw/internal/native/windows/window/winutil"
	"github.com/kevin-yuan/rw/event"
)

var clsName ustrings.Unicode

func (m *HandleManager) Create(b util.Bundle) native.Handle {
	if clsName == nil {
		clsName = ustrings.ToUnicode("BUTTON")
	}
	return window.CreateWindowEx(0, uintptr(clsName), "Button", window.WS_CHILD|window.WS_TABSTOP|window.WS_VISIBLE, 0, 0, 100, 62, winutil.DockerWindow(), 0, window.GetModuleHandle(nil), nil)
}

type buttonImpl struct {
	rw.Control
	onClick event.Hub
    title string
    mnemonic rune
}

func (b *buttonImpl) Mnemonic() rune {
	return b.mnemonic
}

func (b *buttonImpl) SetMnemonic(k rune) {
    if k == b.mnemonic {
        return
    }
    b.mnemonic = k
    b.updateDisplayTitle()
}

func (b *buttonImpl) Title() string {
	return b.title
}

func (b *buttonImpl) SetTitle(title string) {
    if title == b.title {
        return
    }
    b.title = title
    b.updateDisplayTitle()
}

func (b* buttonImpl) updateDisplayTitle () {
    window.SetWindowText(b.Wrapper().Handle(), util.Windows_ControlTitleWithMnemonic(b.title, b.mnemonic))
}

func (b *buttonImpl) OnClick() *event.Hub {
	return &b.onClick
}

func (b *buttonImpl) Windows_ReflectedWndProc(handle native.Handle, msg uint, wParam, lParam uintptr) (result uintptr, processed bool) {
	switch msg {
	case window.WM_COMMAND:
		if b.onClick.HasHandler() {
			if b.onClick.Send(&event.SimpleEvent{b.Self()}) {
				return 0, true
			}
		}
	}
	return b.Control.Windows_ReflectedWndProc(handle, msg, wParam, lParam)
}

func New() Button {
	b := Alloc()
	rw.Init(b)
	return b
}

func Alloc() Button {
	b := &buttonImpl{Control: rw.NewControlTemplate()}
	b.Wrapper().SetHandleManager(hm)
	return b
}