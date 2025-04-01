package warmup

import "strings"

func CountCharacters(s string) map[rune]int {
	res := make(map[rune]int)

	for _, r := range strings.ToLower(s) {

		res[r] = res[r] + 1
	}

	return res
}
