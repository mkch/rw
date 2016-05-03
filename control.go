package rw

import (
	"fmt"
	"github.com/kevin-yuan/rw/util"
)

// Control is the base interface for all window widgets that can be added into a Container.
// If a Control is added in a Container and that Container is released, the Container automatically releases the Control, frees
// resources and memory occupied by the Control. An orphan Control must be released by the application code(calling Control.Release). Otherwise, resource/memory
// leak may occur. Releasing a control will remove itself from it's Container.
type Control interface {
	Object
	// Visible returns whether the window is visible.
	// A visible control is not visible on screen if its parent, its parent's parent, and so forth, is not visible.
	Visible() bool
	// SetVisible sets the visibility state of the control.
	SetVisible(v bool)
	// Parent returns the parent container of this control.
	Parent() Container
	// SetFrame sets the position and size of the window, in parent coordinates.
	SetFrame(frame Rect)
	// Frame returns the position and size of the window, in parent coordinates.
	Frame() Rect
	// TabsStop returns whether this control can be focused using Tab key on keyboard.
	TabStop() bool
	// SetTabStop sets whether this control can be focused using Tab key on keyboard.
	SetTabStop(bool)
	// Focus make this control the focused control. A focused control receives keyboard events.
	Focus()
	// Focused returns whether this control is currently the focused control.
	Focused() bool
	// TabOrder returns the order in which child controls are visited when the user presses the Tab key.
	// Smaller tab order means earlier visiting.
	// Tab order is meaningful only if the TabStop() is true and if the control has a parent.
	// Controls with the same tab order are ordered by the order they were added to the parent.
	TabOrder() uint
	// SetTabOrder sets the order in which child controls are visited when the user presses the Tab key.
	SetTabOrder(uint)

	util.WrapperHolder

	ControlPlatformSpecific

	setParent(Container)
	controlExtra
}

func (c *controlBase) Parent() Container {
	return c.parent
}

// setParent just sets the parent value.
func (c *controlBase) setParent(parent Container) {
	c.parent = parent
}

func (c *controlBase) Wrapper() util.Wrapper {
	return &c.wrapper
}

func (c *controlBase) String() string {
	if c.Wrapper().Valid() {
		return fmt.Sprintf("Control %#X", c.Wrapper().Handle())
	} else {
		return "Control <Invalid>"
	}
}

// controlsByTabOrderAsc is a sort.Interface, which can be used to sort Control by tab order ascendingly.
type controlsByTabOrderAsc []Control

func (a controlsByTabOrderAsc) Len() int {
	return len(a)
}

func (a controlsByTabOrderAsc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a controlsByTabOrderAsc) Less(i, j int) bool {
	return a[i].TabOrder() < a[j].TabOrder()
}

// NewControlTemplate creates a template of Control.
// Use corresponding NewXxx functions(NewButton in package rw/button for example) to create objects of concrete Control.
func NewControlTemplate() Control {
	return newControlTemplate()
}
