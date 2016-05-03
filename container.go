package rw

import (
	"fmt"
	"github.com/kevin-yuan/rw/event"
	"image/color"
)

// Container is the base interface for all Controls that holde Conrols.
// If a Control is added in a Container and that Container is released, the container automatically releases the Control, frees
// resources and memory occupied by the Control. Releasing a Container will remove itself from it's Parent, frees itself and releases all it's Children.
type Container interface {
	Control
	// Children returns the controls in this container.
	Children() []Control
	// ChildrenCount returns the count of controls in this container.
	ChildrenCount() int
	// Child returns the ith child in this container.
	Child(i int) Control
	// Add adds a new Control to this container. Do nothing if control is already in this container.
	// The control will be removed form it's parent, if any, before added to this container.
	Add(control Control)
	// Remove removes a child from this container. Do nothing if control is not a child of this container.
	Remove(control Control)
	// BackgroundColor returns the background color of the container.
	BackgroundColor() color.Color
	// SetBackgroundColor sets the background color of the container.
	SetBackgroundColor(backgroundColor color.Color)
	// SizeChanged is an event hub where an event is sent after the size of the Container is changed.
	SizeChanged() *event.Hub

	containerExtra
}

func (c *containerBase) String() string {
	if c.Wrapper().Valid() {
		return fmt.Sprintf("Container %#X", c.Wrapper().Handle())
	} else {
		return "Container <Invalid>"
	}
}

func (c *containerBase) SizeChanged() *event.Hub {
	return &c.sizeChanged
}

func (c *containerBase) Children() []Control {
	if len(c.children) == 0 {
		return nil
	}
	return append(([]Control(nil)), c.children...)
}

func (c *containerBase) ChildrenCount() int {
	return len(c.children)
}

func (c *containerBase) Child(i int) Control {
	return c.children[i]
}

func (c *containerBase) Add(control Control) {
	var self = c.Self().(Container)
	if self == control {
		panic("Can't add to self")
	}

	for _, child := range c.children {
		if child == control {
			return // Aleady in this container.
		}
	}
	// Remove from it's parent, if any.
	if prevParent := control.Parent(); prevParent != nil {
		prevParent.Remove(control)
	}
	// Add to this container.
	c.children = append(c.children, control)
	nativeSetParent(control.Wrapper().Handle(), c.Wrapper().Handle())
	// Set parent of control
	control.setParent(self)
	applyTabOrder(self, control)
}

func (c *containerBase) Remove(childToRemove Control) {
	for i, child := range c.children {
		if child == childToRemove {
			c.children = append(c.children[:i], c.children[i+1:]...)
			nativeRemoveFromParent(childToRemove.Wrapper().Handle())
			childToRemove.setParent(nil)
			// Layout
			break
		}
	}
}

// NewContainerTemplate creates a template of Container.
// Use corresponding NewXxx functions(NewPanel in package rw/panel for example) to create objects of concrete Container.
func NewContainerTemplate() Container {
	return newContainerTemplate()
}
