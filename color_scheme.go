package colorutils

type HSV struct {
	Hue               int
	Saturation, Value float64
}

// generate a monochromatic color scheme
// h refers the base hue to generate a color scheme around.
// n refers to the number of colors to generate. If n <= 0, no colors will be generated.
// These colors will have same Saturation and Value.
func ColorSchemeMonochromatic(h int, n int) []*HSV {
	if n <= 0 {
		return []*HSV{}
	}

	v := []*HSV{}

	for i := 0; i < n; i++ {
		p := float64(i) / (float64(n) - 1)
		v = append(v, &HSV{
			Hue:        h,
			Saturation: p,
			Value:      p,
		})
	}

	return v
}

// generate a monochromatic color scheme, but with RGB as input & output.
// baseR, baseG, baseB refer to colors that are used as a base for the scheme
// n refers to the number of colors to generate.
// This is a utility function! It wraps ColorSchemeMonochromatic to convert into RGB values!
// View ColorSchemeMonochromatic's documentation for info on generation
func ColorSchemeMonochromaticRGB(baseR, baseG, baseB uint8, n int) [][3]uint8 {
	h, _, _ := RGBToHSV(baseR, baseG, baseB)
	raw := ColorSchemeMonochromatic(h, n)
	out := make([][3]uint8, len(raw))
	for i, v := range raw {
		r, g, b := HSVToRGB(v.Hue, v.Saturation, v.Value)
		out[i] = [3]uint8{r, g, b}
	}

	return out
}

// Generate a list of 3 hues for an Analogous color scheme, with baseH being the base hue.
func ColorSchemeAnalogous(baseH int) [3]int {
	// normalize it first
	baseH = baseH % 360
	v := [3]int{baseH - 60, baseH, baseH + 60}
	if baseH < 60 {
		v[0] += 360
	}
	if baseH > 300 {
		v[1] -= 360
	}

	return v
}

// Generate a list of 3 hues for an Analogous color scheme.
// Hue, Saturation and lightness are taken from baseR, baseG, baseB.
func ColorSchemeAnalogousRGB(baseR, baseG, baseB uint8) [3][3]uint8 {
	h, s, l := RGBToHSL(baseR, baseG, baseB)

	v := [3][3]uint8{}

	items := ColorSchemeAnalogous(h)

	for i, val := range items {
		r, g, b := HSLToRGB(val, s, l)
		v[i] = [3]uint8{r, g, b}
	}

	return v
}

// generate a complimentary hue to baseH
func ColorSchemeComplementary(baseH int) (h int) {
	h = (baseH + 180) % 360

	return
}

// generate a complimentary color to baseR, baseG, baseB.
// This is not recommended, use ColorSchemeComplementary instead, as it allows for more flexibility for saturation/lightness
func ColorSchemeComplementaryRGB(baseR, baseG, baseB uint8) (r, g, b uint8) {
	h, s, v := RGBToHSV(baseR, baseG, baseB)
	h = ColorSchemeComplementary(h)
	r, g, b = HSVToRGB(h, s, v)

	return
}

// generate 3 hues that are 120 degrees apart from baseH (normalized)
func ColorSchemeTriadic(baseH int) [3]int {
	return [3]int{baseH, (baseH + 120) % 360, (baseH + 240) % 360}
}

// generate 3 colors that are 120 degrees apart from baseH (taken from baseR, baseG, baseB)
// This is not recommended, use ColorSchemeTriadic instead, as it allows for more flexibility for saturation/lightness
func ColorSchemeTriadicRGB(baseR, baseG, baseB uint8) [3][3]uint8 {
	h, s, v := RGBToHSV(baseR, baseG, baseB)

	hues := ColorSchemeTriadic(h)

	vals := [3][3]uint8{}

	for i, h := range hues {
		r, g, b := HSVToRGB(h, s, v)

		vals[i] = [3]uint8{r, g, b}
	}

	return vals
}

// generate 3 hues that are +/- 150 degrees apart from baseH (normalized)
func ColorSchemeCompound(baseH int) [3]int {
	return [3]int{baseH, (baseH + 150) % 360, (baseH + 210) % 360}
}

// generate 3 colors that are +/- 150 degrees apart from baseH (taken from baseR, baseG, baseB)
// This is not recommended, use ColorSchemeCompound instead, as it allows for more flexibility for saturation/lightness
func ColorSchemeCompoundRGB(baseR, baseG, baseB uint8) [3][3]uint8 {
	h, s, v := RGBToHSV(baseR, baseG, baseB)

	hues := ColorSchemeCompound(h)

	vals := [3][3]uint8{}

	for i, h := range hues {
		r, g, b := HSVToRGB(h, s, v)

		vals[i] = [3]uint8{r, g, b}
	}

	return vals
}
