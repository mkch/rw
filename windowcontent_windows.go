package rw

import (
	"github.com/kevin-yuan/rw/util/ustr"
	"github.com/kevin-yuan/rw/internal/native/windows/window"
	"github.com/kevin-yuan/rw/internal/native/windows/window/winutil"
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/util"
)

type winContentHandleManager struct {
	hwndManagerBase
}

var winContentClsName unsafe.Pointer

func (m winContentHandleManager) Create(util.Bundle) native.Handle {
	moduleHandle := window.GetModuleHandle(nil)
	if winContentClsName == nil {
		winContentClsName = ustr.CStringUtf16("rw.WindowContent")
		window.RegisterClassEx(&window.WndClassEx{
			WndProc:    window.DefWindowProcPtr(),
			Instance:   moduleHandle,
			Cursor:     window.LoadCursor(0, window.IDC_ARROW),
			Background: native.Handle(window.COLOR_WINDOW),
			ClassName:  winContentClsName,
		})
	}
	return window.CreateWindowEx(window.WS_EX_CONTROLPARENT, uintptr(winContentClsName), "", window.WS_CHILD|window.WS_VISIBLE, 0, 0, 100, 200, winutil.DockerWindow(), 0, moduleHandle, nil)
}

var winContentHM = winContentHandleManager{}

type winContent struct {
	Container
}

func newWindowContent() Container {
	c := &winContent{NewContainerTemplate()}
	c.Wrapper().SetHandleManager(winContentHM)
	Init(c)
	return c
}
