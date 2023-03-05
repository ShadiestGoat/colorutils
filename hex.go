package colorutils

import (
	"strconv"
	"strings"
)

func createColor(r, g, b uint8) uint64 {
	return (uint64(r)<<16)+(uint64(g)<<8)+uint64(b)
}

// Create a hexadecimal representation of a color. This will always create a 6 character representation, without a prefix like # or 0x
func Hexadecimal(r, g, b uint8) string {
	s := strconv.FormatUint(createColor(r, g, b), 16)
	s = strings.Repeat("0", 6-len(s)) + s

	return s
}


// Create a hexadecimal representation of a color, including opacity. This will always create a 6 character representation, without a prefix like # or 0x
func HexadecimalOpacity(r, g, b, a uint8) string {
	s := strconv.FormatUint((createColor(r, g, b) << 8) + uint64(a), 16)
	s = strings.Repeat("0", 8-len(s)) + s

	return s
}
