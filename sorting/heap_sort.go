package sorting

func HeapSort(array []int) {
	length := len(array)
	if length <= 1 {
		return
	}

	buildMaxHeap(array, length)

	for i := length - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		shiftDown(array, i, 0)
	}
}

func buildMaxHeap(array []int, length int) {
	// skip all leaf nodes
	for i := length / 2; i >= 0; i-- {
		shiftDown(array, length, i)
	}
}

func shiftDown(array []int, length, index int) {
	largest := index
	left := index*2 + 1
	right := left + 1

	if left < length && array[left] > array[largest] {
		largest = left
	}

	if right < length && array[right] > array[largest] {
		largest = right
	}

	if largest != index {
		array[largest], array[index] = array[index], array[largest]
		shiftDown(array, length, largest)
	}
}
