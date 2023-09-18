package colorutils

import "math"

func mColor(p, m float64) uint8 {
	return uint8(math.Round((p + m) * 255))
}

func hSwitch(h int, c, x, m float64) (r, g, b uint8) {
	var (
		rp float64
		gp float64
		bp float64
	)

	switch math.Floor(float64(h%360) / 60) {
	case 0:
		// 0 <= h < 60
		rp, gp = c, x
	case 1:
		// 60 <= h < 120
		rp, gp = x, c
	case 2:
		// 120 <= h < 180
		gp, bp = c, x
	case 3:
		// 180 <= h < 240
		gp, bp = x, c
	case 4:
		// 240 <= h < 300
		rp, bp = x, c
	case 5:
		// 300 <= h < 360
		rp, bp = c, x
	}

	r, g, b = mColor(rp, m), mColor(gp, m), mColor(bp, m)

	return
}

func dSwitch(d float64, cMax rune, primes map[rune]float64) int {
	if d == 0 {
		return 0
	}

	switch cMax {
	case 'r':
		return int(math.Round((primes['g']-primes['b'])/d)) % 6
	case 'g':
		return int(math.Round((primes['b']-primes['r'])/d)) + 2
	case 'b':
		return int(math.Round((primes['r']-primes['g'])/d)) + 4
	}

	return 0
}

func xValue(c float64, h int) float64 {
	return c * (1 - math.Abs(float64(int(math.Round(float64(h)/60))%2-1)))
}
