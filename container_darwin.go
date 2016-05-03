package rw

import (
	"github.com/kevin-yuan/rw/event"
	"github.com/kevin-yuan/rw/internal/native/darwin/object"
	"github.com/kevin-yuan/rw/internal/native/darwin/view"

	"github.com/kevin-yuan/rw/native"
	"image/color"
	"sort"
)

type containerExtra interface{}

// winContainerImpl is native container wrapper.
type containerBase struct {
	controlBase
	children                       []Control
	sizeChanged                    event.Hub
	applyChildrenTabOrderScheduled bool
}

func (c *containerBase) BackgroundColor() color.Color {
	return view.RWFlippedView_backgroundColor(c.Wrapper().Handle())
}

func (c *containerBase) SetBackgroundColor(backgroundColor color.Color) {
	handle := c.Wrapper().Handle()
	view.RWFlippedView_setBackgroundColor(handle, backgroundColor)
	view.NSView_display(handle)
}

func nativeSetParent(child, parent native.Handle) {
	// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSView_Class/#//apple_ref/occ/instm/NSView/addSubview:
	// - addSubview:
	// "The view retains aView."
	view.NSView_addSubview(parent, object.NSObject_autorelease(child))
}

func nativeRemoveFromParent(child native.Handle) {
	// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSView_Class/#//apple_ref/occ/instm/NSView/removeFromSuperview
	// - removeFromSuperview
	// "The view is also released; if you plan to reuse it, be sure to retain it before sending this message and to release it as appropriate when adding it as a subview of another NSView."
	object.NSObject_retain(child)
	view.NSView_removeFromSuperview(child)
}

func applyTabOrder(parent Container, child Control) {
	if win := view.NSView_window(child.Wrapper().Handle()); win != 0 {
		defaultObjectTable.Query(win).(Window).recalculateTabOrder()
	}
}

// recalculateTabOrder apply Key-view chain based on TabOrder.
// c is the container which and whose children(recursively) are in count.
// _tails is the tails in the key-view chain so far.
// _head is a head control in the key-view chain so far.
// head is head control in the key-view chain.
// tails is the tails control in the key-view chain. The caller of this function can set the nextKeyView of tails to head
// to make the chain close(a loop).
func recalculateTabOrder(c Container, _head native.Handle, _tails []Control) (head native.Handle, tails []Control) {
	var children []Control
	if c.ChildrenCount() > 0 {
		children = c.Children()
		sort.Stable(controlsByTabOrderAsc(children))
	}
	containerHandle := c.Wrapper().Handle()
	containerTabStop := c.TabStop()
	if containerTabStop {
		// The container has tab stop. Link _tails(if any) to this container.
		for _, tail := range _tails {
			view.NSView_setNextKeyView(tail.Wrapper().Handle(), containerHandle)
		}
		// All the _tails are linked to this container itself, son othing should be linked to it's children.
		_tails = nil
		if _head == 0 {
			_head = containerHandle
		}
	}
	for _, child := range children {
		childHandle := child.Wrapper().Handle()
		if childContainer, ok := child.(Container); ok {
			// Recurse.
			_head, _tails = recalculateTabOrder(childContainer, _head, _tails)
		} else if child.TabStop() {
			// This child has tab stop. Link _tails(if any) to this control.
			for _, tail := range _tails {
				view.NSView_setNextKeyView(tail.Wrapper().Handle(), childHandle)
			}
			// This control should be linked to the next control.
			_tails = []Control{child}
			if _head == 0 {
				_head = childHandle
			}
		} else {
			// This child does not has tab stop. Nothing is linked to this control, but this control and
			// all the _tails should be linked to the next control.
			_tails = append(_tails, child)
		}
	}
	head = _head
	// This container and all the _tails one should be linked to the next control.
	tails = append(_tails, c)
	return
}

func newContainerTemplate() Container {
	return &containerBase{}
}
