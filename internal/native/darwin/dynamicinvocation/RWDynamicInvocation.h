#import <Foundation/Foundation.h>

// RWInvocationArguments is a wrapper of NSInvocation.
// This class provides a convenient way to get every argument as a pointer.
// The problem to solve here is that we do not know the memory size of
// a single argument. In this class, a internal buffer that is large enough
// to holde any parameter is allocated and returned for every argument.
@interface RWInvocationArguments : NSObject {
    NSInvocation* invocation;
    void* buffer;
}

- (id) initWithInvocation:(NSInvocation*)anInvocation;
- (NSUInteger) numberOfArguments;
- (void*) getArgumentAtIndex:(NSInteger)index;
- (void) setReturnValue:(void*)value;
@end

// sel: The selecotr of this invocation. E.g. "windowDidBecomeMain:".
// args: The arguments and return value of this invocation.
//       userData: userData passed to initWithMethods:callback:userData.
// This callback is also called when the RWDynamicInvocation object is deallocated with sel and args set to NULL.
typedef void(*FN_DYNAMIC_INVOCATION_CALLBACL)(char* sel, RWInvocationArguments* args, void* userData);

@interface RWDynamicInvocation : NSObject {
    CFMutableDictionaryRef methods;
    FN_DYNAMIC_INVOCATION_CALLBACL callback;
    id delegate;
    void* userData;
}

// methods: The names and signatures of methods should be added and forwarded to callback.
//         The string is separated by '\0' character and terminated by double '\0' characters,
//         in format of "method1_name\0method1_signature\0method2_name\0method2_signature\0".
//         E.g. "windowDidBecomeMain:\0v@:@\0windowWillClose:\0v@:@\0windowShouldClose:\0i@:@\0"
//         For type encoding: https://developer.apple.com/library/mac/documentation/Cocoa/Conceptual/ObjCRuntimeGuide/Articles/ocrtTypeEncodings.html#//apple_ref/doc/uid/TP40008048-CH100
// aCallback: The implementation of all dynamic added methods.
- (id) initWithMethods:(char*) methods callback:(FN_DYNAMIC_INVOCATION_CALLBACL)aCallback userData:(void*) anUserData;
- (void) addMethod:(char*)method withSignature:(char*)signature;
- (void) setDelegate:(id)aDelegate;
- (id) delegate;

@end
