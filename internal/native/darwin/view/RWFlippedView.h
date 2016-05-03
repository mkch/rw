
@interface RWFlippedView : NSView {
	BOOL _acceptFirstResponsder;
	NSColor* _backgroundColor;
}

- (void)setBackgroundColor:(NSColor*)aColor;
- (NSColor*)backgroundColor;
- (void)setAcceptFirstResponder:(BOOL)aAccept;

@end