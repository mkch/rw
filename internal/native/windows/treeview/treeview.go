package treeview

//#include "treeview.h"
import "C"

import (
	"unsafe"
	"github.com/kevin-yuan/rw/native"
	"github.com/kevin-yuan/rw/internal/mem"
	"github.com/kevin-yuan/rw/internal/native/windows/nativeutil"
	"github.com/kevin-yuan/rw/internal/native/windows/window"
	"github.com/kevin-yuan/rw/internal/native/windows/nativeutil/ustrings"
)

const (
	TVIS_BOLD = uint(C.TVIS_BOLD)
	TVIS_CUT = uint(C.TVIS_CUT)
	TVIS_DROPHILITED = uint(C.TVIS_DROPHILITED)
	TVIS_EXPANDED = uint(C.TVIS_EXPANDED)
	TVIS_EXPANDEDONCE = uint(C.TVIS_EXPANDEDONCE)
	//TVIS_EXPANDPARTIAL = uint(C.TVIS_EXPANDPARTIAL)
	TVIS_SELECTED = uint(C.TVIS_SELECTED)
	TVIS_OVERLAYMASK = uint(C.TVIS_OVERLAYMASK)
	TVIS_STATEIMAGEMASK = uint(C.TVIS_STATEIMAGEMASK)
	TVIS_USERMASK = uint(C.TVIS_USERMASK)
)

const (
	TVS_HASBUTTONS = uint(C.TVS_HASBUTTONS)
	TVS_HASLINES = uint(C.TVS_HASLINES)
	TVS_LINESATROOT = uint(C.TVS_LINESATROOT)
	TVS_CHECKBOXES  = uint(C.TVS_CHECKBOXES)
	TVS_FULLROWSELECT = uint(C.TVS_FULLROWSELECT)
	TVS_EDITLABELS = uint(C.TVS_EDITLABELS)
	TVS_DISABLEDRAGDROP = uint(C.TVS_DISABLEDRAGDROP)

)

var (
	TVI_FIRST = native.Handle(C.PVOID(C._TVI_FIRST))
	TVI_LAST = native.Handle(C.PVOID(C._TVI_LAST))
	TVI_ROOT = native.Handle(C.PVOID(C._TVI_ROOT))
	TVI_SORT = native.Handle(C.PVOID(C._TVI_SORT))
)

const (
	TVIF_CHILDREN = uint(C.TVIF_CHILDREN)
	TVIF_DI_SETITEM = uint(C.TVIF_DI_SETITEM)
	//TVIF_EXPANDEDIMAGE = uint(C.TVIF_EXPANDEDIMAGE)
	TVIF_HANDLE = uint(C.TVIF_HANDLE)
	TVIF_IMAGE = uint(C.TVIF_IMAGE)
	TVIF_INTEGRAL = uint(C.TVIF_INTEGRAL)
	TVIF_PARAM = uint(C.TVIF_PARAM)
	TVIF_SELECTEDIMAGE = uint(C.TVIF_SELECTEDIMAGE)
	TVIF_STATE = uint(C.TVIF_STATE)
	//TVIF_STATEEX = uint(C.TVIF_STATEEX)
	TVIF_TEXT = uint(C.TVIF_TEXT)
)

const (
	TVGN_CARET = uint(C.TVGN_CARET)
	TVGN_CHILD = uint(C.TVGN_CHILD)
	TVGN_DROPHILITE = uint(C.TVGN_DROPHILITE)
	TVGN_FIRSTVISIBLE = uint(C.TVGN_FIRSTVISIBLE)
	TVGN_NEXT = uint(C.TVGN_NEXT)
	//TVGN_NEXTSELECTED = uint(C.TVGN_NEXTSELECTED)
	TVGN_NEXTVISIBLE = uint(C.TVGN_NEXTVISIBLE)
	TVGN_PARENT = uint(C.TVGN_PARENT)
	TVGN_PREVIOUS = uint(C.TVGN_PREVIOUS)
	TVGN_PREVIOUSVISIBLE = uint(C.TVGN_PREVIOUSVISIBLE)
	TVGN_ROOT = uint(C.TVGN_ROOT)
)

const (
	TVE_COLLAPSE = uint(C.TVE_COLLAPSE)
	TVE_COLLAPSERESET = uint(C.TVE_COLLAPSERESET)
	TVE_EXPAND = uint(C.TVE_EXPAND)
	//TVE_EXPANDPARTIAL = uint(C.TVE_EXPANDPARTIAL)
	TVE_TOGGLE = uint(C.TVE_TOGGLE)
)

const (
    TVN_ENDLABELEDIT = uint(C.TVN_ENDLABELEDIT)
)

func TreeView_Expand(tv, item native.Handle, flag uint) {
	if C.SendMessage(C.HWND(C.PVOID(tv)), C.TVM_EXPAND, C.WPARAM(flag), C.LPARAM(item)) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func TreeView_SetItem(tv native.Handle, item *TvItemEx) {
	p := (*C.TVITEMEX)(mem.AllocAutoFree(uintptr(unsafe.Sizeof(C.TVITEMEX{}))))
	// Copy to c struct.
	newTVITEMEX(p, item)
	if C.SendMessage(C.HWND(C.PVOID(tv)), C.TVM_SETITEM, 0, C.LPARAM(uintptr(C.PVOID(p)))) == 0 {
		nativeutil.PanicWithLastError()
	} 
}

func TreeView_DeleteItem(tv, item native.Handle) {
	if C.SendMessage(C.HWND(C.PVOID(tv)), C.TVM_DELETEITEM, 0, C.LPARAM(item)) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func TreeView_GetNextItem(tv, item native.Handle, flag uint) native.Handle {
	return native.Handle(C.SendMessage(C.HWND(C.PVOID(tv)), C.TVM_GETNEXTITEM, C.WPARAM(flag), C.LPARAM(item)))
}

func TreeView_GetItem(tv native.Handle, item *TvItemEx)  {
	p := (*C.TVITEMEX)(mem.AllocAutoFree(uintptr(unsafe.Sizeof(C.TVITEMEX{}))))
	// Copy to c struct.
	newTVITEMEX(p, item)
	if C.SendMessage(C.HWND(C.PVOID(tv)), C.TVM_GETITEM, 0, C.LPARAM(uintptr(C.PVOID(p)))) == 0 {
		nativeutil.PanicWithLastError()
	}
	// Copy back.
	newTvItemEx(item, p)
}

// TreeView_GetCount returns the total number of items, all sub items included.
func TreeView_GetCount(tv native.Handle) int {
	return int(C.SendMessage(C.HWND(C.PVOID(tv)), C.TVM_GETCOUNT, 0, 0))
}

func TreeView_InsertItem(tv native.Handle, st *TvInsertStruct) (item native.Handle) {
	if item = native.Handle(C.SendMessage(C.HWND(C.PVOID(tv)), C.TVM_INSERTITEM, 0, C.LPARAM(uintptr(C.PVOID(newTVINSERTSTRUCT(st)))))); item == 0 {
		nativeutil.PanicWithLastError()
	}
	return
}

type TvInsertStruct struct {
	Parent      native.Handle
	InsertAfter native.Handle
	Item        TvItemEx
}

func newTVINSERTSTRUCT(st *TvInsertStruct) *C.TVINSERTSTRUCT {
	var ret = (*C.TVINSERTSTRUCT)(mem.AllocAutoFree(uintptr(unsafe.Sizeof(C.TVINSERTSTRUCT{}))))
	ret.hParent = C.HTREEITEM(C.PVOID(st.Parent))
	ret.hInsertAfter = C.HTREEITEM(C.PVOID(st.InsertAfter))
	// As Go doesn't have support for C's union type in the general case, C's union types are represented as a Go byte array with the same length.
	itemexPtr := (*C.TVITEMEX)(C.PVOID(uintptr(C.PVOID(&ret.hInsertAfter)) + unsafe.Sizeof(ret.hInsertAfter)))
	newTVITEMEX(itemexPtr, &st.Item)
	return ret	
}

type TvItemEx struct {
	Mask          uint
	Item          native.Handle
	State         uint
	StateMask     uint
	Text          uintptr
	TextMax       int
	Image         int
	SelectedImage int
	Children      int
	LParam        uintptr
	Integral      int
}

func (item *TvItemEx) TextString() string {
	return ustrings.FromUnicode(unsafe.Pointer(item.Text))
}

func (item *TvItemEx) SetTextString (str string) {
	item.Text = uintptr(ustrings.ToUnicodeAutoFree(str))
}

// Count of characters.
const ItemMaxTitleLen = uintptr(261)

// Make the buffer ready to receive text via TreeView_GetItem().
func (item *TvItemEx) PrepareTextBuffer() {
	buf := mem.AllocAutoFree(ItemMaxTitleLen*unsafe.Sizeof(C.WCHAR(0)))
	item.Text = uintptr(buf)
	item.TextMax = int(ItemMaxTitleLen)
}

func newTVITEMEX(ret *C.TVITEMEX, item *TvItemEx) {
	ret.mask = C.UINT(item.Mask)
	ret.hItem = C.HTREEITEM(C.PVOID(item.Item))
	ret.state = C.UINT(item.State)
	ret.stateMask = C.UINT(item.StateMask)
	ret.pszText = C.LPWSTR(C.PVOID(item.Text))
	ret.cchTextMax = C.int(item.TextMax)
	ret.iImage = C.int(item.Image)
	ret.iSelectedImage = C.int(item.SelectedImage)
	ret.cChildren = C.int(item.Children)
	ret.lParam = C.LPARAM(item.LParam)
	ret.iIntegral = C.int(item.Integral)
}

func newTvItemEx(ret *TvItemEx, item *C.TVITEMEX) {
	ret.Mask = uint(item.mask)
	ret.Item = native.Handle(C.PVOID(item.hItem))
	ret.State = uint(item.state)
	ret.StateMask = uint(item.stateMask)
	ret.Text = uintptr(C.PVOID(item.pszText))
	ret.TextMax = int(item.cchTextMax)
	ret.Image = int(item.iImage)
	ret.SelectedImage = int(item.iSelectedImage)
	ret.Children = int(item.cChildren)
	ret.LParam = uintptr(item.lParam)
	ret.Integral = int(item.iIntegral)
}

type TvItem struct {
	Mask uint
	Item native.Handle
	State uint
	StateMask uint
	Text uintptr
	TextMax int
	Image int
	SelectedImage int
	Children int
	LParam uintptr
}

type NmTvDispInfo struct {
	Hdr window.NmHdr
	Item TvItem
}

func ReadNmTvDispInfo(lParam uintptr) *NmTvDispInfo {
	info := (*C.NMTVDISPINFO)(unsafe.Pointer(lParam))
	return &NmTvDispInfo {
		Hdr: window.NmHdr {
				HandleFrom: native.Handle(C.PVOID(info.hdr.hwndFrom)),
				IdFrom: uintptr(info.hdr.idFrom),
				Code: uint(info.hdr.code),
			},
		Item: TvItem {
			Mask: uint(info.item.mask),
			Item: native.Handle(C.PVOID(info.item.hItem)),
			State: uint(info.item.state),
			StateMask: uint(info.item.stateMask),
			Text: uintptr(C.PVOID(info.item.pszText)),
			TextMax: int(info.item.cchTextMax),
			Image: int(info.item.iImage),
			SelectedImage: int(info.item.iSelectedImage),
			Children: int(info.item.cChildren),
			LParam: uintptr(info.item.lParam),
		},
	}
}

