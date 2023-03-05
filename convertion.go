package colorutils

import "math"

// Convert RGB values into HSL
func RGBToHSL(r, g, b uint8) (h int, s, l float64) {
	primes := map[rune]float64{
		'r': float64(r) / 255,
		'g': float64(g) / 255,
		'b': float64(b) / 255,
	}

	cMax := 'r'
	cMin := 'r'

	for _, r := range []rune{'g', 'b'} {
		if primes[r] > primes[cMax] {
			cMax = r
		}
		if primes[r] < primes[cMin] {
			cMin = r
		}
	}

	d := primes[cMax] - primes[cMin]

	h = dSwitch(d, cMax, primes)

	h *= 60

	tmp := primes[cMax] + primes[cMin]
	l = (tmp) / 2
	s = d / (1 - math.Abs(tmp-1))

	return
}

// Convert RGB values into HSV
func RGBToHSV(r, g, b uint8) (h int, s, v float64) {
	primes := map[rune]float64{
		'r': float64(r) / 255,
		'g': float64(g) / 255,
		'b': float64(b) / 255,
	}

	cMax := 'r'
	cMin := 'r'

	for _, r := range []rune{'g', 'b'} {
		if primes[r] > primes[cMax] {
			cMax = r
		}
		if primes[r] < primes[cMin] {
			cMin = r
		}
	}

	d := primes[cMax] - primes[cMin]

	h = dSwitch(d, cMax, primes)

	h *= 60

	v = primes[cMax]

	if primes[cMax] != 0 {
		s = d / primes[cMax]
	}

	return
}

// Convert HSL values into RGB
func HSLToRGB(h int, s, l float64) (r, g, b uint8) {
	c := (1 - math.Abs(2*l-1)) * s
	x := xValue(c, h)
	m := l - c/2

	r, g, b = hSwitch(h, c, x, m)

	return
}

// Convert HSL values into HSV
func HSLToHSV(hInp int, sInp, l float64) (h int, s, v float64) {
	h = hInp

	m := l
	if l > 0.5 {
		m = 1 - l
	}

	v = l + sInp*m

	if v != 0 {
		s = 2 - 2*l/v
	}

	return
}

// Convert HSV values into RGB
func HSVToRGB(h int, s, v float64) (r, g, b uint8) {
	c := v * s
	x := xValue(c, h)
	m := v - c

	r, g, b = hSwitch(h, c, x, m)

	return
}

// Convert HSV values into HSL
func HSVToHSL(hInp int, sInp, v float64) (h int, s, l float64) {
	h = hInp
	l = v - v*sInp

	if l != 0 && l != 1 {
		m := l
		if l > 0.5 {
			m = 1 - l
		}
		s = (v - l) / m
	}

	return
}
