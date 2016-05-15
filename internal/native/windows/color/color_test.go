package color

import (
	"image/color"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	var c1 = color.RGBA{101, 202, 255, 255}
	var c2 = Color_getColor(Color_color(c1))
	if c1 != c2 {
		t.Errorf("RoundTrip failed: %v vs. %v\n", c1, c2)
	}
}
