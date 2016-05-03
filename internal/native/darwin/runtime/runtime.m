#import <objc/runtime.h>
#import "runtime.h"

UINTPTR registerSelector(char* name) {
	return (UINTPTR)(void*)sel_registerName(name);
}


