package array

import (
	"testing"
)

func TestArray(t *testing.T) {
	t.Run("Test init", func(t *testing.T) {
		initialCapacity := 10
		arr := NewArray(initialCapacity)

		if arr.IsEmpty() == false || arr.Size() != 0 {
			t.Error("New array must be empty")
		}

		if actualCapacity := arr.Capacity(); actualCapacity != initialCapacity {
			t.Errorf("New array capacity must equal %d, found %d", initialCapacity, actualCapacity)
		}
	})

	t.Run("Test append", func(t *testing.T) {
		arr := NewArray(10)
		expectedSize := 5

		for i := 0; i < expectedSize; i++ {
			arr.Append(i * 10)
		}

		if actualSize := arr.Size(); actualSize != expectedSize {
			t.Errorf("Array size must equal %d, found %d", expectedSize, actualSize)
		}

		for i := 0; i < expectedSize; i++ {
			if value, _ := arr.ValueAt(i); value != i*10 {
				t.Errorf("Value at index %d must equal %d, found %d", i, i*10, value)
			}
		}
	})

	t.Run("Test insert", func(t *testing.T) {
		arr := NewArray(10)
		insertIndex := 2
		insertValue := 99

		if err := arr.Insert(insertValue, insertIndex); err != nil {
			t.Errorf("Insert to index %d must not fail", insertIndex)
		}

		if value, _ := arr.ValueAt(insertIndex); value != insertValue {
			t.Errorf("Inserted value at index %d must equal %d", insertIndex, insertValue)
		}

		insertIndex = -1

		if err := arr.Insert(insertValue, insertIndex); err == nil {
			t.Errorf("Insert to index %d must fail", insertIndex)
		}

		insertIndex = arr.Capacity()

		if err := arr.Insert(insertValue, insertIndex); err == nil {
			t.Errorf("Insert to index %d must fail", insertIndex)
		}
	})

	t.Run("Test pop", func(t *testing.T) {
		arr := NewArray(10)

		for i := 0; i < 5; i++ {
			arr.Append(i * 10)
		}

		if value, _ := arr.Pop(); value != 40 {
			t.Errorf("Popped value must equal 40, found %d", value)
		}

		for i := 0; i < 4; i++ {
			arr.Pop()
		}

		if _, err := arr.Pop(); err == nil {
			t.Error("Popping an empty array must fail")
		}
	})

	t.Run("Test remove", func(t *testing.T) {
		arr := NewArray(10)
		removeIndex := 2

		if _, err := arr.RemoveAt(removeIndex); err == nil {
			t.Error("Remove a value from an empty array must fail")
		}

		for i := 0; i < 5; i++ {
			arr.Append(i * 10)
		}

		if _, err := arr.RemoveAt(removeIndex); err != nil {
			t.Errorf("Remove value at index %d must not fail", removeIndex)
		}
	})

	t.Run("Test shrink", func(t *testing.T) {
		arr := NewArray(10)
		expectedCapacity := 80

		for i := 0; i < 50; i++ {
			arr.Append(i)
		}

		if actualCapacity := arr.Capacity(); actualCapacity != expectedCapacity {
			t.Errorf("Array capacity must equal %d, found %d", expectedCapacity, actualCapacity)
		}
	})

	t.Run("Test reduce", func(t *testing.T) {
		arr := NewArray(80)
		expectedCapacity := 40

		for i := 0; i < 20; i++ {
			arr.Append(i)
		}

		arr.Pop()

		if actualCapacity := arr.Capacity(); actualCapacity != expectedCapacity {
			t.Errorf("Array capacity must equal %d, found %d", expectedCapacity, actualCapacity)
		}
	})
}
