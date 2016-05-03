
#import <objc/runtime.h>
#import <CoreFoundation/CoreFoundation.h>
#import <Foundation/Foundation.h>
#import <string.h>
#import "RWDynamicInvocation.h"

// https://developer.apple.com/library/mac/documentation/Cocoa/Conceptual/ObjCRuntimeGuide/Articles/ocrtTypeEncodings.html
// 

@implementation RWInvocationArguments

- (id) initWithInvocation:(NSInvocation*)anInvocation {
    self = [super init];
    if(self) {
        invocation = anInvocation;
        const NSMethodSignature* signature = [invocation methodSignature];
        const NSUInteger frameLength = [signature frameLength];
        if(frameLength > 0) {
            // Large enough to hold all arguments. Actually we only use it to hold a single one at a time.
            buffer = malloc(frameLength);
        }
    }
    return self;
}

-(void) dealloc {
    free(buffer);
    [super dealloc];
}

- (NSUInteger) numberOfArguments {
    return [[invocation methodSignature] numberOfArguments]-2;
}

- (void*) getArgumentAtIndex:(NSInteger)index {
    [invocation getArgument:buffer atIndex:index+2];
    return buffer;
}


- (void) setReturnValue:(void*)value {
    [invocation setReturnValue:value];
}

@end

@implementation RWDynamicInvocation

- (id) initWithMethods:(char*) aMethods callback:(FN_DYNAMIC_INVOCATION_CALLBACL)aCallback userData:(void*) anUserData {
    self = [super init];
    if(self) {
        methods = CFDictionaryCreateMutable(NULL, 0, NULL, NULL);
        callback = aCallback;
        userData = anUserData;
    }
    
    if(aMethods) {
        char* selName = aMethods;
        char* mtdSignature = NULL;
        size_t len = strlen(selName);
        while(len) {
            mtdSignature = selName + len + 1;
            [self addMethod:selName withSignature:mtdSignature];
            
            selName = mtdSignature + strlen(mtdSignature) + 1;
            len = strlen(selName);
        }
    }
    return self;
}

-(void) dealloc {
    callback(NULL, NULL, userData); // Notify client the deallocation.
    [delegate release];
    CFRelease(methods);
    [super dealloc];
}

- (void) setDelegate:(id)aDelegate {
    delegate = [aDelegate retain];
}
- (id) delegate {
    return delegate;
}

-(void) addMethod:(char*)aMethod withSignature:(char*)aSignature {
    SEL sel = sel_registerName(aMethod);
    NSMethodSignature* signature = [NSMethodSignature signatureWithObjCTypes:aSignature];
    CFDictionaryAddValue(methods, sel, signature);
}

- (BOOL) respondsToSelector:(SEL)aSelector {
    if(!CFDictionaryContainsKey(methods, aSelector)) {
        return (delegate && [delegate respondsToSelector:aSelector]) || [super respondsToSelector:aSelector];
    }
    return YES;
}

- (NSMethodSignature *)methodSignatureForSelector:(SEL)aSelector {
    NSMethodSignature* signature = CFDictionaryGetValue(methods, aSelector);
    if(signature == nil && delegate) {
        signature = [delegate methodSignatureForSelector:aSelector];
    }
    if(signature == nil) {
        return [super methodSignatureForSelector:aSelector];
    }
    return signature;
}

- (void)forwardInvocation:(NSInvocation *)anInvocation {
    SEL selector = [anInvocation selector];
    if(!CFDictionaryContainsKey(methods, selector)) {
        if(delegate && [delegate respondsToSelector:selector]) {
            [anInvocation invokeWithTarget:delegate];
        } else {
            [super forwardInvocation:anInvocation];
        }
        return;
    }
    
    RWInvocationArguments* args = [[RWInvocationArguments alloc] initWithInvocation:anInvocation];
    callback((char*)sel_getName(selector), args, userData);
    [args release];
}



@end
