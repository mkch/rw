package button_test

import (
	"github.com/mkch/rw"
	"github.com/mkch/rw/window"
	"github.com/mkch/rw/button"
	"github.com/mkch/rw/event"
	"testing"
	"fmt"
)

func TestBasic(t *testing.T) {
	rw.Debug = true
	rw.Run(func(){startup(t)})
}


func startup(t *testing.T) {
	win := window.New()
	btn := button.New()
	btn.SetTitle("Test 测试")
    btn.SetMnemonic('t')
	btn.SetFrame(rw.Rect{100, 100, 80, 30})
	btn.OnClick().SetHandler(func(evt event.Event) bool {
			fmt.Println("Button clicked")
            win.Content().SetTabStop(!win.Content().TabStop())
            fmt.Printf("Content tabstop=%v\n", win.Content().TabStop())
			return true
		})
    btn2 := button.New()
    btn2.SetTitle("A button with 记忆键")
    btn2.SetMnemonic('键')
    btn2.SetFrame(rw.Rect{100, 140, 200, 25})
    btn2.OnClick().SetListener(func() {
        fmt.Println("button2 clicked")
        btn2.SetTabStop(!btn2.TabStop())
        fmt.Printf("button2 tabstop=%v\n", btn2.TabStop())
    })
    win.Content().Add(btn)
    win.Content().Add(btn2)
	win.OnClose().SetHandler(func(evt event.Event) bool {
			rw.Exit()
			return true
		})
	win.ShowActive()
    btn.Focus()
}