#ifndef _RW_TREEVIEW_H
#define _RW_TREEVIEW_H

#include "../types.h"
#include <CoreGraphics/CGGeometry.h>

////////////////////////////////////////////////////////////////////////////////////////
// RWTreeViewItem
OBJC_PTR RWTreeViewItem_alloc();
OBJC_PTR RWTreeViewItem_initWithTitle(OBJC_PTR ptr, char* title);
char* RWTreeViewItem_title(OBJC_PTR ptr);
OBJC_PTR RWTreeViewItem_parent(OBJC_PTR ptr);
void RWTreeViewItem_insertChild_atIndex(OBJC_PTR ptr, OBJC_PTR item, unsigned long index);
void RWTreeViewItem_removeChildAtIndex(OBJC_PTR ptr, unsigned long index);
unsigned long RWTreeViewItem_indexOfChild(OBJC_PTR ptr, OBJC_PTR child);
unsigned long RWTreeViewItem_numberOfChildren(OBJC_PTR ptr);
OBJC_PTR RWTreeViewItem_childAtIndex(OBJC_PTR ptr, unsigned long index);
////////////////////////////////////////////////////////////////////////////////////////
// RWTreeViewDataSource
OBJC_PTR RWTreeViewDataSource_alloc();
void RWTreeViewDataSource_insertItem_atIndex(OBJC_PTR ptr, OBJC_PTR item, unsigned long index);
void RWTreeViewDataSource_removeItemAtIndex(OBJC_PTR ptr, unsigned long index);
unsigned long RWTreeViewDataSource_indexOfItem(OBJC_PTR ptr, OBJC_PTR item);
unsigned long RWTreeViewDataSource_numberOfItems(OBJC_PTR ptr);
OBJC_PTR RWTreeViewDataSource_itemAtIndex(OBJC_PTR ptr, unsigned long index);
////////////////////////////////////////////////////////////////////////////////////////
// RWTreeView
OBJC_PTR RWTreeView_alloc();
OBJC_PTR RWTreeView_initWithFrame(OBJC_PTR ptr, CGRect frame);
OBJC_PTR RWTreeView_outlineView(OBJC_PTR ptr);
bool RWTreeView_isEditable(OBJC_PTR ptr);
void RWTreeView_setEditable(OBJC_PTR ptr, bool editable);
void RWTreeView_sizeColumnToFit(OBJC_PTR ptr);

#endif
