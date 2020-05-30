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
func NewArray(args ...int) *Array {
	capacity := defaultCapacity
	if args != nil && args[0] > 0 {
		capacity = args[0]
	}
	return &Array{
		items: make([]int, capacity),
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
	if index < 0 || index >= a.size {
		return errors.New(indexOutOfRange)
	}

	if a.getLoadFactor() >= upperLoadFactor {
		a.changeCapacity(true)
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
		a.changeCapacity(true)
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
		a.changeCapacity(false)
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
		a.changeCapacity(false)
	}

	a.size--
	return a.items[a.size]
}

func (a *Array) getLoadFactor() float64 {
	return float64(a.size) / float64(len(a.items))
}

func (a *Array) changeCapacity(isGrow bool) {
	oldCapacity := len(a.items)
	var newCapacity int
	if isGrow {
		newCapacity = oldCapacity * growShrinkFactor
	} else {
		newCapacity = oldCapacity / growShrinkFactor
	}

	newArr := make([]int, newCapacity)
	for i := 0; i < a.size; i++ {
		newArr[i] = a.items[i]
	}

	a.items = newArr
}
