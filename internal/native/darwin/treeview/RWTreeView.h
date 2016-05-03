#import <Cocoa/Cocoa.h>

// TreeViewItem
@interface RWTreeViewItem : NSObject {
    NSMutableArray* _children;
}

@property (copy,nonatomic) NSString* title;
@property (assign,nonatomic) RWTreeViewItem* parent;

- (instancetype)initWithTitle:(NSString*) title;
- (void)insertChild:(RWTreeViewItem*)child atIndex:(NSUInteger)index;
- (void)removeChildAtIndex:(NSUInteger)index;
- (NSUInteger)indexOfChild:(RWTreeViewItem*)child;
- (NSUInteger)numberOfChildren;
- (RWTreeViewItem*)childAtIndex:(NSUInteger)index;

@end

///////////////////////////////////////////////////////////////////////////////////////

// TreeViewDataSource
@interface RWTreeViewDataSource : NSObject<NSOutlineViewDataSource> {
    NSMutableArray* _items;
}

@property(getter=isEditable) BOOL editable;

- (instancetype)init;
- (void)insertItem:(RWTreeViewItem*)item atIndex:(NSUInteger)index ;
- (void)removeItemAtIndex:(NSUInteger)index ;
- (NSUInteger)indexOfItem:(RWTreeViewItem*)item;
- (RWTreeViewItem*)itemAtIndex:(NSUInteger)index;
- (NSUInteger)numberOfItems ;
- (RWTreeViewItem*)itemAtIndex:(NSUInteger)index;

@end

///////////////////////////////////////////////////////////////////////////////////////

// TreeView
@interface RWTreeView : NSScrollView

- (instancetype)initWithFrame:(NSRect)frameRect;
// outlineView returns the NSOutlineView instance of the tree view.
- (NSOutlineView*)outlineView;
- (BOOL)isEditable;
- (void)setEditable:(BOOL)editable;
// Resizes the outline column to fit it's content width.
- (void)sizeColumnToFit;

@end