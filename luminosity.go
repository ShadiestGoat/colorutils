package colorutils

import "math/rand"

// A factor for the luminosity calculation
const (
	LUMINOSITY_FACTOR_R = 0.2126
	LUMINOSITY_FACTOR_G = 0.7152
	LUMINOSITY_FACTOR_B = 0.0722
)

// returns the contrast ratio between L1 and L2, where L1 ad L2 are relative luminosities of 2 colors
func ContrastRatio(L1, L2 float64) float64 {
	if L1 < L2 {
		L2, L1 = L1, L2
	}

	return (L1 + 0.05) / (L2 + 0.05)
}

// Calculate the relative luminosity of an rgb value
func RelativeLuminosity(r, g, b uint8) float64 {
	rv, gv, bv := srgbToV(uint8ToSRGB(r)), srgbToV(uint8ToSRGB(g)), srgbToV(uint8ToSRGB(b))
	return LUMINOSITY_FACTOR_R*rv + LUMINOSITY_FACTOR_G*gv + LUMINOSITY_FACTOR_B*bv
}

// This creates a new WCAG AAA contrasting color to a background with the following assumptions:
// - this color is used for text in a heading (ie. this makes a > 7 ratio, closer to 7.5)
// - the background is darker than the text
// 
// The background is given by the RGB representation of it - bgR, bgG, bgB
func NewContrastColor(bgR, bgG, bgB uint8) (r, g, b uint8) {
	// 7.5 < (cY + 0.05)/(bgY + 0.05)
	// 7.5*(bgY + 0.05) - 0.05 = cY

	lhs := 7.5*(RelativeLuminosity(bgR, bgG, bgB)+0.05) - 0.05

	// we must now basically get the upper and lower bounds for a rand func, repeatedly
	// lets do in r, g, b for now.
	// a maximum value is the assumption that the rest are at their minimums, which for now we can assume to be 0s
	// a minimum value is teh assumption that the rest sRGB values are 1s

	var VMap = map[rune]float64{
		'r': 0,
		'g': 0,
		'b': 0,
	}

	var FMap = map[rune]float64{
		'r': LUMINOSITY_FACTOR_R,
		'g': LUMINOSITY_FACTOR_G,
		'b': LUMINOSITY_FACTOR_B,
	}

	pos := []rune{'r', 'g', 'b'}

	rand.Shuffle(len(pos), func(i, j int) {
		pos[i], pos[j] = pos[j], pos[i]
	})

	for i, r := range pos {
		vMAX := solveMax(lhs, FMap[r])

		factors := []float64{}
		for _, f := range pos[i+1:] {
			factors = append(factors, FMap[f])
		}

		vMIN := solveMin(lhs, FMap[r], factors...)
		V := randFloat(vMIN, vMAX)

		VMap[r] = V

		lhs -= srgbToV(V) * FMap[r]
	}

	r, g, b = srgbToUint8(VMap['r']), srgbToUint8(VMap['g']), srgbToUint8(VMap['b'])

	return
}
