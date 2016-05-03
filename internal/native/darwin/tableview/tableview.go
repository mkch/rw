package tableview

import (
	"github.com/kevin-yuan/rw/native"
)

//#include "tableview.h"
import "C"

func NSTableView_reloadData(tv native.Handle) {
	C.NSTableView_reloadData(C.OBJC_PTR(tv))
}

func NsTableView_dataSource(tv native.Handle) native.Handle {
	return native.Handle(C.NSTableView_dataSource(C.OBJC_PTR(tv)))
}