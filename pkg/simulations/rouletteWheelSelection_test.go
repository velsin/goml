package simulations

import (
	"math"
	"testing"
)

func TestSampleMany(t *testing.T) {
	testCases := [][]float64{
		{0.2, 0.3, 0.5},
		{0.2, 0.1, 0.1, 0.6},
		{0.2, 0.8},
	}

	tol := 0.1
	n := 10000

	for _, testCase := range testCases {
		for idx, val := range SampleMany(n, testCase) {
			if math.Abs(val-testCase[idx])/testCase[idx] >= tol {

				t.Errorf("Failed for prob %f, got probability %f that was more that %v off from expected value.\n", testCase[idx], val, tol)
			}
		}
	}
}

func TestSampleManyPanic(t *testing.T) {
	dist := []float64{0.2, 0.5, 0.2} // invalid PMF

	defer func() {
		r := recover()

		if r == nil {
			t.Errorf("Sample many did not panic for invalid dist.")
		}

	}()

	_ = SampleMany(10, dist)

}
