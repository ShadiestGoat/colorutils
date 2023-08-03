package colorutils

import (
	"fmt"
	"math/rand"
	"testing"
)

const TEST_AMT = 100_000_000

func TestContrastColorDarkBg(t *testing.T) {
	for i := 0; i < TEST_AMT; i++ {
		bgR, bgG, bgB := HSLToRGB(rand.Intn(360), rand.Float64()*0.7, 0.15)
	
		r, g, b := NewContrastColorDarkBg(bgR, bgG, bgB)
	
		c := ContrastRatio(
			RelativeLuminosity(bgR, bgG, bgB),
			RelativeLuminosity(r, g, b),
		)
	
		if i % 100_000 == 0 {
			t.Logf("%f%% tries passed <3", 100 * float64(i)/float64(TEST_AMT))
		}

		if c < 7 {
			fmt.Println(bgR, bgB, bgG)
			fmt.Println(r, g, b)

			t.FailNow()
		}
	}
}

func TestContrastColorLightBg(t *testing.T) {
	for i := 0; i < TEST_AMT; i++ {
		bgR, bgG, bgB := HSLToRGB(rand.Intn(360), rand.Float64()*0.7, 0.15)
	
		r, g, b := NewContrastColorLightBg(bgR, bgG, bgB)
	
		c := ContrastRatio(
			RelativeLuminosity(bgR, bgG, bgB),
			RelativeLuminosity(r, g, b),
		)
	
		if i % 100_000 == 0 {
			t.Logf("%f%% tries passed <3", 100 * float64(i)/float64(TEST_AMT))
		}

		if c < 7 {
			fmt.Println(bgR, bgB, bgG)
			fmt.Println(r, g, b)

			t.FailNow()
		}
	}
}