package mathskills

func CalculateStandardDeviation(variance float64) int {
	if variance < 0 {
		panic("Variance cannot be negative")
	}

	std := variance
	for std*std > variance {
		std = (std + variance/std) / 2
	}

	// Round up the standard deviation to the nearest integer
	return int(std + 0.5)
}
