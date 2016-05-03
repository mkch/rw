#ifndef _RW_WINDOWSTYLE_H
#define _RW_WINDOWSTYLE_H

#include "../types.h"
#include "../../windowstyle.h"

void getWindowStyle(WindowStyleFeatures* features, UINT* pStyle, UINT* pExStyle);
void disableCloseButton(HWND hwnd);

#endif