package warmup

import "testing"

func TestIsPrime(t *testing.T) {
	test_cases := []struct {
		in  int
		exp bool
	}{
		{2, true},
		{4, false},
		{13, true},
		{13192, false},
		{1091, true},
		{7883, true},
	}

	for _, test_case := range test_cases {
		if IsPrime(test_case.in) != test_case.exp {
			t.Errorf("Expected %t for %d, got %t",
				test_case.exp, test_case.in, IsPrime(test_case.in))
		}
	}
}
