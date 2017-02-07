package window_test

import (
	"fmt"
	"github.com/mkch/rw"
	"github.com/mkch/rw/alert"
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/window"
	"image/color"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	rw.Debug = true
	rw.Run(func() { startup(t) })
}

type Window struct {
	rw.Window
}

func NewWindow() *Window {
	win := &Window{Window: window.Alloc()}
	rw.Init(win)

	win.Content().SetBackgroundColor(color.RGBA{0xFF, 0, 0, 0xFF})
	return win
}

func startup(t *testing.T) {
	alert.Show("Hello")
	var win rw.Window = NewWindow()
	win.SetVisible(true)
	if win.Enabled() != true {
		t.Errorf("Window is not initially enabled.")
	}
	win.SetEnabled(false)
	if win.Enabled() != false {
		t.Errorf("Disable window failed.")
	}
	go func() {
		time.Sleep(time.Second * 3)
		rw.Post(func() {
			win.SetEnabled(true)
			if win.Enabled() != true {
				t.Errorf("Disabling window failed.")
			}
			win.Content().SetBackgroundColor(color.RGBA{0xFF, 0xFF, 0, 0xFF})
		})
	}()
	win.OnClose().SetHandler(func(event event.Event) bool {
		fmt.Printf("Sender=%v\n", event.Sender())
		rw.Exit()
		return true
	})
	win.Content().SizeChanged().SetHandler(func(event event.Event) bool {
		fmt.Printf("%v SizeChanged %v\n", event.Sender(), event.Sender().(rw.Control).Frame())
		return true
	})

}
