package colorutils

import (
	"math"
)

const (
	value_THRESHOLD   = 0.03928
	value_DENOMINATOR = 12.92
)

func uint8ToSRGB(bit uint8) float64 {
	return float64(bit) / 255
}

func srgbToUint8(sRGB float64) uint8 {
	if sRGB > 1 {
		sRGB = 1
	}
	if sRGB < 0 {
		sRGB = 0
	}

	return uint8(math.Round(sRGB * 255))
}

func srgbToV(sRGB float64) float64 {
	if sRGB <= value_THRESHOLD {
		return sRGB / value_DENOMINATOR
	}
	return math.Pow((sRGB+0.055)/1.055, 2.4)
}

func vToSRGB(V float64) float64 {
	if V <= value_THRESHOLD/value_DENOMINATOR {
		return V * value_DENOMINATOR
	}

	return math.Pow(V, 1/2.4)*1.055 - 0.055
}

// return the maximum sRGB value that this can take. This functions limits it to 1.
func solveMax(lhs float64, factor float64) float64 {
	max := lhs / factor

	sRGB := vToSRGB(max)

	if sRGB > 1 {
		sRGB = 1
	}

	return sRGB
}

// return the minimum sRGB value that this can take. This functions limits it to 0.
func solveMin(lhs float64, factor float64, remainingFactors ...float64) float64 {
	for _, r := range remainingFactors {
		lhs -= srgbToV(1) * r
	}

	if lhs <= 0 {
		return 0
	}

	return solveMax(lhs, factor)
}
