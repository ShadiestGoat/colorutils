package colorutils

import (
	"math/rand"
	"testing"
)

const TEST_AMT = 1_000_000
const TOO_MANY = 10_000

func contrastFactory(t *testing.T, lightness float64, title string, f func (uint8, uint8, uint8) (uint8, uint8, uint8)) {
	hexToAmt := map[string]int{}

	t.Parallel()

	for i := 0; i < TEST_AMT; i++ {
		bgR, bgG, bgB := HSLToRGB(rand.Intn(360), rand.Float64()*0.7, lightness)
	
		r, g, b := f(bgR, bgG, bgB)
	
		c := ContrastRatio(
			RelativeLuminosity(bgR, bgG, bgB),
			RelativeLuminosity(r, g, b),
		)
	
		if i % 100_000 == 0 {
			t.Logf("%f%% tries passed <3", 100 * float64(i)/float64(TEST_AMT))
		}

		hex := Hexadecimal(r, g, b)
		hexToAmt[hex]++
		
		if hexToAmt[hex] > TOO_MANY {
			t.Errorf("%s generated too many of %s\nEither retry the test, or something is wrong\n", title, hex)
			t.FailNow()
		}
		
		if c < 7 {
			t.Errorf("%s generated a contrast ratio < 7:\nBG: %v %v %v\nOut: %v %v %v\n", title, bgR, bgB, bgG, r, g, b)
			t.FailNow()
		}
	}
}

func TestContrastColorDarkBg(t *testing.T) {
	contrastFactory(t, 0.15, "Dark BG", NewContrastColorDarkBg)
}

func TestContrastColorLightBg(t *testing.T) {
	contrastFactory(t, 0.85, "Light BG", NewContrastColorLightBg)
}