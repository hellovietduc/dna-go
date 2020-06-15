package array

import (
	"errors"
)

const (
	// EmptyValue is a constant indicates that
	// a value returned from the Array is empty.
	EmptyValue       = 0
	defaultCapacity  = 8
	growShrinkFactor = 2
	upperLoadFactor  = 1.0
	lowerLoadFactor  = 0.25
	indexOutOfRange  = "Index is out of range"
)

// Array is a struct for storing linear equal-size data.
type Array struct {
	items []int
	size  int
}

// NewArray creates a new Array instance.
func NewArray() *Array {
	return &Array{
		items: make([]int, defaultCapacity),
	}
}

// Size returns the number of items being stored in the Array.
func (a *Array) Size() int {
	return a.size
}

// Capacity returns the current maximum number of items
// that can be stored in the Array.
func (a *Array) Capacity() int {
	return len(a.items)
}

// IsEmpty returns whether the Array has no items or not.
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// ValueAt returns the value stored at the given index.
func (a *Array) ValueAt(index int) (int, error) {
	if index < 0 || index >= a.size {
		return EmptyValue, errors.New(indexOutOfRange)
	}
	return a.items[index], nil
}

// Insert adds the value to the given index and shift
// the items from that index to the right.
func (a *Array) Insert(value int, index int) error {
	if index < 0 || index > a.size {
		return errors.New(indexOutOfRange)
	}

	if a.getLoadFactor() >= upperLoadFactor {
		a.grow()
	}

	// shift array to the right
	for i := a.size; i > index; i-- {
		a.items[i] = a.items[i-1]
	}

	a.items[index] = value
	a.size++

	return nil
}

// Append adds the value to the end of the Array.
func (a *Array) Append(value int) {
	if a.getLoadFactor() >= upperLoadFactor {
		a.grow()
	}

	a.items[a.size] = value
	a.size++
}

// RemoveAt removes the value at the given index and
// shift the items from that index to the left.
func (a *Array) RemoveAt(index int) error {
	if index < 0 || index >= a.size {
		return errors.New(indexOutOfRange)
	}

	if a.getLoadFactor() < lowerLoadFactor {
		a.shrink()
	}

	// shift array to the left
	for i := index; i < a.size; i++ {
		a.items[i] = a.items[i+1]
	}

	a.size--
	return nil
}

// Pop removes that last item in the Array.
func (a *Array) Pop() int {
	if a.IsEmpty() {
		return EmptyValue
	}

	if a.getLoadFactor() < lowerLoadFactor {
		a.shrink()
	}

	a.size--
	return a.items[a.size]
}

// Swap swaps the values at the given indices in-place.
func (a *Array) Swap(index1, index2 int) {
	if index1 < 0 || index1 >= a.size || index2 < 0 || index2 >= a.size {
		return
	}
	tmp := a.items[index1]
	a.items[index1] = a.items[index2]
	a.items[index2] = tmp
}

func (a *Array) getLoadFactor() float64 {
	return float64(a.size) / float64(len(a.items))
}

func (a *Array) grow() {
	newCapacity := len(a.items) * growShrinkFactor
	a.resize(newCapacity)
}

func (a *Array) shrink() {
	newCapacity := len(a.items) / growShrinkFactor
	a.resize(newCapacity)
}

func (a *Array) resize(newCapacity int) {
	newArr := make([]int, newCapacity)
	for i := 0; i < a.size; i++ {
		newArr[i] = a.items[i]
	}
	a.items = newArr
}
