package rw

import (
	"github.com/mkch/rw/internal/native/darwin/view"
	"github.com/mkch/rw/internal/native/darwin/window"

	"github.com/mkch/rw/util"
)

// ControlPlatformSpecific contains extra platform specific methods.
type ControlPlatformSpecific interface{}
type controlExtra interface{}

// controlBase can be used to construct controls.
type controlBase struct {
	objectBase
	wrapper  util.WrapperImpl
	parent   Container
	tabStop  bool
	tabOrder uint
}

func (c *controlBase) Release() {
	util.Release(c)
}

func (c *controlBase) Visible() bool {
	return !view.NSView_isHidden(c.Wrapper().Handle())
}

func (c *controlBase) SetVisible(v bool) {
	view.NSView_setHidden(c.Wrapper().Handle(), !v)
}

func (c *controlBase) Focus() {
	handle := c.Wrapper().Handle()
	if win := view.NSView_window(handle); win != 0 {
		window.NSWindow_makeFirstResponder(win, handle)
	}
}

func (c *controlBase) Focused() bool {
	handle := c.Wrapper().Handle()
	if win := view.NSView_window(handle); win != 0 {
		return window.NSWindow_firstResponder(win) == handle
	}
	return false
}

func (c *controlBase) TabStop() bool {
	return c.tabStop
}

func (c *controlBase) SetTabStop(stop bool) {
	if stop == c.tabStop {
		return
	}
	c.tabStop = stop
	if win := view.NSView_window(c.Wrapper().Handle()); win != 0 {
		defaultObjectTable.Query(win).(Window).recalculateTabOrder()
	}
}

func (c *controlBase) TabOrder() uint {
	return c.tabOrder
}

func (c *controlBase) SetTabOrder(order uint) {
	if order == c.tabOrder {
		return
	}
	c.tabOrder = order
	if win := view.NSView_window(c.Wrapper().Handle()); win != 0 {
		defaultObjectTable.Query(win).(Window).recalculateTabOrder()
	}
}

func (c *controlBase) SetFrame(frame Rect) {
	handle := c.Wrapper().Handle()
	view.NSView_setFrameSize(handle, frame.Width, frame.Height)
	view.NSView_setFrameOrigin(handle, frame.X, frame.Y)
}

func (c *controlBase) Frame() Rect {
	x, y, w, h := view.NSView_frame(c.Wrapper().Handle())
	return Rect{X: x, Y: y, Width: w, Height: h}
}

func newControlTemplate() Control {
	return &controlBase{tabStop: true}
}

// ControlHandleManagerBase is the building block of HandleManager of concrete Control types.
type ControlHandleManagerBase struct {objcHandleManagerBase}

