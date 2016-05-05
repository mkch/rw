package rw

import (
	"github.com/mkch/rw/internal/native/windows/dialog"
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/internal/stackescape"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
)

func (dlg *windowBase) CloseModal(result interface{}) {
	dlg.dialogResult = result
	dialog.EndDialog(dlg.Wrapper().Handle(), 0)
}

type dialogHandleManager struct {
	hwndManagerBase
	handle native.Handle
}

func (m *dialogHandleManager) Create(util.Bundle) native.Handle {
	return m.handle
}

func dlgMsgFilterProc(msg window.PMsg) bool {
	if receiver, ok := defaultObjectTable.Query(msg.Hwnd()).(Windows_WindowMessageReceiver); ok {
		return receiver.Windows_PreTranslateMessage(msg)
	}
	return false
}

func dialogProc(handle native.Handle, msg uint, wParam, lParam uintptr) bool {
	switch msg {
	case dialog.WM_INITDIALOG:
		dlg := stackescape.Get(stackescape.Id(lParam)).(*windowBase)
		title := dlg.Title()
		content := dlg.Content()
		menu := dlg.Menu()
		dlg.SetMenu(nil)
		frame := dlg.Frame()
		// To release the old content.
		dlg.SetContent(nil)
		dlg.Wrapper().SetHandleManager(&dialogHandleManager{handle: handle})
		util.Recreate(dlg, nil)
		dlg.SetTitle(title)
		dlg.SetFrame(frame)
		dlg.SetMenu(menu)
		dlg.SetContent(content)
		return true
	case window.WM_COMMAND:
		if lParam != 0 { // Dialog box generates control WM_COMMAND to dialog itself, not the container.
			if _, processed := handleWmCommandForControl(handle, msg, wParam, lParam); processed {
				return true
			}
		} else {
			switch window.LOWORD(uint(wParam)) {
			// Closing by user is converted to IDCANCEL.
			case dialog.IDOK, dialog.IDCANCEL:
				// dlg := defaultObjectTable.Query(handle).(Window)
				// if dlg.ShouldCloseEvent().HasHandler() && !dlg.ShouldCloseEvent().Send(&event{sender:dlg}) {
				// 	return true
				// }
				dialog.EndDialog(handle, 0)
				return true
			}
		}
	}
	return false
}

// The WH_MSGFILTER hook set to TranslateAccelerator().
var msgFilterHook native.Handle
var activeDialogs []Window

func createDialogDefer() {
	activeDialogs = activeDialogs[0 : len(activeDialogs)-1]
	if len(activeDialogs) == 0 {
		window.UnhookWindowsHookEx(msgFilterHook)
		msgFilterHook = 0
	}
}

func (w *windowBase) ShowModal(parent Window) interface{} {
	if len(activeDialogs) == 0 {
		// 1. https://support.microsoft.com/en-us/kb/100770
		// 2. https://msdn.microsoft.com/en-us/library/windows/desktop/ms644990(v=vs.85).aspx
		// WH_MSGFILTER
		// "Installs a hook procedure that monitors messages generated as a result of an input event in a dialog box, message box, menu, or scroll bar."
		msgFilterHook = dialog.SetupMsgFilterHookToTranslateAccelerator()
	}
	activeDialogs = append(activeDialogs, w.Self().(Window))
	defer createDialogDefer()
	dataId := stackescape.Add(w)
	defer stackescape.Remove(dataId)
	style := uint(window.GetWindowLongPtr(w.Wrapper().Handle(), window.GWL_STYLE))
	exStyle := uint(window.GetWindowLongPtr(w.Wrapper().Handle(), window.GWL_EXSTYLE))
	dialog.CreateDialog(0, parent.Wrapper().Handle(), style|dialog.DS_NOIDLEMSG, exStyle, uintptr(dataId))
	return w.dialogResult
}
