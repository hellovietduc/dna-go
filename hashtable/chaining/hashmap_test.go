package chaining

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	hashMap := NewHashMap()

	t.Run("Insert new keys, get correct values for inserted keys", func(t *testing.T) {
		for i := 1; i < 4; i++ {
			hashMap.Insert(i, i)

			if value := hashMap.Search(i); value != i {
				t.Errorf("Wrong value for inserted key %d, expected %d, found %d", i, i, value)
			}
		}
	})

	t.Run("Overwrite existed keys, get correct values for overwritten keys", func(t *testing.T) {
		for i := 1; i < 4; i++ {
			hashMap.Insert(i, i*10)

			if value := hashMap.Search(i); value != i*10 {
				t.Errorf("Wrong value for overwritten key %d, expected %d, found %d", i, i*10, value)
			}
		}
	})

	t.Run("Insert more keys so the HashMap must grow bigger in capacity", func(t *testing.T) {
		t.Logf("Capacity before: %d", hashMap.capacity())

		for i := 4; i < 20; i++ {
			hashMap.Insert(i, i)
		}

		t.Logf("Capacity after: %d", hashMap.capacity())
	})

	t.Run("Search for non-existed keys, get KeyNotFound", func(t *testing.T) {
		for i := 30; i > 34; i-- {
			if value := hashMap.Search(i); value != KeyNotFound {
				t.Errorf("Search for non-existed key %d should return KeyNotFound, found %d", i, value)
			}
		}
	})

	t.Run("Delete existed keys, get KeyNotFound for deleted keys", func(t *testing.T) {
		for i := 1; i < 4; i++ {
			hashMap.Delete(i)

			if value := hashMap.Search(i); value != KeyNotFound {
				t.Errorf("Search for deleted key %d should return KeyNotFound, found %d", i, value)
			}
		}
	})

	t.Run("Delete more keys so the HashMap must reduce smaller in capacity", func(t *testing.T) {
		t.Logf("Capacity before: %d", hashMap.capacity())

		for i := 4; i < 20; i++ {
			hashMap.Delete(i)
		}

		t.Logf("Capacity after: %d", hashMap.capacity())
	})
}
