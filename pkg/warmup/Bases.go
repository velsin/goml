package warmup

import (
	"strconv"
	"strings"
)

var HexToDecMap = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15,
}
var DecToHexMap = map[int]rune{
	0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9', 10: 'A', 11: 'B', 12: 'C', 13: 'D', 14: 'E', 15: 'F',
}

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

func HexToDecimal(in string) string {
	// Transform the hexadecimal format string representation of a number to its decimal format
	// string representation

	var res int
	runes := []rune(in)
	maxExp := len(runes) - 1

	for i := 0; i <= maxExp; i++ {
		// We get the powers of two easily by bit-shifting by that amount while accounting for the
		// fact that a hexadecimal value can be represented as 4 bits
		res += HexToDecMap[runes[i]] * (1 << (4 * (maxExp - i)))
	}

	return strconv.Itoa(res)
}

func DecimalToHex(in string) string {
	// Transfrom the hexadecimal format string representation of a number to its binary format
	// string representation

	// Find the smallest n s.t. 16^n+1 > x
	x, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	var maxExp int
	var sb strings.Builder

	for (1 << (4 * maxExp)) <= x {
		maxExp += 1
	}

	var n int
	if maxExp > 0 {
		n = maxExp - 1
	} else {
		n = 0
	}

	for x > 0 || n >= 0 {
		y := 1 << (4 * n)
		if y <= x {
			// Now we find the largest multiple of the power of 16 that we can fit in the number
			z := x / y
			x -= z * y
			sb.WriteRune(DecToHexMap[z])
		} else {
			sb.WriteString("0")
		}
		n -= 1
	}

	return sb.String()
}

// For the remaining conversions we just compose the first two.
func BinaryToHex(in string) string {
	return DecimalToHex(BinaryToDecimal(in))
}

func HexToBinary(in string) string {
	return DecimalToBinary(HexToDecimal(in))
}
