
#import <AppKit/NSControl.h>
#import <AppKit/NSTextView.h>

@interface RWTextFieldDelegate : NSObject {
	BOOL multiline;
}

- (BOOL) multiline;
- (void) setMultiline:(BOOL)value;

@end