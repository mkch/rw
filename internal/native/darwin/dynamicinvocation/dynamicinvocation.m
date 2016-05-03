#import "RWDynamicInvocation.h"
#import "dynamicinvocation.h"
#import "_cgo_export.h"


static void callback(char* sel, RWInvocationArguments* args, void* userData) {
	callDynamicInvocationCallback((UINTPTR)userData, sel, args);
}

OBJC_PTR RWDynamicInvocation_initWithMethodsUserData(char* methods, UINTPTR userData) {
	return [[RWDynamicInvocation alloc] initWithMethods:methods callback:callback userData:(void*)userData];
}

void RWDynamicInvocation_addMethodwithSignature(OBJC_PTR ptr, char* method, char* signature) {
	[(RWDynamicInvocation*)ptr addMethod:method withSignature:signature];
}

void RWDynamicInvocation_setDelegateRetain(OBJC_PTR ptr, OBJC_PTR delegate) {
	[(RWDynamicInvocation*)ptr setDelegate:(id)delegate];
}

OBJC_PTR RWDynamicInvocation_delegate(OBJC_PTR ptr) {
	return [(RWDynamicInvocation*)ptr delegate];
}

unsigned int RWInvocationArguments_numberOfArguments(OBJC_PTR ptr) {
	return [(RWInvocationArguments*)ptr numberOfArguments];
}

void* RWInvocationArguments_getArgumentAtIndex(OBJC_PTR ptr, unsigned int index) {
	return [(RWInvocationArguments*)ptr getArgumentAtIndex:index];
}

void RWInvocationArguments_setReturnValue(OBJC_PTR ptr, void* value) {
	[(RWInvocationArguments*)ptr setReturnValue:value];
}
