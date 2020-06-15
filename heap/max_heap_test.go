package heap

import "testing"

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap()

	t.Run("Insert values into the heap, peek max value at root", func(t *testing.T) {
		expectedMax := 9
		for i := 0; i <= expectedMax; i++ {
			heap.Insert(i)
		}

		if root := heap.Peek(); root < expectedMax {
			t.Errorf("Expected root to be %d, found %d", expectedMax, root)
		}
	})

	t.Run("Delete values from the heap, peek new max value", func(t *testing.T) {
		deletedValues := []int{9, 8, 7, 6}
		for _, value := range deletedValues {
			heap.Delete(value)
		}

		expectedNewRoot := deletedValues[len(deletedValues)-1]
		if root := heap.Peek(); root < expectedNewRoot {
			t.Errorf("Expected root to be %d, found %d", expectedNewRoot, root)
		}
	})
}
