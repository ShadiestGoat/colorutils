package colorutils

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewContrastColor(t *testing.T) {
	m := 100_000_000

	for i := 0; i < m; i++ {
		bgR, bgG, bgB := HSLToRGB(rand.Intn(360), rand.Float64()*0.7, 0.15)
	
		r, g, b := NewContrastColor(bgR, bgG, bgB)
	
		c := ContrastRatio(
			RelativeLuminosity(bgR, bgG, bgB),
			RelativeLuminosity(r, g, b),
		)
	
		if i % 100_000 == 0 {
			t.Logf("%f%% tries passed <3", 100 * float64(i)/float64(m))
		}

		if c < 7 {
			fmt.Println(bgR, bgB, bgG)
			fmt.Println(r, g, b)

			t.FailNow()
		}

	}
}