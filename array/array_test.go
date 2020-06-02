package array

import (
	"testing"
)

func TestArray(t *testing.T) {
	t.Run("Create new array with default capacity", func(t *testing.T) {
		arr := NewArray()
		if capacity := arr.Capacity(); capacity != defaultCapacity {
			t.Errorf("New array capacity should equal %d, found %d", defaultCapacity, capacity)
		}
	})

	arr := NewArray()

	t.Run("Check size for empty array", func(t *testing.T) {
		if arr.Size() != 0 || arr.IsEmpty() != true {
			t.Error("Empty array should have 0 size")
		}
	})

	t.Run("Get value from empty array, get indexOutOfRange", func(t *testing.T) {
		if _, err := arr.ValueAt(0); err == nil {
			t.Error("Calling ValueAt on empty array should return indexOutOfRange")
		}
	})

	t.Run("Insert value to out-of-range index, get indexOutOfRange", func(t *testing.T) {
		if err := arr.Insert(1, 1); err == nil {
			t.Error("Insert to index 1 on empty array should return indexOutOfRange")
		}
	})

	t.Run("Remove value at out-of-range index, get indexOutOfRange", func(t *testing.T) {
		if err := arr.RemoveAt(1); err == nil {
			t.Error("Remove value at index 1 on empty array should return indexOutOfRange")
		}
	})

	t.Run("Pop value from empty array, get EmptyValue", func(t *testing.T) {
		if value := arr.Pop(); value != EmptyValue {
			t.Error("Value popped from empty array should be EmptyValue")
		}
	})

	t.Run("Append values to array, get correct appended values", func(t *testing.T) {
		for i := 1; i < 4; i++ {
			arr.Append(i)

			if value, _ := arr.ValueAt(i - 1); value != i {
				t.Errorf("Wrong value at index %d, expected %d, found %d", i-1, i, value)
			}
		}
	})

	t.Run("Insert value at specified index, get correct inserted value", func(t *testing.T) {
		arr.Insert(4, 0)

		for i := 0; i < arr.Size(); i++ {
			value, _ := arr.ValueAt(i)
			if i == 0 {
				if value != 4 {
					t.Errorf("Wrong value at index 0, expected 4, found %d", value)
				}
			} else {
				if value != i {
					t.Errorf("Wrong value at index %d, expected %d, found %d", i, i, value)
				}
			}
		}
	})

	t.Run("Append more values so array must grow in capacity", func(t *testing.T) {
		capacityBefore := arr.Capacity()
		for i := 4; i < 20; i++ {
			arr.Append(i)
		}
		capacityAfter := arr.Capacity()
		if capacityAfter <= capacityBefore {
			t.Errorf("Capacity should grow when inserting more values, before %d, after %d", capacityBefore, capacityAfter)
		}
	})

	t.Run("Remove value at specified index, get different value from that index", func(t *testing.T) {
		arr.RemoveAt(0)

		if value, _ := arr.ValueAt(0); value == 4 {
			t.Error("Wrong value at index 0, value 4 removed")
		}
	})

	t.Run("Pop from unempty array, get correct value", func(t *testing.T) {
		if arr.Pop() != 19 {
			t.Error("Pop value should be 19")
		}
	})

	t.Run("Remove more values so array must shrink in capacity", func(t *testing.T) {
		capacityBefore := arr.Capacity()
		for i := 0; i < 18; i++ {
			arr.Pop()
		}
		capacityAfter := arr.Capacity()
		if capacityAfter >= capacityBefore {
			t.Errorf("Capacity should shrink when removing more values, before %d, after %d", capacityBefore, capacityAfter)
		}
	})
}
