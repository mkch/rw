#include "menu.h"

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
	HBITMAP*   hbmpItem) {

	*fMask = info->fMask;
	*fType = info->fType;
	*fState = info->fState;
	*wID = info->wID;
	*hSubMenu = info->hSubMenu;
	*hbmpChecked = info->hbmpChecked;
	*hbmpUnchecked = info->hbmpUnchecked;
	*dwItemData = info->dwItemData;
	*dwTypeData = info->dwTypeData;
	*cch = info->cch;
	*hbmpItem = info->hbmpItem;
}
