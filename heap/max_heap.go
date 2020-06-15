package heap

import "github.com/vietduc01100001/dna-go/array"

const (
	// EmptyValue is a constant indicates that
	// a value returned from the MaxHeap is empty.
	EmptyValue = 0
)

// MaxHeap is a struct for storing values into a Max Heap.
type MaxHeap struct {
	items *array.Array
}

// NewMaxHeap creates a new MaxHeap instance.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		items: array.NewArray(),
	}
}

// Peek returns the root value of the MaxHeap.
func (h *MaxHeap) Peek() int {
	value, err := h.items.ValueAt(0)
	if err != nil {
		return EmptyValue
	}
	return value
}

// Insert adds the value to the MaxHeap.
func (h *MaxHeap) Insert(value int) {
	h.items.Append(value)
	h.shiftUp(h.getLastIndex())
}

// Delete removes the value from the MaxHeap.
func (h *MaxHeap) Delete(value int) {
	h.items.Swap(0, h.getLastIndex())
	h.items.RemoveAt(h.getLastIndex())
	h.shiftDown(0)
}

func (h *MaxHeap) shiftUp(childIndex int) {
	if !h.hasParent(childIndex) {
		return
	}

	parentIndex := h.getParentIndex(childIndex)
	parentValue, _ := h.items.ValueAt(parentIndex)
	childValue, _ := h.items.ValueAt(childIndex)

	if childValue > parentValue {
		h.items.Swap(childIndex, parentIndex)
		h.shiftUp(parentIndex)
	}
}

func (h *MaxHeap) shiftDown(parentIndex int) {
	hasLeftChild := h.hasLeftChild(parentIndex)
	hasRightChild := h.hasRightChild(parentIndex)

	if !hasLeftChild && !hasRightChild {
		return
	}

	parentValue, _ := h.items.ValueAt(parentIndex)
	swapValue := 0
	swapIndex := -1

	if hasLeftChild && hasRightChild {
		leftChildIndex := h.getLeftChildIndex(parentIndex)
		rightChildIndex := h.getRightChildIndex(parentIndex)
		leftChildValue, _ := h.items.ValueAt(leftChildIndex)
		rightChildValue, _ := h.items.ValueAt(rightChildIndex)
		if leftChildValue > rightChildValue {
			swapValue = leftChildValue
			swapIndex = leftChildIndex
		} else {
			swapValue = rightChildIndex
			swapIndex = rightChildIndex
		}
	} else if hasLeftChild {
		leftChildIndex := h.getLeftChildIndex(parentIndex)
		leftChildValue, _ := h.items.ValueAt(leftChildIndex)
		swapValue = leftChildValue
		swapIndex = leftChildIndex
	} else {
		rightChildIndex := h.getRightChildIndex(parentIndex)
		rightChildValue, _ := h.items.ValueAt(rightChildIndex)
		swapValue = rightChildValue
		swapIndex = rightChildIndex
	}

	if swapValue > parentValue {
		h.items.Swap(swapIndex, parentIndex)
		h.shiftDown(swapIndex)
	}
}

func (h *MaxHeap) getLastIndex() int {
	return h.items.Size() - 1
}

func (h *MaxHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *MaxHeap) getLeftChildIndex(parentIndex int) int {
	return (parentIndex * 2) + 1
}

func (h *MaxHeap) getRightChildIndex(parentIndex int) int {
	return (parentIndex * 2) + 2
}

func (h *MaxHeap) hasParent(childIndex int) bool {
	if childIndex == 0 {
		return false
	}
	return h.getParentIndex(childIndex) >= 0
}

func (h *MaxHeap) hasLeftChild(parentIndex int) bool {
	return h.getLeftChildIndex(parentIndex) < h.items.Size()
}

func (h *MaxHeap) hasRightChild(parentIndex int) bool {
	return h.getRightChildIndex(parentIndex) < h.items.Size()
}
