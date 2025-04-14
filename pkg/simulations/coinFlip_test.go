package simulations

import (
	"math"
	"testing"
)

func TestSimProbability(t *testing.T) {
	testCases := []struct {
		nHeads int
		nFlips int
		exp    float64
	}{
		{1, 1, 0.5},
		{2, 2, 0.25},
		{5, 5, math.Pow(0.5, 5)},
		{2, 10, math.Pow(0.5, 10) * 45},
		{3, 5, math.Pow(0.5, 5) * 10},
	}

	tol := 0.2

	for _, testCase := range testCases {
		res := SimProbability(testCase.nHeads, testCase.nFlips)
		if math.Abs(res-testCase.exp)/testCase.exp > tol {
			t.Errorf("Failed for %d heads in %d flips, probability was more that %v off from expected value of %f. Got %f.\n", testCase.nHeads, testCase.nFlips, tol, testCase.exp, res)
		}

	}
}
