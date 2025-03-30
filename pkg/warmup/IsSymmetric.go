package warmup

func IsSymmetric(input string) bool {
	n := len(input)

	if n == 0 || n == 1 {
		// Trivial cases are always true, prevents iteration
		return true
	}

	// We don't actually need to convert to runes instead of bytes for symmetry check
	for i := 0; i < n/2; i++ {
		if input[i] != input[n-i-1] {
			// Terminate as early as possible, on first difference
			return false
		}
	}
	// If no differences, then must be symmetric
	return true
}
