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

func sequence1(n int) int {
	if n == 1 {
		return 5
	}
	return (sequence1(n-1) * 3) - 4
}

var Sequence1 = Memoized(sequence1)

func sequence2(n int) int {
	if n == 1 {
		return 25
	}

	prev := sequence2(n - 1)

	if prev%2 == 0 {
		return prev / 2
	} else {
		return (prev * 3) + 1
	}
}

var Sequence2 = Memoized(sequence2)

func sequence3(n int) int {
	if n == 1 {
		return 2
	}
	if n == 2 {
		return -3
	}
	return sequence3(n-1) * sequence3(n-2)
}

var Sequence3 = Memoized(sequence3)

func fibonacci(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

var Fibonacci = Memoized(fibonacci)
