package mathskills

func CalculateAverage(data []float64) int {
	var sum float64
	for _, value := range data {
		sum += value
	}
	average := sum / float64(len(data))

	// Round the average to the nearest integer
	if average < 0 {
		return int(average - 0.5)
	}
	return int(average + 0.5)
}
