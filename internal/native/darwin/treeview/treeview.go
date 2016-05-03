package treeview

//#include "treeview.h"
import "C"

import (
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/internal/native/darwin/object"
	"github.com/kevin-yuan/rw/util/ustr"
)


func NewRWTreeViewItem(title string) native.Handle {
	return native.Handle(C.RWTreeViewItem_initWithTitle(C.RWTreeViewItem_alloc(), (*C.char)(ustr.CStringUtf8(title))))
}

func RWTreeView_title(item native.Handle) string {
	return C.GoString(C.RWTreeViewItem_title(C.OBJC_PTR(item)))
}

func RWTreeView_parent(item native.Handle) native.Handle {
	return native.Handle(C.RWTreeViewItem_parent(C.OBJC_PTR(item)))
}

func RWTreeViewItem_insertChild_atIndex(item native.Handle, child native.Handle, index uint) {
	C.RWTreeViewItem_insertChild_atIndex(C.OBJC_PTR(item), C.OBJC_PTR(child), C.ulong(index));
}

func RWTreeViewItem_removeChildAtIndex(item native.Handle, index uint) {
	C.RWTreeViewItem_removeChildAtIndex(C.OBJC_PTR(item), C.ulong(index))
}

// Returns array.NotFound if no such child.
func RWTreeViewItem_indexOfChild(item native.Handle, child native.Handle) uint {
	return uint(C.RWTreeViewItem_indexOfChild(C.OBJC_PTR(item), C.OBJC_PTR(child)))
}

func RWTreeViewItem_numberOfChildren(item native.Handle) uint {
	return uint(C.RWTreeViewItem_numberOfChildren(C.OBJC_PTR(item)))
}

func RWTreeViewItem_childAtIndex(item native.Handle, index uint) native.Handle {
	return native.Handle(C.RWTreeViewItem_childAtIndex(C.OBJC_PTR(item), C.ulong(index)))
}

func NewRWTreeViewDataSource() native.Handle {
	return native.Handle(object.NSObject_init(native.Handle(C.RWTreeViewDataSource_alloc())))
}

func RWTreeViewDataSource_insertItem_atIndex(ds native.Handle, item native.Handle, index uint) {
	C.RWTreeViewDataSource_insertItem_atIndex(C.OBJC_PTR(ds), C.OBJC_PTR(item), C.ulong(index))
}

func RWTreeViewDataSource_removeItemAtIndex(ds native.Handle, index uint) {
	C.RWTreeViewDataSource_removeItemAtIndex(C.OBJC_PTR(ds), C.ulong(index))
}

// Returns array.NotFound if no such item.
func RWTreeViewDataSource_indexOfItem(ds native.Handle, item native.Handle) uint {
	return uint(C.RWTreeViewDataSource_indexOfItem(C.OBJC_PTR(ds), C.OBJC_PTR(item)))
}

func RWTreeViewDataSource_numberOfItems(ds native.Handle) uint {
	return uint(C.RWTreeViewDataSource_numberOfItems(C.OBJC_PTR(ds)))
}

func RWTreeViewDataSource_itemAtIndex(ds native.Handle, index uint) native.Handle {
	return native.Handle(C.RWTreeViewDataSource_itemAtIndex(C.OBJC_PTR(ds), C.ulong(index)))
}

func NewRWTreeView(x, y, w, h int) native.Handle {
	return native.Handle(C.RWTreeView_initWithFrame(C.RWTreeView_alloc(), C.CGRectMake(C.CGFloat(x), C.CGFloat(y), C.CGFloat(w), C.CGFloat(h))))
}

func RWTreeView_outlineView(v native.Handle) native.Handle {
	return native.Handle(C.RWTreeView_outlineView(C.OBJC_PTR(v)))
}

func RWTreeView_isEditable(v native.Handle) bool {
	return bool(C.RWTreeView_isEditable(C.OBJC_PTR(v)))
}

func RWTreeView_setEditable(v native.Handle, editable bool) {
	C.RWTreeView_setEditable(C.OBJC_PTR(v), C.bool(editable))
}

func RWTreeView_sizeColumnToFit(v native.Handle) {
	C.RWTreeView_sizeColumnToFit(C.OBJC_PTR(v))
}

