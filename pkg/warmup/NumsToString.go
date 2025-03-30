package warmup

func NumsToString(input []int) string {
	res := make([]rune, len(input))

	for i, r := range input {
		res[i] = numToRune(r)
	}
	return string(res)
}

func numToRune(i int) rune {
	// Convert an integer into a "natural" rune representation. a is mapped to 1, z to 26, A to
	// 27, and Z to 52. 0 is mapped to space, along with all other integers outside the latin
	// alphabet range.
	if 1 <= i && i <= 26 {
		return rune(i + 'a' - 1)
	} else if 27 <= i && i <= 52 {
		return rune(i + 'A' - 27)
	}
	return ' '
}
