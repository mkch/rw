#ifndef _RW_TABLEVIEW_H
#define _RW_TABLEVIEW_H

#include "../types.h"

void NSOutlineView_reloadItem_reloadChildren(OBJC_PTR ptr, OBJC_PTR item, bool reloadChildren);
bool NSOutlineView_isItemExpanded(OBJC_PTR ptr, OBJC_PTR item);
void NSOutlineView_expandItem_expandChildren(OBJC_PTR ptr, OBJC_PTR item, bool expandChildren);
void NSOutlineView_collapseItem_collapseChildren(OBJC_PTR ptr, OBJC_PTR item, bool collapseChildren);

#endif
