#import <Foundation/NSNotification.h>
#import "notification.h"


OBJC_PTR NSNotification_object(OBJC_PTR ptr) {
	return [(NSNotification*)ptr object];
}

