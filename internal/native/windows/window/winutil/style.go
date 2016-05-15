package winutil

import (
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/native"
)

func HasStyle(handle native.Handle, styleToTest uint) bool {
	return uint(window.GetWindowLongPtr(handle, window.GWL_STYLE))&styleToTest != 0
}

func HasExStyle(handle native.Handle, styleToTest uint) bool {
	return uint(window.GetWindowLongPtr(handle, window.GWL_EXSTYLE))&styleToTest != 0
}

func ModifyStyles(handle native.Handle, styleToAdd, styleToRemove, exStyleToAdd, exStyleToRemove uint, frameChanged bool) {
	if styleToAdd != 0 || styleToRemove != 0 {
		style := uint(window.GetWindowLongPtr(handle, window.GWL_STYLE))
		style |= styleToAdd
		style &^= styleToRemove
		window.SetWindowLongPtr(handle, window.GWL_STYLE, uintptr(style))
	}
	if exStyleToAdd != 0 || exStyleToRemove != 0 {
		exStyle := uint(window.GetWindowLongPtr(handle, window.GWL_EXSTYLE))
		exStyle |= exStyleToAdd
		exStyle &^= exStyleToRemove
		window.SetWindowLongPtr(handle, window.GWL_EXSTYLE, uintptr(exStyle))
	}
	if frameChanged {
		window.SetWindowPos(handle, 0, 0, 0, 0, 0, window.SWP_NOACTIVATE|window.SWP_NOOWNERZORDER|window.SWP_NOSIZE|window.SWP_NOMOVE|window.SWP_FRAMECHANGED)
	}
}
