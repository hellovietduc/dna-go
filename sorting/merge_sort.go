package sorting

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	left := MergeSort(array[0:mid])
	right := MergeSort(array[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	sorted := make([]int, len(left)+len(right))
	i := 0
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			sorted[i] = left[l]
			i++
			l++
		} else {
			sorted[i] = right[r]
			i++
			r++
		}
	}
	for l < len(left) {
		sorted[i] = left[l]
		i++
		l++
	}
	for r < len(right) {
		sorted[i] = right[r]
		i++
		r++
	}
	return sorted
}
