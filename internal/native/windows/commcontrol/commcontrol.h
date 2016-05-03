#ifndef _RW_COMMCONTROL_H
#define _RW_COMMCONTROL_H

// https://msdn.microsoft.com/en-us/library/6sehtctf.aspx
#define WINVER 0x0601
#define _WIN32_WINNT 0x0601
#define _WIN32_IE 0x0600
#include <Windows.h>
#include <Commctrl.h>

#include "../types.h"

extern const LPCWSTR _WC_BUTTON;
extern const LPCWSTR _WC_EDIT;
extern const LPCWSTR _WC_TREEVIEW;

BOOL InitializeCommonControls();

#endif
