package warmup

// Recursive functions will absolutely need some kind of memoization
// so we will go with the naive approach first, not-concurrency safe?
// all given recursive sequences are int->int, so we start with that
// Based on example from gopl.io repo

type PureFuncIntInt func(int) int

func Memoized(fn PureFuncIntInt) func(int) int {

	cache := make(map[int]int)

	wrappedFunc := func(in int) int {
		// Start by checking for a cache hit
		val, ok := cache[in]
		if !ok {
			// Compute if missing
			val = fn(in)
			cache[in] = val
		}
		return val
	}

	return wrappedFunc
}

// We use another functional for sequence generation, that can be shared
// across the sequences and used for timing comparisons

func GenerateSequence(fn PureFuncIntInt, n int) []int {
	// Use the given function to generate an array with the function output for
	// all integers from 1 up to and including n

	res := make([]int, n)

	for i := 1; i <= n; i++ {
		res[i-1] = fn(i)
	}

	return res
}

func sequence1(in int) int {
	// Base cases
	if in == 1 {
		return 5
	}
	return (sequence1(in-1) * 3) - 4
}

var Sequence1 = Memoized(sequence1)
