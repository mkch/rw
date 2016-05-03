#import <Foundation/NSThread.h>
#import <AppKit/NSResponder.h>
#import <objc/runtime.h>
#import "deallochook.h"
#import "_cgo_export.h"

@interface DeallocHook : NSObject {
    id obj;
}
@end

@implementation DeallocHook


-(id) initWithObject:(id)object {
    self = [super init];
    if(!self) {
        return nil;
    }
    obj = object;
    return self;
}

-(void) dealloc {
    callGlobalDeallocHook((void*)obj);
    [super dealloc];
}

@end


void hookDealloc(OBJC_PTR obj) {
    static int key;
    id hook = [[DeallocHook alloc] initWithObject:(id)obj];
    objc_setAssociatedObject((id)obj, &key, hook, OBJC_ASSOCIATION_RETAIN);
    [hook release];
}


// static IMP g_oldDeallocImp = NULL;

// static void newDeallocIMP(id self, SEL cmd) {
//     if([NSThread isMainThread] && g_deallocHook) {
//         g_deallocHook((void*)self);
//     }
//     g_oldDeallocImp(self, cmd);
// }


// void setDeallocHook(FN_DEALLOC_HOOK hook) {
// 	g_deallocHook = hook;
// 	// NSMenu NSMenuItem inherit from NSObject.
// 	Method mtdToHook = class_getInstanceMethod([NSObject class], @selector(dealloc));
//     //Method mtdHook = class_getInstanceMethod([DeallocHook class], @selector(releaseHook));
//     //g_oldDeallocImp = method_setImplementation(mtdToHook, method_getImplementation(mtdHook));
//     g_oldDeallocImp = method_setImplementation(mtdToHook, (IMP)newDeallocIMP);
// }
