#ifndef _RW_WINDOW_H
#define _RW_WINDOW_H

#include "../types.h"

WORD FnHIWORD(DWORD value);
WORD FnLOWWORD(DWORD value);
int GetCW_USEDEFAULT();

WNDENUMPROC GetEnumChildProc();

LONG_PTR WINAPI SetWindowLongPtr__LONG_PTR_HACK(HWND hWnd, int nIndex, LONG_PTR dwNewLong);
LONG_PTR WINAPI SetClassLongPtr__LONG_PTR_HACK(HWND hWnd, int nIndex, LONG_PTR dwNewLong);

WNDPROC GetDefWindowProcPtr();

extern LPCTSTR VarIDC_ARROW;
extern UINT VarHTTRANSPARENT;
HWND VarHWND_BOTTOM;
HWND VarHWND_NOTOPMOST;
HWND VarHWND_TOP;
HWND VarHWND_TOPMOST;

#endif
