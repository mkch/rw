package window_test

import (
	"fmt"
	"github.com/kevin-yuan/rw"
	"github.com/kevin-yuan/rw/event"
	"github.com/kevin-yuan/rw/window"
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
	var win rw.Window = NewWindow()
	win.SetVisible(true)
	win.SetEnabled(false)
	fmt.Printf("Enabled: %v\n", win.Enabled())
	go func() {
		time.Sleep(time.Second * 3)
		rw.Post(func() {
			win.SetEnabled(true)
			win.Content().SetBackgroundColor(color.RGBA{0xFF, 0xFF, 0, 0xFF})
			fmt.Printf("Enabled: %v\n", win.Enabled())
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
