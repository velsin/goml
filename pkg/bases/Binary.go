package bases

import (
	"strconv"
)

func BinaryToDecimal(in string) string {
	// Transform the binary format string representation of a number to its decimal format string
	// representation

	var res int
	runes := []rune(in)
	maxExp := len(runes) - 1

	for i := 0; i <= maxExp; i++ {
		// We get the powers of two easily by bit-shifting by that amount
		res += int(runes[i]-'0') * (1 << (maxExp - i))
	}

	return strconv.Itoa(res)
}

func DecimalToBinary(in string) string {
	// Transfrom the decimal format string representation of a number to its binary format string
	// representation

	var res string

	return "1"
}
