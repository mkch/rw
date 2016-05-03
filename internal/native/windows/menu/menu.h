#ifndef _RW_MENU_H
#define _RW_MENU_H

#include <windows.h>
#include "../types.h"

void getMENUITEMINFO(
	MENUITEMINFO* info,
	UINT*      fMask,
	UINT*      fType,
	UINT*      fState,
	UINT*      wID,
	HMENU*     hSubMenu,
	HBITMAP*   hbmpChecked,
	HBITMAP*   hbmpUnchecked,
	ULONG_PTR* dwItemData,
	LPWSTR*    dwTypeData,
	UINT*      cch,
	HBITMAP*   hbmpItem);

#endif
