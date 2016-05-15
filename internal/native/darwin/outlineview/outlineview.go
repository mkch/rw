package outlineview

import (
	"github.com/mkch/rw/native"
)

//#include "outlineview.h"
import "C"

func NSOutlineView_reloadItem_reloadChildren(tv native.Handle, item native.Handle, reloadChildren bool) {
	C.NSOutlineView_reloadItem_reloadChildren(C.OBJC_PTR(tv), C.OBJC_PTR(item), C.bool(reloadChildren))
}

func NSOutlineView_isItemExpanded(tv native.Handle, item native.Handle) bool {
	return bool(C.NSOutlineView_isItemExpanded(C.OBJC_PTR(tv), C.OBJC_PTR(item)))
}

func NSOutlineView_expandItem_expandChildren(tv, item native.Handle, expandChildren bool) {
	C.NSOutlineView_expandItem_expandChildren(C.OBJC_PTR(tv), C.OBJC_PTR(item), C.bool(expandChildren))
}

func NSOutlineView_collapseItem_collapseChildren(tv, item native.Handle, collapseChildern bool) {
	C.NSOutlineView_expandItem_expandChildren(C.OBJC_PTR(tv), C.OBJC_PTR(item), C.bool(collapseChildern))
}
