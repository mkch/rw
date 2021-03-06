package edittext_test

import (
	"fmt"
	"github.com/mkch/rw"
	"github.com/mkch/rw/edittext"
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/window"
	"testing"
)

func TestBasic(t *testing.T) {
	rw.Run(func() { startup(t) })
}

func startup(t *testing.T) {
	win := window.New()
	edt1 := edittext.New()
	edt1.SetText("Test 测试")
	edt1.SetFrame(rw.Rect{100, 100, 150, 30})
	edt1.OnChanged().SetHandler(func(evt event.Event) bool {
		fmt.Println("Edt1 changed " + evt.Sender().(edittext.EditText).Text())
		return true
	})
	win.Content().Add(edt1)
	edt2 := edittext.New()
	edt2.SetText("Test 测试 2")
	edt2.SetFrame(rw.Rect{280, 100, 150, 30})
	edt2.OnChanged().SetHandler(func(evt event.Event) bool {
		fmt.Println("Edt2 changed " + evt.Sender().(edittext.EditText).Text())
		return true
	})
	win.Content().Add(edt2)
	win.OnClose().SetListener(func() { rw.Exit() })
	win.ShowActive()
}
