#include <windows.h>


DWORD fRGB(byte r, byte g, byte b) {
	return RGB(r, g, b);
}

byte fGetRValue(DWORD color) {
	return GetRValue(color);
}

byte fGetGValue(DWORD color) {
	return GetGValue(color);
}

byte fGetBValue(DWORD color) {
	return GetBValue(color);
}