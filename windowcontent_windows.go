package rw

import (
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/internal/native/windows/window/winutil"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
	"github.com/mkch/rw/util/ustr"
	"unsafe"
)

var winContentClsName unsafe.Pointer

func createWindowContent(util.Bundle) native.Handle {
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

type winContent struct {
	Container
}

func newWindowContent() Container {
	c := &winContent{allocContainer(hwndManager(createWindowContent))}
	Init(c)
	return c
}
