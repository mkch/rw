
@interface RWWindow : NSWindow {
	BOOL _disabled;
}

- (BOOL) enabled ;
- (void) setEnabled:(BOOL)aEnabled;
- (void)sendEvent:(NSEvent *)event;

@end