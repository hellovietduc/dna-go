package sorting

func QuickSort(array []int) {
	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(array, start, end)
	quickSort(array, start, pivot-1)
	quickSort(array, pivot+1, end)
}

func partition(array []int, start, end int) int {
	pivot := start
	storeIndex := pivot + 1
	for i := storeIndex; i <= end; i++ {
		if array[i] < array[pivot] {
			// values smaller than pivot are placed after pivot
			array[i], array[storeIndex] = array[storeIndex], array[i]
			storeIndex++
		}
	}

	// storeIndex-1 is now the last value that is smaller than pivot
	// swap them and we have pivot splitted the array into 2 sides
	// the left side contains values smaller than the pivot
	// the right side contains values bigger than the pivot
	storeIndex--
	array[pivot], array[storeIndex] = array[storeIndex], array[pivot]

	// return the index of the new pivot
	return storeIndex
}
