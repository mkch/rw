package menu

//#include "menu.h"
//#include "../types.h"
import "C"

import (
	"github.com/mkch/rw/internal/native/windows/nativeutil"
	"github.com/mkch/rw/native"
	"unsafe"
)

func CreateMenu() native.Handle {
	return native.Handle(C.PVOID(C.CreateMenu()))
}

func DestroyMenu(menu native.Handle) {
	if C.DestroyMenu(C.HMENU(C.PVOID(menu))) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func GetMenuItemCount(menu native.Handle) int {
	return int(C.GetMenuItemCount(C.HMENU(C.PVOID(menu))))
}

func GetMenuItemInfo(menu native.Handle, item uint, byPos bool, itemInfo *MenuItemInfo) {
	var info = newMENUITEMINFO(itemInfo)
	var bp C.BOOL
	if byPos {
		bp = C.TRUE
	}
	if C.GetMenuItemInfo(C.HMENU(C.PVOID(menu)), C.UINT(item), bp, info) == 0 {
		nativeutil.PanicWithLastError()
	}
	*itemInfo = *newMenuItemInfo(info)
}

func SetMenuItemInfo(menu native.Handle, item uint, byPos bool, info *MenuItemInfo) {
	var bp C.BOOL
	if byPos {
		bp = C.TRUE
	}
	if C.SetMenuItemInfo(C.HMENU(C.PVOID(menu)), C.UINT(item), bp, newMENUITEMINFO(info)) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func InsertMenuItem(menu native.Handle, item uint, byPos bool, info *MenuItemInfo) {
	var bp C.BOOL
	if byPos {
		bp = C.TRUE
	}
	if C.InsertMenuItem(C.HMENU(C.PVOID(menu)), C.UINT(item), bp, newMENUITEMINFO(info)) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func DeleteMenu(menu native.Handle, item uint, flags uint) {
	if C.DeleteMenu(C.HMENU(C.PVOID(menu)), C.UINT(item), C.UINT(flags)) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func RemoveMenu(menu native.Handle, item uint, flags uint) {
	if C.RemoveMenu(C.HMENU(C.PVOID(menu)), C.UINT(item), C.UINT(flags)) == 0 {
		nativeutil.PanicWithLastError()
	}
}

// The application must call the DrawMenuBar function whenever a menu changes, whether the menu is in a displayed window.
func DrawMenuBar(win native.Handle) bool {
	return C.DrawMenuBar(C.HWND(C.PVOID(win))) != 0
}

// Adds or removes highlighting from an item in a menu bar.
// The MF_HILITE and MF_UNHILITE flags can be used only with the HiliteMenuItem function; they cannot be used with the ModifyMenu function.
func HiliteMenuItem(win, menu native.Handle, item, hilite uint) {
	C.HiliteMenuItem(C.HWND(C.PVOID(win)), C.HMENU(C.PVOID(menu)), C.UINT(item), C.UINT(hilite))
}

type MenuItemInfo struct {
	Mask            uint
	Type            uint
	State           uint
	ID              uint
	SubMenu         native.Handle
	CheckedBitmap   native.Handle
	UncheckedBitmap native.Handle
	ItemData        uintptr
	TypeData        uintptr
	Cch             uint
	ItemBitmap      native.Handle
}

const (
	MF_BYCOMMAND       = uint(C.MF_BYCOMMAND)
	MF_BYPOSITION      = uint(C.MF_BYPOSITION)
	MF_ENABLED         = uint(C.MF_ENABLED)
	MF_GRAYED          = uint(C.MF_GRAYED)
	MF_DISABLED        = uint(C.MF_DISABLED)
	MF_BITMAP          = uint(C.MF_BITMAP)
	MF_CHECKED         = uint(C.MF_CHECKED)
	MF_MENUBARBREAK    = uint(C.MF_MENUBARBREAK)
	MF_MENUBREAK       = uint(C.MF_MENUBREAK)
	MF_OWNERDRAW       = uint(C.MF_OWNERDRAW)
	MF_POPUP           = uint(C.MF_POPUP)
	MF_SEPARATOR       = uint(C.MF_SEPARATOR)
	MF_STRING          = uint(C.MF_STRING)
	MF_UNCHECKED       = uint(C.MF_UNCHECKED)
	MF_DEFAULT         = uint(C.MF_DEFAULT)
	MF_SYSMENU         = uint(C.MF_SYSMENU)
	MF_HELP            = uint(C.MF_HELP)
	MF_END             = uint(C.MF_END)
	MF_RIGHTJUSTIFY    = uint(C.MF_RIGHTJUSTIFY)
	MF_MOUSESELECT     = uint(C.MF_MOUSESELECT)
	MF_INSERT          = uint(C.MF_INSERT)
	MF_CHANGE          = uint(C.MF_CHANGE)
	MF_APPEND          = uint(C.MF_APPEND)
	MF_DELETE          = uint(C.MF_DELETE)
	MF_REMOVE          = uint(C.MF_REMOVE)
	MF_USECHECKBITMAPS = uint(C.MF_USECHECKBITMAPS)
	MF_UNHILITE        = uint(C.MF_UNHILITE)
	MF_HILITE          = uint(C.MF_HILITE)
)

const (
	MIIM_STATE      = uint(C.MIIM_STATE)
	MIIM_ID         = uint(C.MIIM_ID)
	MIIM_SUBMENU    = uint(C.MIIM_SUBMENU)
	MIIM_CHECKMARKS = uint(C.MIIM_CHECKMARKS)
	MIIM_TYPE       = uint(C.MIIM_TYPE)
	MIIM_DATA       = uint(C.MIIM_DATA)
	MIIM_STRING     = uint(C.MIIM_STRING)
	MIIM_BITMAP     = uint(C.MIIM_BITMAP)
	MIIM_FTYPE      = uint(C.MIIM_FTYPE)

	MFT_BITMAP       = uint(C.MFT_BITMAP)
	MFT_MENUBARBREAK = uint(C.MFT_MENUBARBREAK)
	MFT_MENUBREAK    = uint(C.MFT_MENUBREAK)
	MFT_OWNERDRAW    = uint(C.MFT_OWNERDRAW)
	MFT_RADIOCHECK   = uint(C.MFT_RADIOCHECK)
	MFT_RIGHTJUSTIFY = uint(C.MFT_RIGHTJUSTIFY)
	MFT_SEPARATOR    = uint(C.MFT_SEPARATOR)
	MFT_RIGHTORDER   = uint(C.MFT_RIGHTORDER)
	MFT_STRING       = uint(C.MFT_STRING)

	MFS_CHECKED   = uint(C.MFS_CHECKED)
	MFS_DEFAULT   = uint(C.MFS_DEFAULT)
	MFS_DISABLED  = uint(C.MFS_DISABLED)
	MFS_ENABLED   = uint(C.MFS_ENABLED)
	MFS_GRAYED    = uint(C.MFS_GRAYED)
	MFS_HILITE    = uint(C.MFS_HILITE)
	MFS_UNCHECKED = uint(C.MFS_UNCHECKED)
	MFS_UNHILITE  = uint(C.MFS_UNHILITE)
)

func newMENUITEMINFO(itemInfo *MenuItemInfo) *C.MENUITEMINFO {
	var info = (*C.MENUITEMINFO)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.MENUITEMINFO{}))[0]))
	info.cbSize = C.UINT(unsafe.Sizeof(*info))
	info.fMask = C.UINT(itemInfo.Mask)
	info.fType = C.UINT(itemInfo.Type)
	info.fState = C.UINT(itemInfo.State)
	info.wID = C.UINT(itemInfo.ID)
	info.hSubMenu = C.HMENU(C.PVOID(itemInfo.SubMenu))
	info.hbmpChecked = C.HBITMAP(C.PVOID(itemInfo.CheckedBitmap))
	info.hbmpUnchecked = C.HBITMAP(C.PVOID(itemInfo.UncheckedBitmap))
	info.dwItemData = C.DWORD(itemInfo.ItemData)
	info.dwTypeData = C.LPWSTR(C.PVOID(itemInfo.TypeData))
	info.cch = C.UINT(itemInfo.Cch)
	info.hbmpItem = C.HBITMAP(C.PVOID(itemInfo.ItemBitmap))
	return info
}

func newMenuItemInfo(p *C.MENUITEMINFO) *MenuItemInfo {
	var Mask C.UINT
	var Type C.UINT
	var State C.UINT
	var ID C.UINT
	var SubMenu C.HMENU
	var CheckedBitmap C.HBITMAP
	var UncheckedBitmap C.HBITMAP
	var ItemData C.ULONG_PTR
	var TypeData C.LPWSTR
	var Cch C.UINT
	var ItemBitmap C.HBITMAP

	C.getMENUITEMINFO(p, &Mask, &Type, &State, &ID, &SubMenu, &CheckedBitmap, &UncheckedBitmap, &ItemData, &TypeData, &Cch, &ItemBitmap)
	return &MenuItemInfo{uint(Mask), uint(Type), uint(State), uint(ID), native.Handle(C.PVOID(SubMenu)), native.Handle(C.PVOID(CheckedBitmap)),
		native.Handle(C.PVOID(UncheckedBitmap)), uintptr(ItemData), uintptr(C.PVOID(TypeData)), uint(Cch), native.Handle(C.PVOID(ItemBitmap))}
}
