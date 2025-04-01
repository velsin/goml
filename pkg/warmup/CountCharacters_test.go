package warmup

import (
	"reflect"
	"testing"
)

func TestCountCharacters(t *testing.T) {
	test_cases := []struct {
		in  string
		exp map[rune]int
	}{
		{"A cat!!!", map[rune]int{'a': 2, 'c': 1, 't': 1, ' ': 1, '!': 3}},
		{"WoW, so cool", map[rune]int{'c': 1, 'l': 1, 'o': 4, 's': 1, 'w': 2, ' ': 2, ',': 1}},
	}

	for _, test_case := range test_cases {
		res := CountCharacters(test_case.in)
		eq := reflect.DeepEqual(res, test_case.exp)
		if !eq {
			t.Errorf("Failed on %v, expected %v, got %v", test_case.in, test_case.exp, res)
		}
	}
}
