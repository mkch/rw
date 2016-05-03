package panel_test

import (
	"github.com/kevin-yuan/rw"
	"github.com/kevin-yuan/rw/window"
	"github.com/kevin-yuan/rw/button"
	"github.com/kevin-yuan/rw/menu"
	"github.com/kevin-yuan/rw/panel"
	//"github.com/kevin-yuan/rw/event"
	"testing"
	"image/color"
)

func TestTabOrder(t *testing.T) {
	rw.Run(func(){startup(t)})
}


func startup(t *testing.T) {
	win := window.New()
	win.OnClose().SetListener(func(){rw.Exit()})

	btn1 := button.New()
	btn1.SetTitle("1")
	btn1.SetFrame(rw.Rect{10, 10, 100, 30})

	btn2 := button.New()
	btn2.SetTitle("2")
	btn2.SetFrame(rw.Rect{10, 45, 100, 30})

	btn3 := button.New()
	btn3.SetTitle("3")
	btn3.SetFrame(rw.Rect{10, 80, 100, 30})

	btn4 := button.New()
	btn4.SetTitle("4")
	btn4.SetFrame(rw.Rect{10, 115, 100, 30})

	p := panel.New()
	p.SetFrame(rw.Rect{20, 150, 120, 120})
	p.SetBackgroundColor(color.RGBA{0xFF, 0xFF, 0, 0xFF})

	btn5 := button.New()
	btn5.SetTitle("5")
	btn5.SetFrame(rw.Rect{5, 5, 100, 30})
	p.Add(btn5)

	content := win.Content()
	content.Add(btn1)
	content.Add(btn2)
	content.Add(btn3)
	content.Add(btn4)
	content.Add(p)


	m := menu.NewBuilder().
		BeginItem("Ops").
			BeginSubmenu().
				BeginItem("5 4 3 2 1").
				SetOnClickListener(func(){
					btn4.SetTabOrder(1)
					btn3.SetTabOrder(2)
					btn2.SetTabOrder(3)
					btn1.SetTabOrder(4)
				}).
				End().
				BeginItem("5 1 4 2 3").
				SetOnClickListener(func(){
					btn5.SetTabOrder(8)
					btn1.SetTabOrder(100)
					btn2.SetTabOrder(111)
					btn3.SetTabOrder(111)
					btn4.SetTabOrder(100)
				}).
				End().
				BeginItem("2.SetTabStop(false)").
				SetOnClickListener(func(){
					btn2.SetTabStop(false)
				}).
				End().
				BeginItem("2.SetTabStop(true)").
				SetOnClickListener(func(){
					btn2.SetTabStop(true)
				}).
				End().
				BeginItem("panel.SetTabStop(true)").
				SetOnClickListener(func(){
					p.SetTabStop(true)
				}).
				End().
				BeginItem("panel.SetTabStop(false)").
				SetOnClickListener(func(){
					p.SetTabStop(false)
				}).
				End().
			End().
		End().
	Build()

	win.SetMenu(m)
	rw.SetMainMenu(m)
	win.ShowActive()
}