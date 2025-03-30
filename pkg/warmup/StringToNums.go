package warmup

func StringToNums(input string) []int {
	res := make([]int, len(input))

	for i, r := range input {
		res[i] = runeToNum(r)
	}
	return res
}

func runeToNum(r rune) int {
	// Converts a rune to it's "natural" representation, mapping 'a' to 1, 'A' to 27, etc.
	// Characters that are not part of the core latin alphabet, including ' ', are mapped to 0.
	if 'a' <= r && r <= 'z' {
		return int(r - 'a' + 1)
	} else if 'A' <= r && r <= 'Z' {
		return int(r - 'A' + 27)
	}
	return 0
}
