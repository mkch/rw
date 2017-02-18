package window

//#include <windows.h>
//#include "window.h"
import "C"

import (
	"github.com/mkch/rw/internal/native/windows/nativeutil"
	"github.com/mkch/rw/internal/stackescape"
	"github.com/mkch/rw/native"
	"github.com/mkch/rw/util/ustr"
	"unsafe"
)

type NmHdr struct {
	HandleFrom native.Handle
	IdFrom     uintptr
	Code       uint
}

func ReadNmHdr(lParam uintptr) *NmHdr {
	hdr := (*C.NMHDR)(unsafe.Pointer(lParam))
	return &NmHdr{
		HandleFrom: native.Handle(C.PVOID(hdr.hwndFrom)),
		IdFrom:     uintptr(hdr.idFrom),
		Code:       uint(hdr.code),
	}
}

func HIWORD(v uint) uint16 {
	return uint16(C.FnHIWORD(C.DWORD(v)))
}

func LOWORD(v uint) uint16 {
	return uint16(C.FnLOWWORD(C.DWORD(v)))
}

func GetModuleHandle(moduleName *string) native.Handle {
	if moduleName != nil {
		unicodeModuleName := ustr.CStringUtf16(*moduleName)
		return native.Handle(C.Ptr(C.GetModuleHandle(C.LPWSTR(unicodeModuleName))))
	}
	return native.Handle(C.Ptr(C.GetModuleHandle(nil)))
}

const (
	WS_OVERLAPPEDWINDOW = uint(C.WS_OVERLAPPEDWINDOW)
	WS_OVERLAPPED       = uint(C.WS_OVERLAPPED)
	WS_POPUP            = uint(C.WS_POPUP)
	WS_CHILD            = uint(C.WS_CHILD)
	WS_BORDER           = uint(C.WS_BORDER)
	WS_VISIBLE          = uint(C.WS_VISIBLE)
	WS_CAPTION          = uint(C.WS_CAPTION)
	WS_SYSMENU          = uint(C.WS_CAPTION)
	WS_MAXIMIZEBOX      = uint(C.WS_CAPTION)
	WS_MINIMIZEBOX      = uint(C.WS_CAPTION)
	WS_SIZEBOX          = uint(C.WS_CAPTION)
	WS_TABSTOP          = uint(C.WS_TABSTOP)

	WS_EX_CLIENTEDGE    = uint(C.WS_EX_CLIENTEDGE)
	WS_EX_CONTROLPARENT = uint(C.WS_EX_CONTROLPARENT)
	WS_EX_TOPMOST       = uint(C.WS_EX_TOPMOST)
)

const (
	MONITORINFOF_PRIMARY = uint(C.MONITORINFOF_PRIMARY)
)

type MonitorInfo struct {
	MonitorRectLeft, MonitorRectTop, MonitorRectWidth, MonitorRectHeight,
	WorkRectLeft, WorkRectTop, WorkRectWidth, WorkRectHeight int
	Flags uint
}

func GetMonitorInfo(monitor native.Handle) *MonitorInfo {
	info := (*C.MONITORINFO)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.MONITORINFO{}))[0]))
	info.cbSize = C.DWORD(unsafe.Sizeof(C.MONITORINFO{}))
	if C.GetMonitorInfo(C.HMONITOR(C.PVOID(monitor)), info) == 0 {
		nativeutil.PanicWithLastError()
	}
	return &MonitorInfo{
		MonitorRectLeft: int(info.rcMonitor.left), MonitorRectTop: int(info.rcMonitor.top), MonitorRectWidth: int(info.rcMonitor.right - info.rcMonitor.left), MonitorRectHeight: int(info.rcMonitor.bottom - info.rcMonitor.top),
		WorkRectLeft: int(info.rcWork.left), WorkRectTop: int(info.rcWork.top), WorkRectWidth: int(info.rcWork.right - info.rcWork.left), WorkRectHeight: int(info.rcWork.bottom - info.rcWork.top),
		Flags: uint(info.dwFlags),
	}
}

// Call C.GetMonitor with MONITORINFOEX as the 2nd argument, so the device name is retrieved.
func GetMonitorInfoEx(monitor native.Handle) (monitorInfo *MonitorInfo, deviceName string) {
	info := (*C.MONITORINFOEX)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.MONITORINFOEX{}))[0]))
	info.cbSize = C.DWORD(unsafe.Sizeof(C.MONITORINFOEX{}))
	if C.GetMonitorInfo(C.HMONITOR(C.PVOID(monitor)), (*C.MONITORINFO)(C.PVOID(info))) == 0 {
		nativeutil.PanicWithLastError()
	}
	return &MonitorInfo{
			MonitorRectLeft: int(info.rcMonitor.left), MonitorRectTop: int(info.rcMonitor.top), MonitorRectWidth: int(info.rcMonitor.right - info.rcMonitor.left), MonitorRectHeight: int(info.rcMonitor.bottom - info.rcMonitor.top),
			WorkRectLeft: int(info.rcWork.left), WorkRectTop: int(info.rcWork.top), WorkRectWidth: int(info.rcWork.right - info.rcWork.left), WorkRectHeight: int(info.rcWork.bottom - info.rcWork.top),
			Flags: uint(info.dwFlags),
		},
		ustr.GoStringFromUtf16(unsafe.Pointer(&info.szDevice[0]))
}

const (
	MONITOR_DEFAULTTONEAREST = uint(C.MONITOR_DEFAULTTONEAREST)
	MONITOR_DEFAULTTONULL    = uint(C.MONITOR_DEFAULTTONULL)
	MONITOR_DEFAULTTOPRIMARY = uint(C.MONITOR_DEFAULTTOPRIMARY)
)

func MonitorFromWindow(hwnd native.Handle, flags uint) native.Handle {
	return native.Handle(C.PVOID(C.MonitorFromWindow(C.HWND(C.PVOID(hwnd)), C.DWORD(flags))))
}

var (
	CW_USEDEFAULT = int(C.GetCW_USEDEFAULT())
)

func ScreenToClient(hwnd native.Handle, x, y *int) {
	var pt = (*C.POINT)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.POINT{}))[0]))
	pt.x, pt.y = C.LONG(*x), C.LONG(*y)
	if C.ScreenToClient(C.HWND(C.PVOID(hwnd)), pt) == 0 {
		nativeutil.PanicWithLastError()
	}
	*x, *y = int(pt.x), int(pt.y)
}

func ClientToScreen(hwnd native.Handle, x, y *int) {
	var pt = (*C.POINT)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.POINT{}))[0]))
	pt.x, pt.y = C.LONG(*x), C.LONG(*y)
	if C.ClientToScreen(C.HWND(C.PVOID(hwnd)), pt) == 0 {
		nativeutil.PanicWithLastError()
	}
	*x, *y = int(pt.x), int(pt.y)
}

func CreateWindowEx(exStyle uint, className uintptr, windowName string, style uint,
	x int, y int, width int, height int, parent native.Handle, menu native.Handle, instance native.Handle, lParam unsafe.Pointer) native.Handle {

	pWindowName := ustr.CStringUtf16(windowName)

	hwnd := C.CreateWindowEx(C.DWORD(exStyle), C.LPWSTR(C.PVOID(className)), C.LPWSTR(pWindowName), C.DWORD(style), C.int(x), C.int(y), C.int(width), C.int(height), C.HWND(C.Ptr(parent)), C.HMENU(C.Ptr(menu)), C.HMODULE(C.Ptr(instance)), C.LPVOID(lParam))
	if hwnd == nil {
		nativeutil.PanicWithLastError()
	}
	return native.Handle(C.Ptr(hwnd))
}

func DestroyWindow(h native.Handle) {
	if C.DestroyWindow(C.HWND(C.Ptr(h))) == 0 {
		nativeutil.PanicWithLastError()
	}
}

var (
	SW_SHOW   = int(C.SW_SHOW)
	SW_HIDE   = int(C.SW_HIDE)
	SW_SHOWNA = int(C.SW_SHOWNA)
)

func ShowWindow(h native.Handle, cmd int) bool {
	return C.ShowWindow(C.HWND(C.Ptr(h)), C.int(cmd)) != 0
}

func GetWindowText(handle native.Handle) string {
	len := C.GetWindowTextLength(C.HWND(C.Ptr(handle)))
	if len == 0 {
		return ""
	}
	buf := C.LPWSTR(unsafe.Pointer(&make([]byte, uintptr(len+1)*unsafe.Sizeof(C.WCHAR(0)))[0]))
	C.GetWindowText(C.HWND(C.Ptr(handle)), buf, len+1)
	return ustr.GoStringFromUtf16(unsafe.Pointer(buf))
}

func PostQuitMessage(exitCode int) {
	C.PostQuitMessage(C.int(exitCode))
}

func PostMessage(h native.Handle, msg uint, wParam uintptr, lParam uintptr) {
	if C.PostMessage(C.HWND(C.Ptr(h)), C.UINT(msg), C.WPARAM(wParam), C.LPARAM(lParam)) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func GetClientRect(h native.Handle) (x, y, width, height int) {
	var rect = (*C.RECT)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.RECT{}))[0]))
	if C.GetClientRect(C.HWND(C.Ptr(h)), rect) == 0 {
		nativeutil.PanicWithLastError()
	}
	return int(rect.left), int(rect.top), int(rect.right - rect.left), int(rect.bottom - rect.top)
}

// The dimensions are given in screen coordinates
func GetWindowRect(h native.Handle) (x, y, width, height int) {
	var rect = (*C.RECT)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.RECT{}))[0]))
	if C.GetWindowRect(C.HWND(C.Ptr(h)), rect) == 0 {
		nativeutil.PanicWithLastError()
	}
	return int(rect.left), int(rect.top), int(rect.right - rect.left), int(rect.bottom - rect.top)
}

// The dimensions are given in parent client coordinates
func GetChildWindowRect(handle, parent native.Handle) (x, y, width, height int) {
	var rect = (*C.RECT)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.RECT{}))[0]))
	if C.GetWindowRect(C.HWND(C.Ptr(handle)), rect) == 0 ||
		C.MapWindowPoints(nil, C.HWND(C.Ptr(parent)), C.LPPOINT(unsafe.Pointer(&rect)), 2) == 0 {
		nativeutil.PanicWithLastError()
	}
	return int(rect.left), int(rect.top), int(rect.right - rect.left), int(rect.bottom - rect.top)
}

type WindowEnumCallback func(h native.Handle, lParam interface{}) bool

type enumChildWindowsParam struct {
	Callback WindowEnumCallback
	Param    interface{}
}

//export callEnumChildWindowCallback
func callEnumChildWindowCallback(h C.HWND, paramId uintptr) bool {
	param := stackescape.Get(stackescape.Id(paramId)).(*enumChildWindowsParam)
	return param.Callback(native.Handle(C.Ptr(h)), param.Param)
}

func EnumChildWindows(h native.Handle, callback WindowEnumCallback, lParam interface{}) (ok bool) {
	lp := stackescape.Add(&enumChildWindowsParam{callback, lParam})
	ok = bool(C.EnumChildWindows(C.HWND(C.Ptr(h)), C.GetEnumChildProc(), C.LPARAM(lp)) != 0)
	stackescape.Remove(lp)
	return
}

const (
	SWP_NOACTIVATE    uint = uint(C.SWP_NOACTIVATE)
	SWP_NOOWNERZORDER      = uint(C.SWP_NOOWNERZORDER)
	SWP_NOZORDER           = uint(C.SWP_NOZORDER)
	SWP_NOSIZE             = uint(C.SWP_NOSIZE)
	SWP_NOMOVE             = uint(C.SWP_NOMOVE)
	SWP_FRAMECHANGED       = uint(C.SWP_FRAMECHANGED)
)

var (
	HWND_BOTTOM    = native.Handle(C.PVOID(C.VarHWND_BOTTOM))
	HWND_NOTOPMOST = native.Handle(C.PVOID(C.VarHWND_NOTOPMOST))
	HWND_TOP       = native.Handle(C.PVOID(C.VarHWND_TOP))
	HWND_TOPMOST   = native.Handle(C.PVOID(C.VarHWND_TOPMOST))
)

func SetWindowPos(h, insertAfter native.Handle, x, y, cx, cy int, flags uint) {
	if C.SetWindowPos(C.HWND(C.Ptr(h)), C.HWND(C.Ptr(insertAfter)), C.int(x), C.int(y), C.int(cx), C.int(cy), C.UINT(flags)) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func SetWindowText(h native.Handle, text string) {
	if C.SetWindowText(C.HWND(C.Ptr(h)), C.LPWSTR(ustr.CStringUtf16(text))) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func SetParent(child, parent native.Handle) {
	if C.SetParent(C.HWND(C.Ptr(child)), C.HWND(C.Ptr(parent))) == nil {
		nativeutil.PanicWithLastError()
	}
}

func GetParent(h native.Handle) (parent native.Handle) {
	p := C.GetParent(C.HWND(C.Ptr(h)))
	return native.Handle(C.Ptr(p))
}

var (
	GWLP_WNDPROC = int(C.GWLP_WNDPROC)
	GWL_EXSTYLE  = int(C.GWL_EXSTYLE)
	GWL_STYLE    = int(C.GWL_STYLE)
)

func GetWindowLongPtr(h native.Handle, index int) uintptr {
	return uintptr(C.GetWindowLongPtr(C.HWND(C.Ptr(h)), C.int(index)))
}

func SetWindowLongPtr(h native.Handle, index int, newLong uintptr) uintptr {
	return uintptr(C.SetWindowLongPtr__LONG_PTR_HACK(C.HWND(C.Ptr(h)), C.int(index), C.LONG_PTR(newLong)))
}

func CallWindowProc(f uintptr, handle native.Handle, msg uint, wParam, lParam uintptr) uintptr {
	return uintptr(C.CallWindowProc(C.WNDPROC(C.Ptr(f)), C.HWND(C.Ptr(handle)), C.UINT(msg), C.WPARAM(wParam), C.LPARAM(lParam)))
}

const (
	MSGF_DIALOGBOX = uint(C.MSGF_DIALOGBOX)
	MSGF_MENU      = uint(C.MSGF_MENU)
)

const (
	WA_INACTIVE = uint(C.WA_INACTIVE)
)

const (
	WM_ACTIVATE    = uint(C.WM_ACTIVATE)
	WM_DESTROY     = uint(C.WM_DESTROY)
	WM_NCDESTROY   = uint(C.WM_NCDESTROY)
	WM_COMMAND     = uint(C.WM_COMMAND)
	WM_SIZE        = uint(C.WM_SIZE)
	WM_CLOSE       = uint(C.WM_CLOSE)
	WM_ENTERIDLE   = uint(C.WM_ENTERIDLE)
	WM_NOTIFY      = uint(C.WM_NOTIFY)
	WM_ERASEBKGND  = uint(C.WM_ERASEBKGND)
	WM_NCHITTEST   = uint(C.WM_NCHITTEST)
	WM_LBUTTONDOWN = uint(C.WM_LBUTTONDOWN)
	WM_KEYDOWN     = uint(C.WM_KEYDOWN)
	WM_SYSKEYDOWN  = uint(C.WM_SYSKEYDOWN)
	WM_MENUCOMMAND = uint(C.WM_MENUCOMMAND)
)

var (
	HTTRANSPARENT = uintptr(C.VarHTTRANSPARENT)
)

var (
	BN_CLICKED = uint(C.BN_CLICKED)
)

func InvalidateRect(hwnd native.Handle, l, t, r, b int, erase bool) bool {
	var rect = (*C.RECT)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.RECT{}))[0]))
	rect.left = C.LONG(l)
	rect.top = C.LONG(t)
	rect.right = C.LONG(r)
	rect.bottom = C.LONG(b)
	var bErase C.BOOL = 0
	if erase {
		bErase = 1
	}
	return C.InvalidateRect(C.HWND(C.PVOID(hwnd)), rect, bErase) != 0
}

func InvalidateRectNull(hwnd native.Handle, erase bool) bool {
	var b C.BOOL = 0
	if erase {
		b = 1
	}
	return C.InvalidateRect(C.HWND(C.PVOID(hwnd)), nil, b) != 0
}

func DefWindowProc(handle native.Handle, msg uint, wParam, lParam uintptr) uintptr {
	return uintptr(C.DefWindowProc(C.HWND(C.Ptr(handle)), C.UINT(msg), C.WPARAM(wParam), C.LPARAM(lParam)))
}

func SetProp(handle native.Handle, prop uintptr, data uintptr) bool {
	return C.SetProp(C.HWND(C.Ptr(handle)), C.LPWSTR(C.PVOID(prop)), C.HANDLE(data)) != 0
}

func SetPropString(handle native.Handle, prop string, data uintptr) bool {
	return C.SetProp(C.HWND(C.Ptr(handle)), C.LPWSTR(ustr.CStringUtf16(prop)), C.HANDLE(data)) != 0
}

func RemoveProp(handle native.Handle, prop uintptr) uintptr {
	return uintptr(C.RemoveProp(C.HWND(C.Ptr(handle)), C.LPWSTR(C.PVOID(prop))))
}

func RemovePropString(handle native.Handle, prop string) uintptr {
	return RemoveProp(handle, uintptr(ustr.CStringUtf16(prop)))
}

func GetPropString(handle native.Handle, prop string) uintptr {
	return GetProp(handle, uintptr(ustr.CStringUtf16(prop)))
}

func GetProp(handle native.Handle, prop uintptr) uintptr {
	return uintptr(C.GetProp(C.HWND(C.Ptr(handle)), C.LPWSTR(C.PVOID(prop))))
}

var (
	GA_PARENT    = uint(C.GA_PARENT)
	GA_ROOT      = uint(C.GA_ROOT)
	GA_ROOTOWNER = uint(C.GA_ROOTOWNER)
)

func GetAncestor(handle native.Handle, falgs uint) native.Handle {
	return native.Handle(C.Ptr(C.GetAncestor(C.HWND(C.Ptr(handle)), C.UINT(falgs))))
}

func EnableWindow(handle native.Handle, enable bool) {
	var value C.BOOL = 0
	if enable {
		value = 1
	}
	C.EnableWindow(C.HWND(C.Ptr(handle)), value)
}

func IsWindowEnabled(handle native.Handle) bool {
	return C.IsWindowEnabled(C.HWND(C.Ptr(handle))) != 0
}

func SendMessage(handle native.Handle, msg uint, wParam, lParam uintptr) uintptr {
	return uintptr(C.SendMessage(C.HWND(C.PVOID(handle)), C.UINT(msg), C.WPARAM(wParam), C.LPARAM(lParam)))
}

func SetFocus(handle native.Handle) bool {
	return C.SetFocus(C.HWND(C.PVOID(handle))) != nil
}

func GetFocus() native.Handle {
	return native.Handle(C.PVOID(C.GetFocus()))
}

func SetMenu(hwnd, hmenu native.Handle) {
	if C.SetMenu(C.HWND(C.PVOID(hwnd)), C.HMENU(C.PVOID(hmenu))) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func GetMenu(hwnd native.Handle) native.Handle {
	return native.Handle(C.PVOID(C.GetMenu(C.HWND(C.PVOID(hwnd)))))
}

func DrawMenuBar(hwnd native.Handle) bool {
	return C.DrawMenuBar(C.HWND(C.PVOID(hwnd))) != 0
}

func UnhookWindowsHookEx(hook native.Handle) {
	if C.UnhookWindowsHookEx(C.HHOOK(C.PVOID(hook))) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func TranslateAccelerator(hwnd, accel native.Handle, msg PMsg) bool {
	return C.TranslateAccelerator(C.HWND(C.PVOID(hwnd)), C.HACCEL(C.PVOID(accel)), (*C.MSG)(unsafe.Pointer(msg))) != 0
}

const (
	COLOR_WINDOW = int(C.COLOR_WINDOW)
)

func GetSysColor(index int) uint32 {
	return uint32(C.GetSysColor(C.int(index)))
}

func BringWindowToTop(handle native.Handle) bool {
	return C.BringWindowToTop(C.HWND(C.PVOID(handle))) != 0
}

type WndClassEx struct {
	Style      uint
	WndProc    unsafe.Pointer
	ClsExtra   int
	WndExtra   int
	Instance   native.Handle
	Icon       native.Handle
	Cursor     native.Handle
	Background native.Handle
	MenuName   unsafe.Pointer
	ClassName  unsafe.Pointer
	IconSm     native.Handle
}

func newWNDCLASSEX(cls *WndClassEx) *C.WNDCLASSEX {
	p := (*C.WNDCLASSEX)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(C.WNDCLASSEX{}))[0]))
	p.cbSize = C.UINT(unsafe.Sizeof(C.WNDCLASSEX{}))
	p.style = C.UINT(cls.Style)
	p.lpfnWndProc = C.WNDPROC(cls.WndProc)
	p.cbClsExtra = C.int(cls.ClsExtra)
	p.cbWndExtra = C.int(cls.WndExtra)
	p.hInstance = C.HINSTANCE(C.PVOID(cls.Instance))
	p.hIcon = C.HICON(C.PVOID(cls.Icon))
	p.hCursor = C.HCURSOR(C.PVOID(cls.Cursor))
	p.hbrBackground = C.HBRUSH(C.PVOID(cls.Background))
	p.lpszMenuName = C.LPCWSTR(cls.MenuName)
	p.lpszClassName = C.LPCWSTR(cls.ClassName)
	p.hIconSm = C.HICON(C.PVOID(cls.IconSm))
	return p
}

func RegisterClassEx(class *WndClassEx) {
	if C.RegisterClassEx(newWNDCLASSEX(class)) == 0 {
		nativeutil.PanicWithLastError()
	}
}

func DefWindowProcPtr() unsafe.Pointer {
	return unsafe.Pointer(C.GetDefWindowProcPtr())
}

var (
	IDC_ARROW = uintptr(C.PVOID(C.VarIDC_ARROW))
)

func LoadCursor(instance native.Handle, name uintptr) native.Handle {
	result := native.Handle(C.PVOID(C.LoadCursor(C.HINSTANCE(C.PVOID(instance)), C.LPCWSTR(unsafe.Pointer(name)))))
	if result == 0 {
		nativeutil.PanicWithLastError()
	}
	return result
}

func GetActiveWindow() native.Handle {
	return native.Handle(C.PVOID(C.GetActiveWindow()))
}

type PMsg uintptr

// Must deallocated after use with PMsg.Free.
func AllocMsg() PMsg {
	return PMsg(C.malloc(C.size_t(unsafe.Sizeof(C.MSG{}))))
}

func (msg PMsg) pMsg() C.PMSG {
	return C.PMSG(unsafe.Pointer(msg))
}

func (msg PMsg) Free() {
	C.free(unsafe.Pointer(msg))
}

func (msg PMsg) Hwnd() native.Handle {
	return native.Handle(C.PVOID(msg.pMsg().hwnd))
}

func (msg PMsg) SetHwnd(handle native.Handle) {
	msg.pMsg().hwnd = C.HWND(C.PVOID(handle))
}

func (msg PMsg) Message() uint {
	return uint(msg.pMsg().message)
}

func (msg PMsg) SetMessage(message uint) {
	msg.pMsg().message = C.UINT(message)
}

func (msg PMsg) WParam() uintptr {
	return uintptr(msg.pMsg().wParam)
}

func (msg PMsg) SetWParam(wParam uintptr) {
	msg.pMsg().wParam = C.WPARAM(wParam)
}

func (msg PMsg) LParam() uintptr {
	return uintptr(msg.pMsg().lParam)
}

func (msg PMsg) SetLParam(lParam uintptr) {
	msg.pMsg().lParam = C.LPARAM(lParam)
}

func (msg PMsg) Time() uint {
	return uint(msg.pMsg().time)
}

func (msg PMsg) SetTime(time uint) {
	msg.pMsg().time = C.DWORD(time)
}

func (msg PMsg) Point() (x, y int) {
	return int(msg.pMsg().pt.x), int(msg.pMsg().pt.y)
}

func (msg PMsg) SetPoint(x, y int) {
	msg.pMsg().pt.x, msg.pMsg().pt.y = C.LONG(x), C.LONG(y)
}

func GetMessage(msg PMsg, hwnd native.Handle, msgMin, msgMax uint) bool {
	ret := C.GetMessage(msg.pMsg(), C.HWND(C.PVOID(hwnd)), C.UINT(msgMin), C.UINT(msgMax))
	if ret == -1 {
		nativeutil.PanicWithLastError()
	}
	return ret != 0
}

func IsDialogMessage(hwnd native.Handle, msg PMsg) bool {
	return C.IsDialogMessage(C.HWND(C.PVOID(hwnd)), msg.pMsg()) != 0
}

func TranslateMessage(msg PMsg) bool {
	return C.TranslateMessage(msg.pMsg()) != 0
}

func DispatchMessage(msg PMsg) bool {
	return C.DispatchMessage(msg.pMsg()) != 0
}
