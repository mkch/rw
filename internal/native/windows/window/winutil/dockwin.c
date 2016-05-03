#include <windows.h>
#include "dockwin.h"

static LRESULT CALLBACK DockWndProc(HWND hwnd,  UINT uMsg, WPARAM wParam, LPARAM lParam) {
	switch(uMsg){
		case WM_COMMAND:
			// http://www.programforge.com/2731204788/
			// http://stackoverflow.com/questions/30708760/edit-control-doesnt-generate-wm-command-messages
			if(lParam != 0) {
				HWND hChild = (HWND)lParam;
				HWND hRealParent = GetParent(hChild);
				if(hRealParent != hwnd){
					return SendMessage(hRealParent, uMsg, wParam, lParam);
				}
			}
			break;
	}
    return DefWindowProc(hwnd, uMsg, wParam, lParam);
}

WNDPROC GetDockWndProc() {
	return DockWndProc;
}

HWND VarHWND_MESSAGE = HWND_MESSAGE;