package icon_test

import (
	"github.com/mkch/rw/util/icon"
	"os"
	"testing"
)

func Test(t *testing.T) {
	var f *os.File
	var err error
	if f, err = os.Open("test_data/ico128.ico"); err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var ico *icon.Icon
	if ico, err = icon.Read(f); err != nil {
		t.Fatal(err)
	}
	t.Logf("type=%d len=%d\n", ico.Type, len(ico.Images))
	var totalBytes uint32
	for i, Image := range ico.Images {
		t.Logf("%v) %vx%v\t%vbpp\t%v bytes\n", i, *Image.Entry.Width(), *Image.Entry.Height(), *Image.Entry.BitCount(), len(Image.Data))
		totalBytes += uint32(len(Image.Data))
	}
	t.Logf("Total data size: %v\n", totalBytes)
}
