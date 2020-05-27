package array

import (
	"errors"
)

type Array struct {
	arr  []int
	size int
}

func NewArray(capacity int) *Array {
	return &Array{
		arr: make([]int, capacity),
	}
}

func (a *Array) Size() int {
	return a.size
}

func (a *Array) Capacity() int {
	return len(a.arr)
}

func (a *Array) IsEmpty() bool {
	return a.Size() == 0
}

func (a *Array) ValueAt(index int) (int, error) {
	if index < 0 || index >= a.Size() {
		return -1, errors.New("Index is out of range")
	}
	return a.arr[index], nil
}

func (a *Array) Insert(item int, index int) error {
	if index < 0 || index >= a.Capacity() {
		return errors.New("Index is out of range")
	}

	a.checkAndShrink()

	// shift array to the right
	for i := a.size; i > index; i-- {
		a.arr[i] = a.arr[i-1]
	}

	a.arr[index] = item

	// insert at index out of size range
	// will increase the size to index + 1
	if index >= a.Size() {
		a.size = index + 1
	} else {
		a.size++
	}

	return nil
}

func (a *Array) Append(item int) {
	a.checkAndShrink()

	a.arr[a.size] = item
	a.size++
}

func (a *Array) Pop() (int, error) {
	if a.IsEmpty() {
		return -1, errors.New("Array is empty")
	}

	a.checkAndReduce()

	a.size--
	value := a.arr[a.size]

	return value, nil
}

func (a *Array) RemoveAt(index int) (int, error) {
	if index < 0 || index >= a.Capacity() {
		return -1, errors.New("Index is out of range")
	}

	if a.IsEmpty() {
		return -1, errors.New("Array is empty")
	}

	a.checkAndReduce()

	a.size--
	value := a.arr[index]

	// shift array to the left
	for i := index; i < a.size; i++ {
		a.arr[i] = a.arr[i+1]
	}

	return value, nil
}

func (a *Array) checkAndShrink() {
	capacity := len(a.arr)
	if a.size < capacity {
		return
	}

	// size = capacity, double the capacity
	newArr := make([]int, capacity*2)
	for i := 0; i < a.size; i++ {
		newArr[i] = a.arr[i]
	}
	a.arr = newArr
}

func (a *Array) checkAndReduce() {
	capacity := len(a.arr)
	if a.size > capacity/4 {
		return
	}

	// size is <= 1/4 of capacity
	// cut half of the capacity
	newArr := make([]int, capacity/2)
	for i := 0; i < a.size; i++ {
		newArr[i] = a.arr[i]
	}
	a.arr = newArr
}
