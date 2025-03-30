package warmup

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	test_cases := []struct {
		in  string
		exp bool
	}{
		{"notapalindrome", false},
		{"racecar", true},
		{"rotator", true},
		{"rot.a.tor", true},
		{"mumbojumbo", false},
		{"abccba", true},
		{"", true},
		{"a", true},
	}

	for _, test_case := range test_cases {
		res := IsSymmetric(test_case.in)
		if res != test_case.exp {
			t.Errorf("Got %t for %s, expected %t", res, test_case.in, test_case.exp)
		}
	}

}
