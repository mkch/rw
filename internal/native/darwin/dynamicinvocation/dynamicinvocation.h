#ifndef RW_DYNAMIC_INVOCATION_H
#define RW_DYNAMIC_INVOCATION_H

#include "../types.h"


OBJC_PTR RWDynamicInvocation_initWithMethodsUserData(char* methods, UINTPTR userData);
void RWDynamicInvocation_addMethodwithSignature(OBJC_PTR ptr, char* method, char* signature);
void RWDynamicInvocation_setDelegateRetain(OBJC_PTR ptr, OBJC_PTR delegate);
OBJC_PTR RWDynamicInvocation_delegate(OBJC_PTR ptr);

unsigned int RWInvocationArguments_numberOfArguments(OBJC_PTR ptr);
void* RWInvocationArguments_getArgumentAtIndex(OBJC_PTR ptr, unsigned int index);
void RWInvocationArguments_setReturnValue(OBJC_PTR ptr, void* value);

#endif
