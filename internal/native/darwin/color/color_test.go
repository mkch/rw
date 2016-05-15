package color

import (
	"image/color"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	var c1 = color.RGBA{101, 202, 255, 14}
	var c2 = NSColor_getRGBA(NSColor_colorWithRGB(c1))
	if c1 != c2 {
		t.Errorf("RoundTrip failed: %v vs. %v\n", c1, c2)
	}
}
