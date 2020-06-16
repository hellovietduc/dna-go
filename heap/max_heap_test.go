package heap

import "testing"

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap()

	t.Run("Insert values into the heap, peek max value at root", func(t *testing.T) {
		tasks := map[string]int{
			"Wash the dishes":                    1,
			"Clean the house":                    3,
			"Take out the trash":                 5,
			"Go out with girlfriend":             10,
			"Have dinner with another girlfiend": 20,
			"Go home with wife":                  100,
		}

		for task, priority := range tasks {
			heap.Insert(priority, task)
		}

		if task, _ := heap.Peek(); task != "Go home with wife" {
			t.Errorf("Expected root to be 'Go home with wife', found %s", task)
		}
	})

	// t.Run("Delete values from the heap, peek new max value", func(t *testing.T) {
	// 	deletedValues := []int{9, 8, 7, 6}
	// 	for _, value := range deletedValues {
	// 		heap.Delete(value)
	// 	}

	// 	expectedNewRoot := deletedValues[len(deletedValues)-1]
	// 	if root := heap.Peek(); root < expectedNewRoot {
	// 		t.Errorf("Expected root to be %d, found %d", expectedNewRoot, root)
	// 	}
	// })
}
