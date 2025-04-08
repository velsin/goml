package bases

import (
	"strconv"
	"strings"
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

	var sb strings.Builder

	// Find the smallest n s.t. 2^n+1 > x
	x, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	var maxExp int

	// Could also do interval bisection to find this? Depends on if we want to care more about smaller inputs or not.
	for (1 << maxExp) <= x {
		maxExp += 1
	}

	var n int
	if maxExp > 0 {
		n = maxExp - 1
	} else {
		n = 0
	}

	for x > 0 || n >= 0 {
		y := 1 << n
		if y <= x {
			x -= y
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		n -= 1
	}

	return sb.String()
}
