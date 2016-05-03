#import <AppKit/NSOutlineView.h>
#import "outlineview.h"

void NSOutlineView_reloadItem_reloadChildren(OBJC_PTR ptr, OBJC_PTR item, bool reloadChildren) {
	[(NSOutlineView*)ptr reloadItem:(NSOutlineView*)item reloadChildren:reloadChildren];
}

bool NSOutlineView_isItemExpanded(OBJC_PTR ptr, OBJC_PTR item) {
	return [(NSOutlineView*)ptr isItemExpanded:(id)item];
}

void NSOutlineView_expandItem_expandChildren(OBJC_PTR ptr, OBJC_PTR item, bool expandChildren) {
	[(NSOutlineView*)ptr expandItem:(id)item expandChildren:expandChildren];
}

void NSOutlineView_collapseItem_collapseChildren(OBJC_PTR ptr, OBJC_PTR item, bool collapseChildren) {
	[(NSOutlineView*)ptr collapseItem:(id)item collapseChildren:collapseChildren];
}

