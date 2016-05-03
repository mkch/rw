#include "makegui_windows.h"
#include "_cgo_export.h"

DWORD _LOAD_LIBRARY_AS_DATAFILE = LOAD_LIBRARY_AS_DATAFILE;
LPCTSTR _RT_GROUP_ICON = RT_GROUP_ICON;
LPCTSTR _RT_ICON = RT_ICON;
WORD LANG_NUTRAL = MAKELANGID(LANG_NEUTRAL, SUBLANG_NEUTRAL);
WORD LANG_DEFAULT = MAKELANGID(LANG_SYSTEM_DEFAULT, SUBLANG_SYS_DEFAULT);
WCHAR* _MAKEINTRESOURCE(WORD id) {
	return MAKEINTRESOURCE(id);
}


BOOL CALLBACK EnumResNameProc(HMODULE hModule, LPCTSTR lpszType, LPTSTR lpszName, LONG_PTR lParam) {
	go_enumResourceNamesCallback(lpszName, (void*)lParam);
	return TRUE;
}

BOOL WINAPI EnumResourceNames_(HMODULE hModule, LPCTSTR lpszType, void* goEnumFunc) {
	return EnumResourceNames(hModule, lpszType, EnumResNameProc, (LONG_PTR)goEnumFunc);
}