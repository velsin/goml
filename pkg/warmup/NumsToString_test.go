package warmup

import (
	"testing"
)

func TestNumsToString(t *testing.T) {
	test_cases := []struct {
		in  []int
		exp string
	}{
		{[]int{1, 2, 3}, "abc"},
		{[]int{26, 26, 26}, "zzz"},
		{[]int{27, 0, 2, 9, 7, 0, 19, 20, 18}, "A big str"},
		{[]int{1, 0, 3, 1, 20}, "a cat"},
	}

	for _, test_case := range test_cases {
		res := NumsToString(test_case.in)
		if res != test_case.exp {
			t.Errorf("Got %s for %d, expected %s", res, test_case.in, test_case.exp)
		}
	}
}

func TestNumToRune(t *testing.T) {
	test_cases := []struct {
		in  int
		exp rune
	}{
		{1, 'a'},
		{0, ' '},
		{26, 'z'},
		{27, 'A'},
		{52, 'Z'},
	}
	for _, test_case := range test_cases {
		res := numToRune(test_case.in)
		if res != test_case.exp {
			t.Errorf("Got %d for %c, expected %d", res, test_case.in, test_case.exp)
		}
	}
}
