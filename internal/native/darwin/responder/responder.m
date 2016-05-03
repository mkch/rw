#import <AppKit/NSResponder.h>
#import "responder.h"


bool NSResponder_acceptsFirstResponder(OBJC_PTR ptr) {
	return [(NSResponder*)ptr acceptsFirstResponder];
}
