#import <Foundation/NSNotification.h>
#import <Foundation/NSString.h>
#import "notificationcenter.h"

OBJC_PTR NSNotificationCenter_defaultCenter() {
	return [NSNotificationCenter defaultCenter];
}

void NSNotificationCenter_addObserver_selector_name_object(OBJC_PTR ptr, OBJC_PTR observer, PVOID sel, char* name, OBJC_PTR object) {
	NSString* notificationName = name ? [NSString stringWithUTF8String:name] : nil;
	[(NSNotificationCenter*)ptr addObserver:(id)observer
								   selector:(SEL)sel
								       name:notificationName
								     object:(id)object];
}

void NSNotificationCenter_removeObserver_name_object(OBJC_PTR ptr, OBJC_PTR observer, char* name, OBJC_PTR object) {
	NSString* notificationName = name ? [NSString stringWithUTF8String:name] : nil;
	[(NSNotificationCenter*)ptr removeObserver:(id)observer
										  name:notificationName
										object:(id)object];
}