package sorting_test

import (
	"sort"
	"testing"
	"time"

	"github.com/vietduc01100001/dna-go/sorting"
)

func TestHeapSort(t *testing.T) {
	empty := []int{}
	allOne := []int{1, 1, 1, 1, 1}
	noDuplicates := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	hasDuplicates := []int{3, 44, 38, 5, 5, 15, 36, 26, 26, 2, 46, 4, 4, 50, 48}
	ascSorted := []int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50}
	descSorted := []int{50, 48, 47, 46, 44, 38, 36, 27, 26, 19, 15, 5, 4, 3, 2}

	t.Run("on empty array", func(t *testing.T) {
		start := time.Now()
		sorting.HeapSort(empty)
		elapsed := time.Since(start)
		t.Logf("Took %s", elapsed)
		if !sort.IntsAreSorted(empty) {
			t.Errorf("Expected array to be sorted but got %v", empty)
		}
	})

	t.Run("on array whose elements are all one value", func(t *testing.T) {
		start := time.Now()
		sorting.HeapSort(allOne)
		elapsed := time.Since(start)
		t.Logf("Took %s", elapsed)
		if !sort.IntsAreSorted(allOne) {
			t.Errorf("Expected array to be sorted but got %v", allOne)
		}
	})

	t.Run("on array which has no duplicates", func(t *testing.T) {
		start := time.Now()
		sorting.HeapSort(noDuplicates)
		elapsed := time.Since(start)
		t.Logf("Took %s", elapsed)
		if !sort.IntsAreSorted(noDuplicates) {
			t.Errorf("Expected array to be sorted but got %v", noDuplicates)
		}
	})

	t.Run("on array which has duplicates", func(t *testing.T) {
		start := time.Now()
		sorting.HeapSort(hasDuplicates)
		elapsed := time.Since(start)
		t.Logf("Took %s", elapsed)
		if !sort.IntsAreSorted(hasDuplicates) {
			t.Errorf("Expected array to be sorted but got %v", hasDuplicates)
		}
	})

	t.Run("on array which is already ascending sorted", func(t *testing.T) {
		start := time.Now()
		sorting.HeapSort(ascSorted)
		elapsed := time.Since(start)
		t.Logf("Took %s", elapsed)
		if !sort.IntsAreSorted(ascSorted) {
			t.Errorf("Expected array to be sorted but got %v", ascSorted)
		}
	})

	t.Run("on array which is already descending sorted", func(t *testing.T) {
		start := time.Now()
		sorting.HeapSort(descSorted)
		elapsed := time.Since(start)
		t.Logf("Took %s", elapsed)
		if !sort.IntsAreSorted(descSorted) {
			t.Errorf("Expected array to be sorted but got %v", descSorted)
		}
	})
}
