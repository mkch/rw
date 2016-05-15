package main

/*
#cgo CFLAGS: -D UNICODE
#include "makegui_windows.h"
*/
import "C"

import (
	"flag"
	"fmt"
	"github.com/mkch/rw/util/icon"
	"io"
	"os"
	"unicode/utf16"
	"unsafe"
)

func main() {
	var icon string
	flag.StringVar(&icon, "icon", "", "The icon file(*.ico) used as the icon of the exe.")
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Make a console program a UI program and optionally set it's icon.")
		flag.PrintDefaults()
		os.Exit(-1)
	}

	var file = flag.Arg(0)
	if makegui(file) && (len(icon) == 0 || changeIcon(file, icon)) {
		os.Exit(0)
	}
	os.Exit(-2)
}

func makegui(path string) bool {
	fmt.Fprintln(os.Stdin, path)
	f, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		printError(err)
		return false
	}
	defer f.Close()

	// Read IMAGE_DOS_HEADER.e_magic
	var dosHdrMagic C.WORD
	if _, err = io.ReadFull(f, (*[unsafe.Sizeof(dosHdrMagic)]byte)(unsafe.Pointer(&dosHdrMagic))[:]); err != nil {
		printError(err)
		return false
	}
	if dosHdrMagic != C.IMAGE_DOS_SIGNATURE {
		fmt.Fprintf(os.Stderr, "Invalid file format: Wrong DOS header magic.\n")
		return false
	}
	// Read IMAGE_DOS_HEADER.e_lfanew
	if _, err = f.Seek(int64(unsafe.Offsetof(C.PIMAGE_DOS_HEADER(nil).e_lfanew)), 0); err != nil {
		printError(err)
		return false
	}
	var fileOffsetToNtHdrs C.LONG
	if _, err = io.ReadFull(f, (*[unsafe.Sizeof(fileOffsetToNtHdrs)]byte)(unsafe.Pointer(&fileOffsetToNtHdrs))[:]); err != nil {
		printError(err)
		return false
	}
	// Seek to IMAGE_NT_HEADERS
	if _, err = f.Seek(int64(fileOffsetToNtHdrs), 0); err != nil {
		printError(err)
		return false
	}
	// Read IMAGE_NT_HEADERS.Signature
	var ntSig C.DWORD
	if _, err = io.ReadFull(f, (*[unsafe.Sizeof(ntSig)]byte)(unsafe.Pointer(&ntSig))[:]); err != nil {
		printError(err)
		return false
	}
	if ntSig != C.IMAGE_NT_SIGNATURE {
		fmt.Fprintf(os.Stderr, "Invalid file format: Wrong NT signature.\n")
	}
	// Skip IMAGE_FILE_HEADER
	if _, err = f.Seek(int64(unsafe.Sizeof(*(C.PIMAGE_FILE_HEADER)(nil))), 1); err != nil {
		printError(err)
		return false
	}
	// Read IMAGE_OPTIONAL_HEADER.Magic
	var optHdrMagic C.WORD
	if _, err = io.ReadFull(f, (*[unsafe.Sizeof(optHdrMagic)]byte)(unsafe.Pointer(&optHdrMagic))[:]); err != nil {
		printError(err)
		return false
	}
	// The offset from IMAGE_OPTIONAL_HEADER to IMAGE_OPTIONAL_HEADER.Subsystem.
	var subsystemOffset int64
	switch optHdrMagic {
	case C.IMAGE_NT_OPTIONAL_HDR32_MAGIC:
		subsystemOffset = int64(unsafe.Offsetof(C.PIMAGE_OPTIONAL_HEADER32(nil).Subsystem))
	case C.IMAGE_NT_OPTIONAL_HDR64_MAGIC:
		subsystemOffset = int64(unsafe.Offsetof(C.PIMAGE_OPTIONAL_HEADER64(nil).Subsystem))
	default:
		fmt.Fprintf(os.Stderr, "Wrong optional header magic.\n")
	}
	// Seek to IMAGE_OPTIONAL_HEADER.Subsystem
	var fileOffsetToSubsystem int64
	if fileOffsetToSubsystem, err = f.Seek(subsystemOffset-int64(unsafe.Sizeof(optHdrMagic)), 1); err != nil {
		printError(err)
		return false
	}
	// Read IMAGE_OPTIONAL_HEADER.Subsystem
	var subsystem C.WORD
	if _, err = io.ReadFull(f, (*[unsafe.Sizeof(subsystem)]byte)(unsafe.Pointer(&subsystem))[:]); err != nil {
		printError(err)
		return false
	}
	switch subsystem {
	case C.IMAGE_SUBSYSTEM_WINDOWS_GUI:
		fmt.Fprintln(os.Stdout, "Already a GUI program.")
	case C.IMAGE_SUBSYSTEM_WINDOWS_CUI:
	default:
		fmt.Fprintf(os.Stderr, "Unsupported subsystem 0x%X.\n", subsystem)
	}

	// Seek back to IMAGE_OPTIONAL_HEADER.Subsystem
	if _, err = f.Seek(fileOffsetToSubsystem, 0); err != nil {
		printError(err)
		return false
	}
	// Overwrite IMAGE_OPTIONAL_HEADER.Subsystem to C.IMAGE_SUBSYSTEM_WINDOWS_GUI
	subsystem = C.IMAGE_SUBSYSTEM_WINDOWS_GUI
	if _, err = f.Write((*[unsafe.Sizeof(subsystem)]byte)(unsafe.Pointer(&subsystem))[:]); err != nil {
		printError(err)
		return false
	}

	return true
}

//export go_enumResourceNamesCallback
func go_enumResourceNamesCallback(name C.LPCTSTR, param unsafe.Pointer) {
	names := (*[]*C.CHAR)(param)
	*names = append(*names, (*C.CHAR)(unsafe.Pointer(name)))
}

// firstGroupIconResource returns the first RT_GROUP_ICON resource(ico in resource format) ID
// and the RT_ICON ids in the group, in the module file(exe or dll).
// path is the UTF-16 format file path to the module.
func firstGroupIconResource(path *C.WCHAR) (ok bool, groupId *C.WCHAR, iconIds []*C.WCHAR) {
	h := C.LoadLibraryEx(path, nil, C._LOAD_LIBRARY_AS_DATAFILE)
	if h == nil {
		printLastError("LoadLibraryEx")
		return
	}
	defer C.FreeLibrary(h)

	var groupIconIds []*C.CHAR
	C.EnumResourceNames_(h, C._RT_GROUP_ICON, unsafe.Pointer(&groupIconIds))
	if len(groupIconIds) == 0 { // No icon at all.
		ok = true
		return
	}
	// The first icon is the one displays.
	groupId = (*C.WCHAR)(unsafe.Pointer(groupIconIds[0]))
	hRes := C.FindResource(h, groupId, (*C.WCHAR)(unsafe.Pointer(C._RT_GROUP_ICON)))
	if hRes == nil {
		printLastError("FindResource")
		return
	}
	hGlobal := C.LoadResource(h, hRes)
	if hGlobal == nil {
		printLastError("LoadResource")
		return
	}
	pRes := C.LockResource(hGlobal)
	if pRes == nil {
		printLastError("LockResource")
		return
	}
	nRes := C.SizeofResource(h, hRes)
	if nRes == 0 {
		printLastError("SizeofResource")
		return
	}

	grpHdr := (*icon.IconDirHeader)(unsafe.Pointer(uintptr(pRes)))
	iconIds = make([]*C.WCHAR, *grpHdr.Count())
	// N icon.GrpIconDirEntry(s) follow the icon.IconDirHeader
	for i := 0; i < len(iconIds); i++ {
		entry := (*icon.GrpIconDirEntry)(unsafe.Pointer(uintptr(pRes) + unsafe.Sizeof(*grpHdr) + unsafe.Sizeof(icon.GrpIconDirEntry{})*uintptr(i)))
		//fmt.Println(entry)
		iconIds[i] = C._MAKEINTRESOURCE(C.WORD(*entry.ID()))
	}
	ok = true
	return
}

func changeIcon(exe, ico string) bool {
	exe16 := toUtf16(exe)
	ok, grp, icons := firstGroupIconResource(exe16)
	if !ok {
		return false
	} else if grp != nil {
		fmt.Fprintf(os.Stdout, "Already has icon. 0x%X %X\n", grp, icons)
	}

	hUpdate := C.BeginUpdateResource(exe16, C.FALSE)
	if hUpdate == nil {
		printLastError("BeginUpdateResource")
		return false
	}
	var discardChange C.BOOL = C.TRUE
	defer func() {
		C.EndUpdateResource(hUpdate, discardChange)
	}()
	// Load the ico file
	f, err := os.Open(ico)
	if err != nil {
		printError(err)
		return false
	}
	defer f.Close()
	var iconData *icon.Icon
	if iconData, err = icon.Read(f); err != nil {
		printError(err)
		return false
	}
	// Remove all the icons in group.
	for _, iconResId := range icons {
		if C.UpdateResource(hUpdate, (*C.WCHAR)(unsafe.Pointer(C._RT_ICON)), iconResId, C.LANG_NUTRAL, nil, 0) == C.FALSE {
			printLastError(fmt.Sprintf("Delete icon resource 0x%x", iconResId))
			return false
		}
	}
	var addedIconResIds []uint16
	// Add new icon
	hModule := C.LoadLibraryEx(exe16, nil, C._LOAD_LIBRARY_AS_DATAFILE)
	if hUpdate == nil {
		printLastError("LoadLibraryEx")
		return false
	}
	defer C.FreeLibrary(hModule)
	var lastFreeIconId uint16 = 1
	for i := 0; i < len(iconData.Images); i++ {
		var iconID uint16
		if len(icons) > 0 {
			iconID = uint16(uintptr(unsafe.Pointer(icons[0])))
			icons = icons[1:]
		}
		if iconID == 0 {
			// Find a free ID
			// https://msdn.microsoft.com/en-us/library/t2zechd4.aspx
			for nextId := lastFreeIconId; nextId < 0x6FFF; nextId++ {
				if C.FindResource(hModule, C._MAKEINTRESOURCE(C.WORD(nextId)), (*C.WCHAR)(unsafe.Pointer(C._RT_ICON))) == nil {
					iconID = nextId
					lastFreeIconId = nextId + 1
					break
				}
			}
		}
		if iconID == 0 {
			fmt.Fprintln(os.Stderr, "Run out of icon ID.")
			return false
		}
		addedIconResIds = append(addedIconResIds, iconID)
		if C.UpdateResource(hUpdate, (*C.WCHAR)(unsafe.Pointer(C._RT_ICON)), C._MAKEINTRESOURCE(C.WORD(iconID)), C.LANG_NUTRAL, unsafe.Pointer(&iconData.Images[i].Data[0]), C.DWORD(len(iconData.Images[i].Data))) == C.FALSE {
			printLastError(fmt.Sprintf("UpdateResource - Add icon resource 0x%x", iconID))
			return false
		}

	}
	if grp != nil {
		// Remove group icon if needed. Reuse the ID.
		if C.UpdateResource(hUpdate, (*C.WCHAR)(unsafe.Pointer(C._RT_GROUP_ICON)), (*C.WCHAR)(grp), C.LANG_NUTRAL, nil, 0) == C.FALSE {
			printLastError("UpdateResource - Delete group icon resource")
			return false
		}
	} else {
		// Find a free group icon resource ID.
		for i := uint16(1); i < 0x6FFF; i++ {
			if C.FindResource(hModule, C._MAKEINTRESOURCE(C.WORD(i)), (*C.WCHAR)(unsafe.Pointer(C._RT_GROUP_ICON))) == nil {
				grp = C._MAKEINTRESOURCE(C.WORD(i))
				break
			}
		}
		if grp == nil {
			fmt.Fprintln(os.Stderr, "Run out of group icon ID.")
			return false
		}
	}
	// Add/replace group icon.
	groupIconData := make([]byte, unsafe.Sizeof(icon.IconDirHeader{})+uintptr(len(iconData.Images))*unsafe.Sizeof(icon.GrpIconDirEntry{}))
	hdr := (*icon.IconDirHeader)(unsafe.Pointer(&groupIconData[0]))
	*hdr.Type() = iconData.Type
	*hdr.Count() = uint16(len(iconData.Images))
	for i := uint16(0); i < uint16(len(iconData.Images)); i++ {
		entry := (*icon.GrpIconDirEntry)(unsafe.Pointer(uintptr(unsafe.Pointer(hdr)) + unsafe.Sizeof(icon.IconDirHeader{}) + uintptr(i)*unsafe.Sizeof(icon.GrpIconDirEntry{})))
		*entry.Width() = *iconData.Images[i].Entry.Width()
		*entry.Height() = *iconData.Images[i].Entry.Height()
		*entry.ColorCount() = *iconData.Images[i].Entry.ColorCount()
		*entry.Planes() = *iconData.Images[i].Entry.Planes()
		*entry.BitCount() = *iconData.Images[i].Entry.BitCount()
		*entry.BytesInRes() = *iconData.Images[i].Entry.BytesInRes()
		*entry.ID() = addedIconResIds[i]
	}
	if C.UpdateResource(hUpdate, (*C.WCHAR)(unsafe.Pointer(C._RT_GROUP_ICON)), grp, C.LANG_NUTRAL, unsafe.Pointer(&groupIconData[0]), C.DWORD(len(groupIconData))) == C.FALSE {
		printLastError("UpdateResource - Add group icon resource")
		return false
	}
	discardChange = C.FALSE
	return true
}

func printError(err error) {
	fmt.Fprintln(os.Stderr, err)
}

// toUtf16 converts go string(UTF-8) to MS Windows unicode string(UTF-16).
func toUtf16(s string) *C.WCHAR {
	s16 := append(utf16.Encode([]rune(s)), 0)
	return (*C.WCHAR)(unsafe.Pointer(&s16[0]))
}

func fromUtf16(s *C.WCHAR) string {
	var count uint
	for p := s; *p != 0; p = (*C.WCHAR)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(C.WCHAR(0)))) {
		count++
	}
	return string(utf16.Decode((*[^uintptr(0) / 2 / unsafe.Sizeof(uint16(0))]uint16)(unsafe.Pointer(s))[:count]))
}

func getLastErrorMessage(lastError C.DWORD) (errorMessage string) {
	var lastErrorMessageBuffer C.LPWSTR
	if C.FormatMessage(C.FORMAT_MESSAGE_ALLOCATE_BUFFER|C.FORMAT_MESSAGE_FROM_SYSTEM, nil, lastError, 0, C.LPWSTR(unsafe.Pointer(&lastErrorMessageBuffer)), 0, nil) != 0 && lastErrorMessageBuffer != nil {
		defer C.LocalFree(C.HLOCAL(lastErrorMessageBuffer))
		return fromUtf16(lastErrorMessageBuffer)
	}
	panic("FormatMessage failed!")
}

func printLastError(name string) {
	err := C.GetLastError()
	fmt.Fprintf(os.Stderr, "%s failed. 0x%X: %v\n", name, err, getLastErrorMessage(err))
}
