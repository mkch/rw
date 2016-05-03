#import "RWTextFieldDelegate.h"

@implementation RWTextFieldDelegate

- (BOOL)control:(NSControl*)control textView:(NSTextView*)textView doCommandBySelector:(SEL)commandSelector {
    if(!multiline) {
    	return NO;
    }

    if (commandSelector == @selector(insertNewline:)) {
        // new line action:
        // always insert a line-break character and don’t cause the receiver to end editing
        [textView insertNewlineIgnoringFieldEditor:self];
        return YES;
    } else if (commandSelector == @selector(insertTab:)) {
        // tab action:
        // always insert a tab character and don’t cause the receiver to end editing
        [textView insertTabIgnoringFieldEditor:self];
        return YES;
    }
 
    return NO;
}

- (BOOL) multiline {
	return multiline;
}

- (void) setMultiline:(BOOL)value {
	multiline = value;
}

@end


