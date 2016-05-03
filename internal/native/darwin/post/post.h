#ifndef _RW_POST_H
#define _RW_POST_H

#include "../types.h"

typedef void(*FN_POST_CALLBACK)(UINTPTR);

void initPostOnMainThread(FN_POST_CALLBACK safeCallback, FN_POST_CALLBACK unsafeCallback);
void postOnMainThread(UINTPTR userData, bool safe);

#endif
