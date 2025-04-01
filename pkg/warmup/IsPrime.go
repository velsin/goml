package warmup

import "math"

func IsPrime(n int) bool {
	// Prime sieve by eliminating base cases early, and then iterating with non-naive upper bound.
	// Only defined for n >= 2

	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 || n%5 == 0 {
		return false
	}
	for i := 5; i <= int(math.Floor(math.Sqrt(float64(n))))+1; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
