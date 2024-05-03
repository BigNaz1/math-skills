package mathskills

func CalculateVariance(data []float64, average float64) int {
	var sumSquaredDiff float64
	for _, value := range data {
		diff := value - average
		sumSquaredDiff += diff * diff
	}

	variance := sumSquaredDiff / float64(len(data))

	// rounds the variance to the nearest int
	if variance < 0 {
		return int(variance - 0.5)
	}
	return int(variance + 0.5)
}
