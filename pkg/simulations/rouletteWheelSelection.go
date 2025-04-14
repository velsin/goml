package simulations

func randomDraw(dist []float64) int {
	// Takes in any discrete probability distribution, and returns an index drawn from that
	// distribution. The sum of all elements in dist must be 1.

	draw := rng.Float64()
	for i, prob := range dist {
		if draw <= prob {
			return i
		}
	}
	return 0
}

func SampleMany(n int, dist []float64) []float64 {
	// Samples from the given discrete probability distribution n time, and returns the
	// relative frequency of each of the resulting draws

	cumDist := make([]float64, len(dist))

	total := 0.0
	for i, prob := range dist {
		total += prob
		cumDist[i] = total
	}

	if cumDist[len(cumDist)-1] != 1.0 {
		panic("Invalid probability distribution provided, total sum != 1.0")
	}

	counts := make([]float64, len(dist))

	for range n {
		counts[randomDraw(cumDist)] += 1
	}

	res := make([]float64, len(dist))
	for i, count := range counts {
		res[i] = count / float64(n)
	}

	return res
}
