package rw

import (
	"github.com/mkch/rw/event"
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/internal/native/windows/window/winutil"

	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util"
)

// ControlPlatformSpecific contains extra platform specific methods.
type ControlPlatformSpecific interface {
	Windows_WindowMessageReceiver
	// Windows_ReflectedWndProc receives all control messages routed by it's parent window.
	// Only available on Windows platform.
	Windows_ReflectedWndProc(handle native.Handle, msg uint, wParam, lParam uintptr) (result uintptr, processed bool)
}

type controlExtra interface {
}

// controlBase can be used to construct controls.
type controlBase struct {
	objectBase
	wrapper     util.WrapperImpl
	parent      Container
	prevWndProc uintptr
	tabOrder    uint
}

func (c *controlBase) Windows_ReflectedWndProc(handle native.Handle, msg uint, wParam, lParam uintptr) (result uintptr, processed bool) {
	return 0, false
}

func (c *controlBase) Windows_PreTranslateMessage(msg window.PMsg) bool {
	if defaultObjectTable.Query(window.GetAncestor(msg.Hwnd(), window.GA_ROOT)).(Window).Windows_PreTranslateMessage(msg) {
		return true
	}
	return false
}

func (c *controlBase) Windows_WndProc(handle native.Handle, msg uint, wParam, lParam uintptr) uintptr {
	return window.CallWindowProc(c.prevWndProc, handle, msg, wParam, lParam)
}

func (c *controlBase) TabStop() bool {
	return winutil.HasStyle(c.Wrapper().Handle(), window.WS_TABSTOP)
}

func (c *controlBase) SetTabStop(s bool) {
	if s {
		winutil.ModifyStyles(c.Wrapper().Handle(), window.WS_TABSTOP, 0, 0, 0, false)
	} else {
		winutil.ModifyStyles(c.Wrapper().Handle(), 0, window.WS_TABSTOP, 0, 0, false)
	}
}

func (c *controlBase) Focus() {
	window.SetFocus(c.Wrapper().Handle())
}

func (c *controlBase) Focused() bool {
	return window.GetFocus() == c.Wrapper().Handle()
}

func (c *controlBase) TabOrder() uint {
	return c.tabOrder
}

func (c *controlBase) SetTabOrder(order uint) {
	if c.tabOrder == order {
		return
	}
	c.tabOrder = order
	if c.parent != nil {
		c.parent.applyChildrenTabOrder()
	}
}

func (c *controlBase) Visible() bool {
	return winutil.HasStyle(c.Wrapper().Handle(), window.WS_VISIBLE)
}

func (c *controlBase) SetVisible(v bool) {
	var cmd = window.SW_HIDE
	if v {
		cmd = window.SW_SHOWNA
	}
	window.ShowWindow(c.Wrapper().Handle(), cmd)
}

func (c *controlBase) SetFrame(frame Rect) {
	setFrame(c.Wrapper().Handle(), frame.X, frame.Y, frame.Width, frame.Height)
}

func (c *controlBase) Release() {
	if c.parent != nil {
		c.parent.Remove(c.Self().(Control))
	}
	util.Release(c)
}

func (c *controlBase) Frame() Rect {
	handle := c.Wrapper().Handle()
	x, y, width, height := window.GetWindowRect(handle)
	window.ScreenToClient(window.GetParent(handle), &x, &y)
	return Rect{x, y, width, height}
}

func (c *controlBase) afterRegistered(event event.Event, nextHook event.Handler) bool {
	c.prevWndProc = setWndProc(c.Wrapper().Handle())
	return nextHook(event)
}

func initControlBase(c *controlBase) *controlBase {
	c.Wrapper().AfterRegistered().AddHook(c.afterRegistered)
	return c
}

func newControlTemplate() Control {
	return initControlBase(&controlBase{})
}

// ControlHandleManagerBase is the building block of HandleManager of concrete Control types.
type ControlHandleManagerBase struct {
	hwndManagerBase
}
