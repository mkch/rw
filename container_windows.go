package rw

import (
	"github.com/mkch/rw/event"
	colorconv "github.com/mkch/rw/internal/native/windows/color"
	"github.com/mkch/rw/internal/native/windows/gdi"
	"github.com/mkch/rw/internal/native/windows/window"
	"github.com/mkch/rw/internal/native/windows/window/winutil"

	"github.com/mkch/rw/native"
	"image/color"
	"sort"
)

type containerExtra interface {
	setWindow(Window)
	window() Window
	applyChildrenTabOrder()
}

// containerBase is native container wrapper.
type containerBase struct {
	controlBase
	children                       []Control
	sizeChanged                    event.Hub
	backgroundColor                color.Color
	backgroundBrush                uintptr
	win                            Window // If this container is the content of a window
	applyChildrenTabOrderScheduled bool
}

func (c *containerBase) setWindow(win Window) {
	c.win = win
}

func (c *containerBase) window() Window {
	return c.win
}

func (c *containerBase) doApplyChildrenTabOrder() {
	if len(c.children) <= 1 {
		return // Zero or only one child, already in the right z-order.
	}
	// Use c.Children(), not c.children, to make a copy of the children.
	children := c.Children()
	// After stable ascending sort, the control with smallest tab order is in foront.
	// The order of controls with equal tab order is kept.
	sort.Stable(controlsByTabOrderAsc(children))
	for _, child := range children {
		// Append one by one to the bottom(aka. end) of z-order.
		window.SetWindowPos(child.Wrapper().Handle(), window.HWND_BOTTOM, 0, 0, 0, 0, window.SWP_NOOWNERZORDER|window.SWP_NOSIZE|window.SWP_NOMOVE)
	}
}

func applyTabOrder(parent Container, child Control) {
	parent.applyChildrenTabOrder()
}

func (c *containerBase) applyChildrenTabOrder() {
	if c.applyChildrenTabOrderScheduled {
		return
	}
	c.applyChildrenTabOrderScheduled = true
	unsafePost(func() { c.doApplyChildrenTabOrder(); c.applyChildrenTabOrderScheduled = false })
}

func (c *containerBase) TabStop() bool {
	return c.controlBase.TabStop() &&
		!winutil.HasExStyle(c.Wrapper().Handle(), window.WS_EX_CONTROLPARENT)
}

func (c *containerBase) SetTabStop(s bool) {
	c.controlBase.SetTabStop(s)
	if s {
		winutil.ModifyStyles(c.Wrapper().Handle(), 0, 0, 0, window.WS_EX_CONTROLPARENT, false)
	} else {
		winutil.ModifyStyles(c.Wrapper().Handle(), 0, 0, window.WS_EX_CONTROLPARENT, 0, false)
	}
}

func (c *containerBase) Release() {
	if c.win != nil {
		c.win.SetContent(nil)
	}
	c.controlBase.Release()
}

func (c *containerBase) BackgroundColor() color.Color {
	if c.backgroundBrush == 0 {
		return colorconv.Color(window.GetSysColor(window.COLOR_WINDOW))
	}
	return c.backgroundColor
}

func (c *containerBase) SetBackgroundColor(backgroundColor color.Color) {
	if c.backgroundBrush != 0 {
		gdi.DeleteObject(c.backgroundBrush)
	}
	c.backgroundColor = backgroundColor
	c.backgroundBrush = gdi.CreateSolidBrush(colorconv.Uint32(c.backgroundColor))
	window.InvalidateRectNull(c.Wrapper().Handle(), true)
}

func (c *containerBase) Windows_WndProc(handle native.Handle, msg uint, wParam, lParam uintptr) uintptr {
	switch msg {
	case window.WM_DESTROY:
		if c.backgroundBrush != 0 {
			gdi.DeleteObject(c.backgroundBrush)
		}
	case window.WM_SIZE:
		if c.sizeChanged.HasHandler() {
			c.sizeChanged.Send(&simpleEvent{sender: c.Self()})
		}
	case window.WM_ERASEBKGND:
		if c.backgroundBrush != 0 {
			dc := wParam
			oldBrush := gdi.SelectObject(dc, c.backgroundBrush)
			oldPen := gdi.SelectObject(dc, gdi.GetStockObject(gdi.NULL_PEN))
			frame := c.Frame()
			gdi.Rectangle(dc, 0, 0, frame.Width+1, frame.Height+1)
			gdi.SelectObject(dc, oldPen)
			gdi.SelectObject(dc, oldBrush)
			return 1
		}
	case window.WM_COMMAND:
		// https://msdn.microsoft.com/en-us/library/windows/desktop/ms647591(v=vs.85).aspx
		if lParam != 0 {
			if result, processed := handleWmCommandForControl(handle, msg, wParam, lParam); processed {
				return result
			}
		}
	}
	return c.controlBase.Windows_WndProc(handle, msg, wParam, lParam)
}

func handleWmCommandForControl(handle native.Handle, msg uint, wParam, lParam uintptr) (result uintptr, processed bool) {
	if ctrl, ok := defaultObjectTable.Query(native.Handle(lParam)).(Control); ok {
		return ctrl.Windows_ReflectedWndProc(handle, msg, wParam, lParam)
	}
	return 0, false
}

func nativeSetParent(child, parent native.Handle) {
	window.SetParent(child, parent)
}

func nativeRemoveFromParent(child native.Handle) {
	window.SetParent(child, winutil.DockerWindow())
}

func newContainerTemplate() Container {
	c := &containerBase{}
	initControlBase(&c.controlBase)
	return c
}
