#import <AppKit/NSTableView.h>
#import "tableview.h"

void NSTableView_reloadData(OBJC_PTR ptr) {
	[(NSTableView*)ptr reloadData];
}

OBJC_PTR NSTableView_dataSource(OBJC_PTR ptr) {
	return [(NSTableView*)ptr dataSource];
}

