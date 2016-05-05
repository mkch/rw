package window_test

import (
	"fmt"
	"github.com/mkch/rw"
	"github.com/mkch/rw/button"
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/menu"
	"github.com/mkch/rw/window"
	"image/color"
	"runtime"
	"testing"
)

func TestDialog(t *testing.T) {
	rw.Debug = true
	rw.Run(func() { dialogStartup(t) })
}

func dialogStartup(t *testing.T) {
	var win rw.Window = window.New()
	win.SetVisible(true)
	win.OnClose().SetHandler(func(event event.Event) bool {
		fmt.Printf("Sender=%v\n", event.Sender())
		rw.Exit()
		return true
	})
	win.Content().SetBackgroundColor(color.RGBA{0xFF, 0, 0xFF, 0xFF})
	win.Content().SizeChanged().SetHandler(func(event event.Event) bool {
		fmt.Printf("%v SizeChanged %v\n", event.Sender(), event.Sender().(rw.Control).Frame())
		return true
	})

	openDialogBtn := button.New()
	openDialogBtn.SetTitle("Open dialog")
	openDialogBtn.SetFrame(rw.Rect{10, 10, 160, 75})
	openDialogBtn.OnClick().SetListener(func() {
		result := newDialog().ShowModal(win)
		fmt.Printf("Dialog result=%v\n", result)
	})
	win.Content().Add(openDialogBtn)
}

func newDialog() rw.Window {
	dlg := window.New()
	dlg.SetFrame(rw.Rect{0, 0, 200, 200})
	dlg.SetTitle("This is a Dialog")
	dlg.CenterToScreen()
	endDialogBtn := button.New()
	endDialogBtn.SetTitle("End dialog")
	endDialogBtn.SetFrame(rw.Rect{5, 5, 160, 75})
	endDialogBtn.OnClick().SetListener(func() { dlg.CloseModal("closed by button") })
	dlg.Content().Add(endDialogBtn)
	dlg.Content().SetBackgroundColor(color.RGBA{0xFF, 0xFF, 0, 0})
	dlg.SetVisible(false)
	dlg.OnClose().SetListener(func() {
		fmt.Printf("Dialog on close\n")
		if runtime.GOOS == "darwin" {
			fmt.Printf("SetMainMenu to nil\n")
			setMenu(nil, nil)
		}
	})

	mainMenu := menu.NewBuilder().
		BeginItem("File").
		BeginSubmenu().
		BeginItem("Hide button").
		SetMnemonic('b').
		SetKeyboardShortcut(rw.DefaultModifierKey, 'B').
		SetOnClickListener(func() {
			fmt.Printf("Hide button menu item clicked!\n")
			endDialogBtn.SetVisible(false)
		}).
		End().
		BeginItem("Exit").
		SetMnemonic('x').
		SetKeyboardShortcut(rw.DefaultModifierKey, 'q').
		SetOnClickListener(func() {
			fmt.Printf("Exit menu item clicked!\n")
			dlg.CloseModal("closed by menu item")
		}).
		End().
		End().
		End().
		Build()

	setMenu(mainMenu, dlg)
	return dlg
}
