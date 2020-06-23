package sorting

import "math"

func CountingSort(array []int) {
	if len(array) <= 1 {
		return
	}

	max := math.MinInt64
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	counts := make([]int, max+1)
	for i := 0; i < len(array); i++ {
		num := array[i]
		counts[num]++
	}

	sorted := 0
	for i := 0; i <= max; i++ {
		for j := 0; j < counts[i]; j++ {
			array[sorted] = i
			sorted++
		}
	}
}
