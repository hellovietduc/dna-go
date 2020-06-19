package sorting

func InsertionSort(array []int) {
	if len(array) <= 1 {
		return
	}

	// consider first element as sorted
	for i := 1; i < len(array); i++ {
		index := i
		for index > 0 {
			// swap until the element before is smaller
			if array[index] < array[index-1] {
				array[index], array[index-1] = array[index-1], array[index]
				index--
			} else {
				break
			}
		}
	}
}
