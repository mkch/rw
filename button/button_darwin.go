package button

import (
	"github.com/mkch/rw"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/internal/native/darwin/object"
	"github.com/mkch/rw/internal/native/darwin/runtime"
	"github.com/mkch/rw/internal/native/darwin/button"
	"github.com/mkch/rw/internal/native/darwin/control"
	"github.com/mkch/rw/internal/native/darwin/deallochook"
	"github.com/mkch/rw/internal/native/darwin/dynamicinvocation"
	"github.com/mkch/rw/event"
)

func (m *HandleManager) Create(b util.Bundle) native.Handle {
	return deallochook.Apply(button.NewButton(0, 0, 100, 60))
}

type buttonImpl struct {
	rw.Control
	onClick event.Hub
}

func (b *buttonImpl) Mnemonic() rune {
	return 0
}

func (b *buttonImpl) SetMnemonic(k rune) {
    // Do nothing.
}

func (b *buttonImpl) Title() string {
	return button.NSButton_title(b.Wrapper().Handle())
}

func (b *buttonImpl) SetTitle(title string) {
    button.NSButton_setTitle(b.Wrapper().Handle(), title)
}

func (b *buttonImpl) ensureTargetAction() {
	handle := b.Wrapper().Handle()
	if control.NSControl_target(handle) != 0 {
		return
	}
	const targetAction = "targetAction:"
	d := dynamicinvocation.RWDynamicInvocation_initWithMethodsCallback([]string{
			targetAction, "v@:@",
		}, func(selector string, args native.Handle) {
			switch selector {
			case targetAction:
				b.onClick.Send(&event.SimpleEvent{b.Self()})
			}
			})
	object.SetTargetRetain(handle, d)
	object.NSObject_release(d)
	control.NSControl_setAction(handle, runtime.RegisterSelector(targetAction))
}


func (b *buttonImpl) OnClick() *event.Hub {
	b.ensureTargetAction()
	return &b.onClick
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