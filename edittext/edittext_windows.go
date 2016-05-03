package edittext

import (
	"github.com/kevin-yuan/rw"
	"github.com/kevin-yuan/rw/util"
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/internal/native/windows/window"
	"github.com/kevin-yuan/rw/internal/native/windows/edit"
	"github.com/kevin-yuan/rw/internal/native/windows/window/winutil"
	"github.com/kevin-yuan/rw/internal/native/windows/commcontrol"
	"github.com/kevin-yuan/rw/event"
)

func (m *HandleManager) Create(b util.Bundle) native.Handle {
	return window.CreateWindowEx(
		/*exStyle uint*/ window.WS_EX_CLIENTEDGE,
		/*className string*/ commcontrol.WC_EDIT,
		/*windowName string*/ "",
		/*style uint*/ window.WS_CHILD|window.WS_VISIBLE|window.WS_BORDER|window.WS_TABSTOP,
		/*x int*/ 0,
		/*y int*/ 0,
		/*width int*/ 80,
		/*height int*/ 30,
		/*parent Handle*/ winutil.DockerWindow(),
		/*menu Handle*/ 0,
		/*instance Handle*/ window.GetModuleHandle(nil),
		/*lParam unsafe.Pointer*/ nil)
}

type editTextImpl struct {
	rw.Control
	onChanged event.Hub
}

func (edt *editTextImpl) Text() string {
	return window.GetWindowText(edt.Wrapper().Handle())
}

func (edt *editTextImpl) SetText(text string) {
	window.SetWindowText(edt.Wrapper().Handle(), text)
}

func (edt *editTextImpl) OnChanged() *event.Hub {
	return &edt.onChanged
}

func (edt *editTextImpl) Windows_ReflectedWndProc(handle native.Handle, msg uint, wParam, lParam uintptr) (result uintptr, processed bool) {
	switch msg {
	case window.WM_COMMAND:
		switch uint(window.HIWORD(uint(wParam))) {
		case edit.EN_CHANGE:
			if edt.onChanged.HasHandler() {
				edt.onChanged.Send(&event.SimpleEvent{edt.Self()})
			}
		}
	}
	return edt.Control.Windows_ReflectedWndProc(handle, msg, wParam, lParam)
}

func Alloc() EditText {
	edt := &editTextImpl{Control: rw.NewControlTemplate()}
	edt.Wrapper().SetHandleManager(hm)
	return edt
}

func New() EditText {
	edt := Alloc()
	rw.Init(edt)
	return edt
}
