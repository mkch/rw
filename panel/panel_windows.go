package panel

import (
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/util"
	"github.com/kevin-yuan/rw/util/ustr"
	"github.com/kevin-yuan/rw/internal/native/windows/window"
	"github.com/kevin-yuan/rw/internal/native/windows/window/winutil"
)

var clsName unsafe.Pointer

func (m *HandleManager) Create(util.Bundle) native.Handle {
	moduleHandle := window.GetModuleHandle(nil)
	if clsName == nil {
		clsName = ustr.CStringUtf16("rw.Panel")
		window.RegisterClassEx(&window.WndClassEx {
			WndProc: window.DefWindowProcPtr(),
			Instance: moduleHandle,
			Cursor: window.LoadCursor(0, window.IDC_ARROW),
			Background: native.Handle(window.COLOR_WINDOW),
			ClassName:	clsName,
			})
	}
	return window.CreateWindowEx(window.WS_EX_CONTROLPARENT, uintptr(clsName), "", window.WS_CHILD|window.WS_TABSTOP|window.WS_VISIBLE, 0, 0, 100, 200, winutil.DockerWindow(), 0, moduleHandle, nil)
}