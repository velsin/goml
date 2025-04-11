package warmup

import "testing"

// --- Tests ---

func TestSequence1(t *testing.T) {
	testCases := []struct {
		in  int
		exp int
	}{
		{1, 5},
		{2, 11},
		{3, 29},
	}

	for _, test_case := range testCases {
		res := Sequence1(test_case.in)
		if res != test_case.exp {
			t.Errorf("Got %d for %d, expected %d", res, test_case.in, test_case.exp)
		}
	}
}

// --- Benchmarks ---

const benchmarkSequenceLength = 50

func BenchmarkSequence1Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateSequence(sequence1, benchmarkSequenceLength)
	}
}

func BenchmarkSequence1Memoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateSequence(Sequence1, benchmarkSequenceLength)
	}
	// It seems like this version might actually persist the same cache across
	// all b.N benchmark runs, favoring the memoized version even more.
}
