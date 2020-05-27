package hash_table

import (
	"testing"
	"time"
)

func TestHashTable(t *testing.T) {
	hashTable := NewHashTable()
	expectedSize := 500

	t.Run("Test insert", func(t *testing.T) {
		for i := 0; i < expectedSize; i++ {
			start := time.Now()
			hashTable.Insert(i, i*10)
			elapsed := time.Since(start)
			if elapsed > time.Microsecond {
				t.Logf("Insert key %d took %s", i, elapsed)
			}
		}

		if actualSize := hashTable.Size(); actualSize != expectedSize {
			t.Errorf("Hash table size must be %d, found %d", expectedSize, actualSize)
		}
	})

	t.Run("Test search", func(t *testing.T) {
		for i := 0; i < expectedSize; i++ {
			value, err := hashTable.Search(i)

			if err != nil {
				t.Errorf("Search with key %d must return value", i)
			}

			if value != i*10 {
				t.Errorf("Search with key %d, expected %d, found %d", i, i*10, value)
			}
		}
	})

	t.Run("Test delete", func(t *testing.T) {
		for i := 0; i < expectedSize; i++ {
			start := time.Now()
			if _, err := hashTable.Delete(i); err != nil {
				t.Errorf("Delete with key %d must not fail", i)
			} else {
				elapsed := time.Since(start)
				if elapsed > time.Microsecond {
					t.Logf("Delete key %d took %s", i, elapsed)
				}
			}
		}
	})
}
