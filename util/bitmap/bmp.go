package bitmap

import (
	"fmt"
	"io"
	"io/ioutil"
	"unsafe"
)

type BitmapFileHeader [14]byte

func (hdr *BitmapFileHeader) Type() *uint16 {
	return (*uint16)(unsafe.Pointer(&hdr[0]))
}

func (hdr *BitmapFileHeader) Size() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[2]))
}

func (hdr *BitmapFileHeader) ReservedZero() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[6]))
}

func (hdr *BitmapFileHeader) OffBits() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[10]))
}

type BitmapInfoHeader [40]byte

func (hdr *BitmapInfoHeader) Size() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[0]))
}

func (hdr *BitmapInfoHeader) Width() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[4]))
}

func (hdr *BitmapInfoHeader) Height() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[8]))
}

func (hdr *BitmapInfoHeader) Planes() *uint16 {
	return (*uint16)(unsafe.Pointer(&hdr[12]))
}

func (hdr *BitmapInfoHeader) BitCount() *uint16 {
	return (*uint16)(unsafe.Pointer(&hdr[14]))
}

func (hdr *BitmapInfoHeader) Compression() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[16]))
}

func (hdr *BitmapInfoHeader) SizeImage() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[20]))
}

func (hdr *BitmapInfoHeader) XPelsPerMeter() *int32 {
	return (*int32)(unsafe.Pointer(&hdr[24]))
}

func (hdr *BitmapInfoHeader) YPelsPerMeter() *int32 {
	return (*int32)(unsafe.Pointer(&hdr[28]))
}

func (hdr *BitmapInfoHeader) ClrUsed() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[32]))
}

func (hdr *BitmapInfoHeader) ClrImportant() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[36]))
}

type BitmapV4Header [40 + 68]byte

func (hdr *BitmapV4Header) BitmapInfoHeader() *BitmapInfoHeader {
	return (*BitmapInfoHeader)(unsafe.Pointer(&hdr[0]))
}

func (hdr *BitmapV4Header) RedMask() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[40]))
}

func (hdr *BitmapV4Header) GreenMask() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[44]))
}

func (hdr *BitmapV4Header) BlueMask() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[48]))
}

func (hdr *BitmapV4Header) AlphaMask() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[52]))
}

func (hdr *BitmapV4Header) CSType() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[56]))
}

func (hdr *BitmapV4Header) Endpoints() *CieXYZTriple {
	return (*CieXYZTriple)(unsafe.Pointer(&hdr[60]))
}

func (hdr *BitmapV4Header) GammaRed() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[96]))
}

func (hdr *BitmapV4Header) GammaGreen() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[100]))
}

func (hdr *BitmapV4Header) GammaBlue() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[104]))
}

type BitmapV5Header [108 + 16]byte

func (hdr *BitmapV5Header) BitmapV4Header() *BitmapV4Header {
	return (*BitmapV4Header)(unsafe.Pointer(&hdr[0]))
}

func (hdr *BitmapV5Header) Intent() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[108]))
}

func (hdr *BitmapV5Header) ProfileData() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[112]))
}

func (hdr *BitmapV5Header) ProfileSize() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[116]))
}

func (hdr *BitmapV5Header) reserved() *uint32 {
	return (*uint32)(unsafe.Pointer(&hdr[120]))
}

type FxPt2Dot30 [4]byte

func (pt *FxPt2Dot30) Uint32() *uint32 {
	return (*uint32)(unsafe.Pointer(&pt[0]))
}

type CieXYZ [3 * 4]byte

func (xyz *CieXYZ) X() *FxPt2Dot30 {
	return (*FxPt2Dot30)(unsafe.Pointer(&xyz[0]))
}

func (xyz *CieXYZ) Y() *FxPt2Dot30 {
	return (*FxPt2Dot30)(unsafe.Pointer(&xyz[4]))
}

func (xyz *CieXYZ) Z() *FxPt2Dot30 {
	return (*FxPt2Dot30)(unsafe.Pointer(&xyz[8]))
}

type CieXYZTriple [3 * 3 * 4]byte

func (t *CieXYZTriple) Red() *CieXYZ {
	return (*CieXYZ)(unsafe.Pointer(&t[0]))
}

func (t *CieXYZTriple) Green() *CieXYZ {
	return (*CieXYZ)(unsafe.Pointer(&t[12]))
}

func (t *CieXYZTriple) Blue() *CieXYZ {
	return (*CieXYZ)(unsafe.Pointer(&t[24]))
}

type RgbQuad struct {
	Blue, Green, Red, Reserved byte
}

// ColorTable is the color table of bitmap.
type ColorTable []RgbQuad

type RGB struct {
	Red, Green, Blue byte
}

type FileFormatError struct {
	err string
}

func (e *FileFormatError) Error() string {
	return e.err
}

// countedReader is a reader that remembers how many bytes it has read.
type countedReader struct {
	R         io.Reader
	bytesRead int
}

func (r *countedReader) Read(p []byte) (n int, err error) {
	n, err = r.R.Read(p)
	r.bytesRead += n
	return
}

func (r *countedReader) BytesRead() int {
	return r.bytesRead
}

type Bitmap struct {
	Info   []byte // BITMAPINFO, info header + color table.
	Pixels []byte
}

// InfoHeaderSize returns the size of bitmap info header(BITMAPINFOHEADER, BITMAPV4HEADER or BITMAPV5HEADER).
func (bmp *Bitmap) InfoHeaderSize() uint32 {
	return *bmp.InfoHeader().Size()
}

// InfoHeader returns the BITMAPINFOHEADER of this bitmap.
// The color table data just follows the header, so the returned pointer can be used as pointer to BITMAPINFO also.
func (bmp *Bitmap) InfoHeader() *BitmapInfoHeader {
	return (*BitmapInfoHeader)(unsafe.Pointer(&bmp.Info[0]))
}

// V4Header returns the BITMAPV4HEADER of this bitmap,
// or nil if InfoHeaderSize() < unsafe.Sizeof(BitmapV4Header{}).
func (bmp *Bitmap) V4Header() *BitmapV4Header {
	if uintptr(bmp.InfoHeaderSize()) >= unsafe.Sizeof(BitmapV4Header{}) {
		return (*BitmapV4Header)(unsafe.Pointer(&bmp.Info[0]))
	}
	return nil
}

// V4Header returns the BITMAPV5HEADER of this bitmap
// or nil if InfoHeaderSize() < unsafe.Sizeof(BitmapV5Header{}).
func (bmp *Bitmap) V5Header() *BitmapV5Header {
	if uintptr(bmp.InfoHeaderSize()) >= unsafe.Sizeof(BitmapV5Header{}) {
		return (*BitmapV5Header)(unsafe.Pointer(&bmp.Info[0]))
	}
	return nil
}

func (bmp *Bitmap) Width() uint32 {
	return *bmp.InfoHeader().Width()
}

func (bmp *Bitmap) Height() uint32 {
	return *bmp.InfoHeader().Height()
}

// BitCount returns the number of bits that define each pixel and the maximum number of colors in the bitmap.
// 1, 4, 8, 16, 24, or 32.
func (bmp *Bitmap) BitCount() uint16 {
	return *bmp.InfoHeader().BitCount()
}

// HorizontalResolution returns the horizontal resolution, in pixels-per-meter, of the target device for the bitmap.
func (bmp *Bitmap) HorizontalResolution() int32 {
	return *bmp.InfoHeader().XPelsPerMeter()
}

// VerticalResolution returns the vertical resolution, in pixels-per-meter, of the target device for the bitmap.
func (bmp *Bitmap) VerticalResolution() int32 {
	return *bmp.InfoHeader().YPelsPerMeter()
}

// ColorTable returns the color table of bitmap or optimization information.
// The color table is returned if bit count is 1, 4, or 8.
// The optimization information is returned if bit count is 16, 24, or 32. See https://msdn.microsoft.com/en-us/library/windows/desktop/dd183376(v=vs.85).aspx
func (bmp *Bitmap) ColorTable() ColorTable {
	const rgbQuadSize = unsafe.Sizeof(RgbQuad{})
	infoHeaderSize := uintptr(bmp.InfoHeaderSize())
	colorTableLen := (uintptr(len(bmp.Info)) - infoHeaderSize) / rgbQuadSize
	return (*[_MAX_ADDR_SPACE / rgbQuadSize]RgbQuad)(unsafe.Pointer(uintptr(unsafe.Pointer(&bmp.Info[0])) + infoHeaderSize))[:colorTableLen]
}

func (bmp *Bitmap) ColorAt(x, y uint32) RGB {
	height := *bmp.InfoHeader().Height()
	width := *bmp.InfoHeader().Width()
	bitCount := *bmp.InfoHeader().BitCount()

	// Flips y.
	// BMP file has pixels data up side down.
	y = height - 1 - y
	bytesPerLine := multipleOf32(width*uint32(bitCount)) / 8
	lineOffset := bytesPerLine * y

	switch bitCount {
	case 1:
		// The byte coutains our bit.
		b := bmp.Pixels[lineOffset+x/8]
		i := (b >> (7 - x%8)) & byte(1)
		q := bmp.ColorTable()[i]
		return RGB{Red: q.Red, Green: q.Green, Blue: q.Blue}
	case 4:
		// The byte coutains our bit.
		b := bmp.Pixels[lineOffset+x/2]
		// s is the right shift bits.
		var s byte
		if x%2 == 0 {
			// We want the most significant 4 bits. Shift right 4 bits.
			s = 4
		}
		i := (b >> s) & byte(0xF)
		q := bmp.ColorTable()[i]
		return RGB{Red: q.Red, Green: q.Green, Blue: q.Blue}
	case 8:
		// The byte coutains our bit.
		b := bmp.Pixels[lineOffset+x]
		q := bmp.ColorTable()[b]
		return RGB{Red: q.Red, Green: q.Green, Blue: q.Blue}
	case 16:
		b1 := bmp.Pixels[lineOffset+x*2]
		b2 := bmp.Pixels[lineOffset+x*2+1]
		blue := b2 & byte(0x1F)
		green := ((b2 >> 5) & byte(0x7)) | ((b1 & 0x3) << 3)
		red := (b1 >> 2) & byte(0x1F)
		return RGB{Red: red, Green: green, Blue: blue}
	case 24:
		blue := bmp.Pixels[lineOffset+x*3]
		green := bmp.Pixels[lineOffset+x*3+1]
		red := bmp.Pixels[lineOffset+x*3+2]
		return RGB{Red: red, Green: green, Blue: blue}
	case 32:
		blue := bmp.Pixels[lineOffset+x*4]
		green := bmp.Pixels[lineOffset+x*4+1]
		red := bmp.Pixels[lineOffset+x*4+2]
		return RGB{Red: red, Green: green, Blue: blue}
	default:
		panic("Bit count not supported.")
	}
}

// ReadBitmapInfo reads the BITMAPINFO(BITMAPINFOHEADER+Color table).
func ReadBitmapInfo(r io.Reader) (info []byte, err error) {
	// Read struct size of next header.
	var infoHeaderSize uint32
	if infoHeaderSize, err = readInfoHeaderSize(r); err != nil {
		return
	}
	var infoHeader unsafe.Pointer
	switch uintptr(infoHeaderSize) {
	case unsafe.Sizeof(BitmapInfoHeader{}):
		var hdr BitmapInfoHeader
		err = readBitmapInfoHeaderWithoutSize(r, &hdr)
		infoHeader = unsafe.Pointer(&hdr)
	case unsafe.Sizeof(BitmapV4Header{}):
		var hdr BitmapV4Header
		err = readBitmapV4Header(r, &hdr)
		infoHeader = unsafe.Pointer(&hdr)
	case unsafe.Sizeof(BitmapV5Header{}):
		var hdr BitmapV5Header
		err = readBitmapV5Header(r, &hdr)
		infoHeader = unsafe.Pointer(&hdr)
	default:
		err = &FileFormatError{err: fmt.Sprintf("Unsupported info header size: 0x%X", infoHeaderSize)}
		return
	}
	// Fill the Size field of info header.
	// V4 or V5 header has BitmapInfoHeader in the front.
	// It is safe to cast the infoHeader to *BitmapInfoHeader.
	ihdr := (*BitmapInfoHeader)(infoHeader)
	*ihdr.Size() = infoHeaderSize

	// The number of planes for the target device. This value must be set to 1.
	// https://msdn.microsoft.com/en-us/library/windows/desktop/dd183376(v=vs.85).aspx
	planes := *ihdr.Planes()
	if planes != 1 {
		err = &FileFormatError{err: fmt.Sprintf("Unsupported count of planes %v. Only exact 1 plan supported.", planes)}
		return
	}
	bitCount := *ihdr.BitCount()
	if bitCount == 0 {
		err = &FileFormatError{err: fmt.Sprintf("JPEG or PNG compression is not supported.")}
		return
	}
	compression := *ihdr.Compression()
	if compression != 0 {
		err = &FileFormatError{err: fmt.Sprintf("Compressed BMP is not supported.")}
		return
	}
	sizeImage := *ihdr.SizeImage()
	if sizeImage == 0 {
		err = &FileFormatError{err: fmt.Sprintf("JPEG or PNG compression is not supported.")}
		return
	}

	// Calculate the entries count of color table.
	// https://msdn.microsoft.com/en-us/library/windows/desktop/dd183376(v=vs.85).aspx
	var colorTableEntryCount uint32
	switch bitCount {
	case 1:
		fallthrough
	case 4:
		fallthrough
	case 8:
		colorTableEntryCount = 2 << (bitCount - 1)
	default:
		// Not a real color table. optimization information.
		colorTableEntryCount = *ihdr.ClrUsed()
	}

	// Alloc BITMAPINFO
	info = make([]byte, uintptr(infoHeaderSize)+uintptr(colorTableEntryCount)*unsafe.Sizeof(RgbQuad{}))
	// copy the already read info header to the head of info.
	copy(info, (*[_MAX_ADDR_SPACE]byte)(infoHeader)[:uintptr(infoHeaderSize)])

	// Read color table if necessary.
	if colorTableEntryCount > 0 {
		colorTable := info[infoHeaderSize:]
		if _, err = io.ReadFull(r, colorTable); err != nil {
			return
		}
	}

	return
}

// Read reads a bitmap from a io.Reader.
func Read(reader io.Reader) (result *Bitmap, err error) {
	r := &countedReader{R: reader}
	// Read BITMAPFILEHEADER.
	var fileHeader *BitmapFileHeader
	if fileHeader, err = readFileHeader(r); err != nil {
		return
	} else if *fileHeader.Type() != 0x4D42 { // "BM"
		err = &FileFormatError{err: fmt.Sprintf("Wrong file type 0x%X", fileHeader.Type)}
		return
	}

	var info []byte
	if info, err = ReadBitmapInfo(r); err != nil {
		return
	}

	ihdr := (*BitmapInfoHeader)(unsafe.Pointer(&info[0]))
	// Calculate the size of pixel data.
	pixelsByteCount := PixelDataLen(*ihdr.Width(), *ihdr.Height(), *ihdr.BitCount())
	// Seek to pixel data.
	if _, err = io.CopyN(ioutil.Discard, r, int64(*fileHeader.OffBits()-uint32(r.BytesRead()))); err != nil {
		return
	}
	pixels := make([]byte, pixelsByteCount)
	if _, err = io.ReadFull(r, pixels); err != nil {
		return
	}

	result = &Bitmap{
		Info:   info,
		Pixels: pixels,
	}

	return
}

// PixelDataLen returns the pixel data length in byte.
func PixelDataLen(width, height uint32, bitCount uint16) uint32 {
	// Right shift 3 bits: divided by 8, aka bits count to bytes count.
	return (multipleOf32(width*uint32(bitCount)) >> 3) * height
}

func multipleOf32(v uint32) uint32 {
	// http://stackoverflow.com/questions/1766535/bit-hack-round-off-to-multiple-of-8
	// http://stackoverflow.com/questions/2022179/c-quick-calculation-of-next-multiple-of-4
	return (v + 31) &^ 31
}

// The max address space in byte.
// Not accurate.
const _MAX_ADDR_SPACE = uintptr(int(1) << (unsafe.Sizeof(int(0)) * 8 / 2))

func readFileHeader(r io.Reader) (header *BitmapFileHeader, err error) {
	var hdr BitmapFileHeader
	if _, err = io.ReadFull(r, hdr[:]); err != nil {
		return
	}
	header = &hdr
	return
}

func readInfoHeaderSize(r io.Reader) (size uint32, err error) {
	sizeAsSlice := (*[unsafe.Sizeof(size)]byte)(unsafe.Pointer(&size))[:]
	_, err = io.ReadFull(r, sizeAsSlice)
	return
}

func readBitmapInfoHeaderWithoutSize(r io.Reader, header *BitmapInfoHeader) (err error) {
	// The Size has been read by readInfoHeaderSize
	_, err = io.ReadFull(r, header[unsafe.Sizeof(*header.Size()):])
	return
}

func readBitmapV4Header(r io.Reader, header *BitmapV4Header) (err error) {
	if err = readBitmapInfoHeaderWithoutSize(r, header.BitmapInfoHeader()); err != nil {
		return
	}
	const remainOffset = unsafe.Sizeof(BitmapInfoHeader{})
	_, err = io.ReadFull(r, header[remainOffset:])
	return
}

func readBitmapV5Header(r io.Reader, header *BitmapV5Header) (err error) {
	if err = readBitmapV4Header(r, header.BitmapV4Header()); err != nil {
		return
	}
	const remainOffset = unsafe.Sizeof(BitmapV4Header{})
	_, err = io.ReadFull(r, header[remainOffset:])
	return
}
