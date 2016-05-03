#include "treeview.h"
#include "RWTreeView.h"

////////////////////////////////////////////////////////////////////////////////////////
// RWTreeViewItem

OBJC_PTR RWTreeViewItem_alloc() {
	return [RWTreeViewItem alloc];
}

OBJC_PTR RWTreeViewItem_initWithTitle(OBJC_PTR ptr, char* title) {
   return [(RWTreeViewItem*)ptr initWithTitle:[NSString stringWithUTF8String:title]];
}

char* RWTreeViewItem_title(OBJC_PTR ptr) {
    return (char*)[[(RWTreeViewItem*)ptr title] UTF8String];
}

OBJC_PTR RWTreeViewItem_parent(OBJC_PTR ptr) {
    return [(RWTreeViewItem*)ptr parent];
}

void RWTreeViewItem_insertChild_atIndex(OBJC_PTR ptr, OBJC_PTR item, unsigned long index) {
    [(RWTreeViewItem*)ptr insertChild:(RWTreeViewItem*)item atIndex:index];
}

void RWTreeViewItem_removeChildAtIndex(OBJC_PTR ptr, unsigned long index) {
    [(RWTreeViewItem*)ptr removeChildAtIndex:index];
}

unsigned long RWTreeViewItem_indexOfChild(OBJC_PTR ptr, OBJC_PTR child) {
    return [(RWTreeViewItem*)ptr indexOfChild:(RWTreeViewItem*)child];
}

unsigned long RWTreeViewItem_numberOfChildren(OBJC_PTR ptr) {
    return [(RWTreeViewItem*)ptr numberOfChildren];
}


OBJC_PTR RWTreeViewItem_childAtIndex(OBJC_PTR ptr, unsigned long index) {
    return [(RWTreeViewItem*)ptr childAtIndex:index];
}

////////////////////////////////////////////////////////////////////////////////////////
// RWTreeViewDataSource
OBJC_PTR RWTreeViewDataSource_alloc() {
	return [RWTreeViewDataSource alloc];
}

void RWTreeViewDataSource_insertItem_atIndex(OBJC_PTR ptr, OBJC_PTR item, unsigned long index) {
    [(RWTreeViewDataSource*)ptr insertItem:(RWTreeViewItem*)item atIndex:index];
}

void RWTreeViewDataSource_removeItemAtIndex(OBJC_PTR ptr, unsigned long index) {
    [(RWTreeViewDataSource*)ptr removeItemAtIndex:index];
}

unsigned long RWTreeViewDataSource_indexOfItem(OBJC_PTR ptr, OBJC_PTR item) {
    return [(RWTreeViewDataSource*)ptr indexOfItem:(RWTreeViewItem*)item];
}

unsigned long RWTreeViewDataSource_numberOfItems(OBJC_PTR ptr) {
    return [(RWTreeViewDataSource*)ptr numberOfItems];
}

OBJC_PTR RWTreeViewDataSource_itemAtIndex(OBJC_PTR ptr, unsigned long index) {
    return [(RWTreeViewDataSource*)ptr itemAtIndex:index];
}


////////////////////////////////////////////////////////////////////////////////////////
// RWTreeView

OBJC_PTR RWTreeView_alloc() {
	return [RWTreeView alloc];
}

OBJC_PTR RWTreeView_initWithFrame(OBJC_PTR ptr, CGRect frame) {
    return [(RWTreeView*)ptr initWithFrame:frame];
}

OBJC_PTR RWTreeView_outlineView(OBJC_PTR ptr) {
    return [(RWTreeView*)ptr outlineView];
}

bool RWTreeView_isEditable(OBJC_PTR ptr) {
    return [(RWTreeView*)ptr isEditable];
}

void RWTreeView_setEditable(OBJC_PTR ptr, bool editable) {
    [(RWTreeView*)ptr setEditable:editable];
}

void RWTreeView_sizeColumnToFit(OBJC_PTR ptr) {
    [(RWTreeView*)ptr sizeColumnToFit];
}

