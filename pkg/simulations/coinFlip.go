package simulations

func coinFlip() int {
	draw := rng.Float64()
	if draw <= 0.5 {
		return 0
	} else {
		return 1
	}
}

func SimProbability(nHeads, nFlips int) float64 {
	// Computes the probability of obtaining nHeads heads in nFlips coin flips by simulation
	var nSims int = 10000
	var successes float64

	for range nSims {
		sumHeads := 0
		for range nFlips {
			sumHeads += coinFlip()
		}
		if sumHeads == nHeads {
			successes += 1
		}
	}

	return successes / float64(nSims)
}
