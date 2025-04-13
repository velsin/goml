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

func TestSequence2(t *testing.T) {
	testCases := []struct {
		in  int
		exp int
	}{
		{1, 25},
		{2, 76},
		{3, 38},
		{4, 19},
	}

	for _, test_case := range testCases {
		res := Sequence2(test_case.in)
		if res != test_case.exp {
			t.Errorf("Got %d for %d, expected %d", res, test_case.in, test_case.exp)
		}
	}
}

func TestSequence3(t *testing.T) {
	testCases := []struct {
		in  int
		exp int
	}{
		{1, 2},
		{2, -3},
		{3, -6},
		{4, 18},
		{5, -108},
	}

	for _, test_case := range testCases {
		res := Sequence3(test_case.in)
		if res != test_case.exp {
			t.Errorf("Got %d for %d, expected %d", res, test_case.in, test_case.exp)
		}
	}
}

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		in  int
		exp int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 3},
		{6, 5},
	}

	for _, test_case := range testCases {
		res := Fibonacci(test_case.in)
		if res != test_case.exp {
			t.Errorf("Got %d for %d, expected %d", res, test_case.in, test_case.exp)
		}
	}
}

// --- Benchmarks ---

const benchmarkSequenceLength = 30

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

func BenchmarkSequence2Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateSequence(sequence2, benchmarkSequenceLength)
	}
}

func BenchmarkSequence2Memoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateSequence(Sequence2, benchmarkSequenceLength)
	}
}

func BenchmarkSequence3Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Set much lower seq length since this one quickly explodes past int64 size?
		_ = GenerateSequence(sequence3, benchmarkSequenceLength)
	}
}

func BenchmarkSequence3Memoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateSequence(Sequence3, benchmarkSequenceLength)
	}
}

func BenchmarkFibonacciNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateSequence(fibonacci, benchmarkSequenceLength)
	}
}

func BenchmarkFibonacciMemoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateSequence(Fibonacci, benchmarkSequenceLength)
	}
}
