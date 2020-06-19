package sorting

func SelectionSort(array []int) {
	if len(array) <= 1 {
		return
	}

	// the sorted part is of length zero now
	sorted := 0
	for i := 0; i < len(array); i++ {
		min := array[i]
		minIndex := i

		// find the min element from the unsorted part
		for j := sorted + 1; j < len(array); j++ {
			if array[j] < min {
				min = array[j]
				minIndex = j
			}
		}

		// swap the min element found with the first element
		// in the unsorted part
		array[sorted], array[minIndex] = array[minIndex], array[sorted]
		sorted++
	}
}
