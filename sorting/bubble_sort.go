package sorting

func BubbleSort(array []int) {
	if len(array) <= 1 {
		return
	}

	for i := 0; i < len(array)-1; i++ {
		swapped := false
		for j := 0; j < len(array)-i-1; j++ {
			if array[j+1] < array[j] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
