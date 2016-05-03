#include <windows.h>
#include "windowstyle.h"

void getWindowStyle(WindowStyleFeatures* features, UINT* pStyle, UINT* pExStyle) {
    *pStyle = WS_SYSMENU|WS_POPUP;
    *pExStyle = 0;
    if(features->hasBorder) {
        *pStyle |= (WS_BORDER|WS_OVERLAPPED);
    }
    if(features->hasTitle) {
        *pStyle |= (WS_CAPTION|WS_OVERLAPPED);
    }
    if(features->hasCloseButton) {
        
    }
    if(features->hasMinimizeButton) {
        *pStyle |= WS_MINIMIZEBOX;
    }
    if(features->hasMaximizeButton) {
        *pStyle |= WS_MAXIMIZEBOX;
    }
    if(features->resizable) {
        *pStyle |= WS_SIZEBOX;
    }
}

void disableCloseButton(HWND hwnd) {
    EnableMenuItem(GetSystemMenu(hwnd, FALSE), SC_CLOSE, MF_BYCOMMAND | MF_DISABLED | MF_GRAYED);
}