#import "RWTreeView.h"

@interface Dummy_class_to_provide_prototype_of__sizeToFitWidthOfColumn_method : NSObject
-(CGFloat)_sizeToFitWidthOfColumn:(NSInteger)col;
@end

@implementation Dummy_class_to_provide_prototype_of__sizeToFitWidthOfColumn_method

-(CGFloat)_sizeToFitWidthOfColumn:(NSInteger)col {
    return 0;
}

@end


@interface NSTableColumn(NSTableColumnSizeToFitContent)

- (void)sizeToFitContent;

@end

@implementation NSTableColumn(NSTableColumnSizeToFitContent)

- (void)sizeToFitContent {
    NSTableView* tableView = [self tableView];
    NSInteger index = (NSInteger)[[tableView tableColumns] indexOfObject:self];
    CGFloat w = [(id)tableView _sizeToFitWidthOfColumn:index];
    [self setWidth:w];
}

@end

///////////////////////////////////////////////////////////////////////////////////////
// TreeViewItem

@implementation RWTreeViewItem

- (instancetype)initWithTitle:(NSString*) title {
    self = [super init];
    if(self != nil) {
        self.title = title;
    }
    return self;
}

- (void)dealloc {
    [_children release];
    [super dealloc];
}


- (void)insertChild:(RWTreeViewItem*)child atIndex:(NSUInteger)index {
    if([child parent] != nil) {
        [NSException raise:@"BadParentException"
                    format:@"RWTreeViewItem \"%@\" is already in parent \"%@\" before being inserted into \"%@\".",
                    child.title, child.parent.title, self.title];
    }
    if(_children == nil) {
        _children = [NSMutableArray new];
    }
    [_children insertObject:child
                    atIndex:index];
    [child setParent:self];
}

- (void)removeChildAtIndex:(NSUInteger)index {
    if(_children == nil) {
        return;
    }
    [[_children objectAtIndex:index] setParent:nil];
    [_children removeObjectAtIndex:index];
}

- (NSUInteger)indexOfChild:(RWTreeViewItem*)child {
    if(_children == nil) {
        return NSNotFound;
    }
    return [_children indexOfObject:child];
}

- (NSUInteger)numberOfChildren {
    if(_children == nil) {
        return 0;
    }
    return [_children count];
}

- (RWTreeViewItem*)childAtIndex:(NSUInteger)index {
    if(_children == nil) {
        return nil;
    }
    return [_children objectAtIndex:index];
}

@end


///////////////////////////////////////////////////////////////////////////////////////
// TreeViewDataSource

@implementation RWTreeViewDataSource

- (instancetype)init {
    self = [super init];
    if(self) {
        _items = [NSMutableArray new];
    }
    return self;
}

- (void)dealloc {
    [_items release];
    [super dealloc];
}

- (void)insertItem:(RWTreeViewItem*)item atIndex:(NSUInteger)index {
    [_items insertObject:item
                    atIndex:index];
}

- (void)removeItemAtIndex:(NSUInteger)index {
    [_items removeObjectAtIndex:index];
}

- (NSUInteger)indexOfItem:(RWTreeViewItem*)item {
    return [_items indexOfObject:item];
}

- (NSUInteger)numberOfItems {
    return [_items count];
}

- (RWTreeViewItem*)itemAtIndex:(NSUInteger)index {
    return [_items objectAtIndex:index];
}

- (id)outlineView:(NSOutlineView *)outlineView child:(NSInteger)index ofItem:(id)item {
    if(item == nil) { // Root item.
        return [_items objectAtIndex:index];
    } else {
        return [(RWTreeViewItem*)item childAtIndex:index];
    }
}

- (BOOL)outlineView:(NSOutlineView *)outlineView isItemExpandable:(id)item {
    return [(RWTreeViewItem*)item numberOfChildren] > 0;
}

- (NSInteger)outlineView:(NSOutlineView *)outlineView numberOfChildrenOfItem:(id)item {
    if(item == nil) {
        return [_items count];
    } else {
        return [(RWTreeViewItem*)item numberOfChildren];
    }
}

- (id)outlineView:(NSOutlineView *)outlineView objectValueForTableColumn:(NSTableColumn *)tableColumn
           byItem:(id)item {
    return [(RWTreeViewItem*)item title];
}

- (void) sizeColumnToFit:(NSTableColumn*)column {
    [column sizeToFitContent];
}

- (void)outlineView:(NSOutlineView *)outlineView setObjectValue:(id)object forTableColumn:(NSTableColumn *)tableColumn byItem:(id)item {
    [(RWTreeViewItem*)item setTitle:object];

    [self performSelectorOnMainThread:@selector(sizeColumnToFit:)
                        withObject:tableColumn
                      waitUntilDone:NO];
}

@end


////////////////////////////////////////////////////////////////////////////////////////

@implementation RWTreeView

- (NSOutlineView*)outlineView; {
    return (NSOutlineView*)[self documentView];
}

- (instancetype)initWithFrame:(NSRect)frameRect; {
    self = [super initWithFrame:frameRect];
    if(self) {
        [self setHasVerticalScroller:YES];
        [self setHasHorizontalScroller:YES];
        // https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSScrollView_Class/index.html#//apple_ref/occ/instm/NSScrollView/setScrollerStyle:
        // "This setting is automatically set at runtime..."
        //[self setScrollerStyle:NSScrollerStyleOverlay];
        // DocumentView of NSScrollView will be autosized.
        NSOutlineView* outlineView = [[[NSOutlineView alloc] initWithFrame:NSMakeRect(0, 0, 0, 0)] autorelease];
        NSTableColumn* column = [[[NSTableColumn alloc] initWithIdentifier:@"Column1"] autorelease];
        [column setEditable:NO];
        [outlineView addTableColumn:column];
        [outlineView setOutlineTableColumn:column];
        [outlineView setHeaderView:nil]; // Hiding header view.
        // setDataSource: The receiver maintains a weak reference to the data source.
        [outlineView setDataSource:[RWTreeViewDataSource new]];
        [self setDocumentView:outlineView];
    }
    return self;
}

- (BOOL)isEditable {
    return [[[self outlineView] outlineTableColumn] isEditable];
}

- (void)setEditable:(BOOL)editable {
    [[[self outlineView] outlineTableColumn] setEditable:editable];
}

- (void)sizeColumnToFit {
    [[[self outlineView] outlineTableColumn] sizeToFitContent];
}

- (void)dealloc {
    [self.outlineView.dataSource release];
    [super dealloc];
}

@end
