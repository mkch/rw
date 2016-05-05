package bitmap_test

import (
	"github.com/mkch/rw/util/bitmap"
	"os"
	"testing"
)

func TestMono(t *testing.T) {
	file, err := os.Open("test_data/mono.bmp")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	bmp, err := bitmap.Read(file)
	if err != nil {
		t.Fatal(err)
	}
	if bmp == nil {
		t.Fatal("Read returned nil.")
	}
	if bmp.Width() != 100 || bmp.Height() != 100 {
		t.Fatalf("Wrong size: %vx%v. Should be 100x100", bmp.Width(), bmp.Height())
	}
	colorTable := bmp.ColorTable()
	if len(colorTable) != 2 || colorTable[0] != (bitmap.RgbQuad{0, 0, 0, 0}) || colorTable[1] != (bitmap.RgbQuad{0xFF, 0xFF, 0xFF, 0}) {
		t.Fatalf("Wrong color table for monochrome bitmap: %v\n", colorTable)
	}

	color := bmp.ColorAt(0, 0)
	if color != (bitmap.RGB{0, 0, 0}) {
		t.Errorf("Wrong color at (0,0): %v\n", color)
	}

	color = bmp.ColorAt(6, 7)
	if color != (bitmap.RGB{0xFF, 0xFF, 0xFF}) {
		t.Errorf("Wrong color at (6,7): %v\n", color)
	}
	color = bmp.ColorAt(8, 8)
	if color != (bitmap.RGB{0, 0, 0}) {
		t.Errorf("Wrong color at (8,8): %v\n", color)
	}
	color = bmp.ColorAt(50, 50)
	if color != (bitmap.RGB{0, 0, 0}) {
		t.Errorf("Wrong color at (50,50): %v\n", color)
	}
	color = bmp.ColorAt(50, 51)
	if color != (bitmap.RGB{0xFF, 0xFF, 0xFF}) {
		t.Errorf("Wrong color at (50,51): %v\n", color)
	}
	color = bmp.ColorAt(99, 99)
	if color != (bitmap.RGB{0, 0, 0}) {
		t.Errorf("Wrong color at (99,99): %v\n", color)
	}
}

func Test4(t *testing.T) {
	testColor(t, "16colors.bmp", 16)
}

func Test8(t *testing.T) {
	testColor(t, "256colors.bmp", 256)
}

// func Test16(t *testing.T) {
// 	testColor(t, "", -1)
// }

func Test24(t *testing.T) {
	testColor(t, "24bits.bmp", -1)
}

func testColor(t *testing.T, filename string, colorTableLen int) {
	file, err := os.Open("test_data/" + filename)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	bmp, err := bitmap.Read(file)
	if err != nil {
		t.Fatal(err)
	}
	if bmp == nil {
		t.Fatal("Read returned nil.")
	}
	if bmp.Width() != 100 || bmp.Height() != 100 {
		t.Fatalf("Wrong size: %vx%v. Should be 100x100", bmp.Width(), bmp.Height())
	}
	colorTable := bmp.ColorTable()
	if colorTableLen != -1 && len(colorTable) != colorTableLen {
		t.Fatalf("Wrong color table for bitmap\n")
	}

	color := bmp.ColorAt(0, 0)
	if color != (bitmap.RGB{0xFF, 0, 0}) {
		t.Errorf("Wrong color at (0,0): %v\n", color)
	}

	color = bmp.ColorAt(6, 7)
	if color != (bitmap.RGB{0xFF, 0xFF, 0xFF}) {
		t.Errorf("Wrong color at (6,7): %v\n", color)
	}
	color = bmp.ColorAt(8, 8)
	if color != (bitmap.RGB{0xFF, 0, 0}) {
		t.Errorf("Wrong color at (8,8): %v\n", color)
	}
	color = bmp.ColorAt(50, 50)
	if color != (bitmap.RGB{0, 0, 0}) {
		t.Errorf("Wrong color at (50,50): %v\n", color)
	}
	color = bmp.ColorAt(50, 51)
	if color != (bitmap.RGB{0xFF, 0xFF, 0xFF}) {
		t.Errorf("Wrong color at (50,51): %v\n", color)
	}
	color = bmp.ColorAt(99, 99)
	if color != (bitmap.RGB{0, 0, 0xFF}) {
		t.Errorf("Wrong color at (99,99): %v\n", color)
	}
}
