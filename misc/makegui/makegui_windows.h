#pragma once
#include <Windows.h>

extern DWORD _LOAD_LIBRARY_AS_DATAFILE;
extern LPCTSTR _RT_GROUP_ICON;
extern LPCTSTR _RT_ICON;
WORD LANG_NUTRAL;
WORD LANG_DEFAULT;
WCHAR* _MAKEINTRESOURCE(WORD id);

BOOL WINAPI EnumResourceNames_(HMODULE hModule, LPCTSTR lpszType, void* goEnumFunc);