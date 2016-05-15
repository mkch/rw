package icon

// https://msdn.microsoft.com/en-us/library/ms997538.aspx
// https://en.wikipedia.org/wiki/ICO_(file_format)

import (
	"fmt"
	"io"
	"io/ioutil"
	"unsafe"
)

type IconDirHeader [6]byte

// Reserved must be 0.
func (hdr *IconDirHeader) Reserved() *uint16 {
	return (*uint16)(unsafe.Pointer(&hdr[0]))
}

// 1 for icon. 2 for cursor.
func (hdr *IconDirHeader) Type() *uint16 {
	return (*uint16)(unsafe.Pointer(&hdr[2]))
}

func (hdr *IconDirHeader) Count() *uint16 {
	return (*uint16)(unsafe.Pointer(&hdr[4]))
}

type IconDirEntry [16]byte

func (entry *IconDirEntry) Width() *uint8 {
	return &entry[0]
}

func (entry *IconDirEntry) Height() *uint8 {
	return &entry[1]
}

func (entry *IconDirEntry) ColorCount() *uint8 {
	return &entry[2]
}

// Reserved must be 0.
func (entry *IconDirEntry) Reserved() *uint8 {
	return &entry[3]
}

// Planes returns the image planes of icon(type == 1).
func (entry *IconDirEntry) Planes() *uint16 {
	return (*uint16)(unsafe.Pointer(&entry[4]))
}

// Hotspot returns the hotspot of cursor(type == 2).
func (entry *IconDirEntry) Hotspot() (x, y *uint8) {
	return &entry[4], &entry[5]
}

func (entry *IconDirEntry) BitCount() *uint16 {
	return (*uint16)(unsafe.Pointer(&entry[6]))
}

func (entry *IconDirEntry) BytesInRes() *uint32 {
	return (*uint32)(unsafe.Pointer(&entry[8]))
}

func (entry *IconDirEntry) ImageOffset() *uint32 {
	return (*uint32)(unsafe.Pointer(&entry[12]))
}

// RT_GROUP_ICON resource dir entry.
type GrpIconDirEntry [14]byte

func (entry *GrpIconDirEntry) Width() *uint8 {
	return &entry[0]
}

func (entry *GrpIconDirEntry) Height() *uint8 {
	return &entry[1]
}

func (entry *GrpIconDirEntry) ColorCount() *uint8 {
	return &entry[2]
}

// Reserved must be 0.
func (entry *GrpIconDirEntry) Reserved() *uint8 {
	return &entry[3]
}

// Planes returns the image planes of icon(type == 1).
func (entry *GrpIconDirEntry) Planes() *uint16 {
	return (*uint16)(unsafe.Pointer(&entry[4]))
}

// Hotspot returns the hotspot of cursor(type == 2).
func (entry *GrpIconDirEntry) Hotspot() (x, y *uint8) {
	return &entry[4], &entry[5]
}

func (entry *GrpIconDirEntry) BitCount() *uint16 {
	return (*uint16)(unsafe.Pointer(&entry[6]))
}

func (entry *GrpIconDirEntry) BytesInRes() *uint32 {
	return (*uint32)(unsafe.Pointer(&entry[8]))
}

func (entry *GrpIconDirEntry) ID() *uint16 {
	return (*uint16)(unsafe.Pointer(&entry[12]))
}

type Image struct {
	Entry IconDirEntry
	Data  []byte
}

type Icon struct {
	Type   uint16 // See IconDirHeader.Type()
	Images []Image
}

type FileFormatError struct {
	err string
}

func (e *FileFormatError) Error() string {
	return e.err
}

func Read(r io.Reader) (result *Icon, err error) {
	var n, bytesBeforeData int
	// Read ICONDIR
	var header IconDirHeader
	if n, err = io.ReadFull(r, header[:]); err != nil {
		return
	}
	bytesBeforeData += n

	if *header.Reserved() != 0 {
		err = &FileFormatError{"Wrong file header"}
		return
	}
	if *header.Type() != 1 && *header.Type() != 2 {
		err = &FileFormatError{fmt.Sprintf("Wrong file type: %v", *header.Type())}
		return
	}

	result = &Icon{Type: *header.Type(), Images: make([]Image, *header.Count())}

	// Read []IconDirEntry
	for i := uint16(0); i < *header.Count(); i++ {
		if n, err = io.ReadFull(r, result.Images[i].Entry[:]); err != nil {
			return
		}
		bytesBeforeData += n
	}

	// Read the remain data
	var data []byte
	if data, err = ioutil.ReadAll(r); err != nil {
		return
	}

	for i := uint16(0); i < *header.Count(); i++ {
		start := *result.Images[i].Entry.ImageOffset() - uint32(bytesBeforeData)
		size := *result.Images[i].Entry.BytesInRes()
		result.Images[i].Data = data[start : start+size]
	}

	return
}
