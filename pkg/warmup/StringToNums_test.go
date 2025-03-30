package warmup

import (
	"testing"
)

func TestStringToNums(t *testing.T) {
	test_cases := []struct {
		in  string
		exp []int
	}{
		{"abc", []int{1, 2, 3}},
		{"zzz", []int{26, 26, 26}},
		{"A big str", []int{27, 0, 2, 9, 7, 0, 19, 20, 18}},
		{"a cat", []int{1, 0, 3, 1, 20}},
	}

	for _, test_case := range test_cases {
		res := StringToNums(test_case.in)
		for i, v := range res {
			if v != test_case.exp[i] {
				t.Errorf("Got %d for %d, expected %d", v, test_case.in[i], test_case.exp[i])
			}
		}
	}
}

func TestRuneToNum(t *testing.T) {
	test_cases := []struct {
		in  rune
		exp int
	}{
		{'a', 1},
		{' ', 0},
		{'!', 0},
		{'z', 26},
		{'A', 27},
		{'Z', 52},
	}
	for _, test_case := range test_cases {
		res := runeToNum(test_case.in)
		if res != test_case.exp {
			t.Errorf("Got %d for %c, expected %d", res, test_case.in, test_case.exp)
		}
	}
}
