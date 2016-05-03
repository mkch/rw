#ifndef _RW_POST_H
#define _RW_POST_H

#include <windows.h>
#include "../types.h"

#define WM_GO_SAFE_POST (WM_APP + 11)
#define WM_GO_UNSAFE_POST (WM_GO_SAFE_POST + 1)

HWND createPostMessageOnlyWindow();


#endif
