package mathskills

import (
	"sort"
)

func CalculateMedian(data []float64) int {
	sort.Float64s(data)
	n := len(data)
	if n == 0 {
		return 0
	}

	if n%2 == 0 {
		// If the numb of elements is even, calc the avg of two middle values
		middleLeft := data[n/2-1]
		middleRight := data[n/2]
		median := (middleLeft + middleRight) / 2

		// rounds the median to the nearest int
		if median < 0 {
			return int(median - 0.5)
		}
		return int(median + 0.5)
	}

	// If the numb of elements is odd, return the middle value
	return int(data[n/2])
}
