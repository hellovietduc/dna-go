package heap

// MaxHeap is a struct for storing values into a Max Heap.
type MaxHeap struct {
	items *NodeArray
}

// NewMaxHeap creates a new MaxHeap instance.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		items: NewNodeArray(),
	}
}

// Peek returns the root value of the MaxHeap.
func (h *MaxHeap) Peek() (string, int) {
	if node := h.items.ValueAt(0); node != nil {
		return node.value, node.key
	}
	return "", 0
}

// Insert adds the value to the MaxHeap.
func (h *MaxHeap) Insert(key int, value string) {
	node := &Node{key, value}
	h.items.Append(node)
	h.shiftUp(h.getLastIndex(), key)
}

// Delete removes the value from the MaxHeap.
func (h *MaxHeap) ExtractMax() {
	h.items.Swap(0, h.getLastIndex())
	h.items.Pop()
	_, newMax := h.Peek()
	h.shiftDown(0, newMax)
}

func (h *MaxHeap) shiftUp(childIndex, childKey int) {
	if !h.hasParent(childIndex) {
		return
	}

	parentIndex := h.getParentIndex(childIndex)
	parent := h.items.ValueAt(parentIndex)

	if childKey > parent.key {
		h.items.Swap(childIndex, parentIndex)
		h.shiftUp(parentIndex, childKey)
	}
}

func (h *MaxHeap) shiftDown(parentIndex, parentKey int) {
	hasLeftChild := h.hasLeftChild(parentIndex)
	hasRightChild := h.hasRightChild(parentIndex)
	hasBothChildren := hasLeftChild && hasRightChild

	if !hasBothChildren {
		return
	}

	leftChildIndex := h.getLeftChildIndex(parentIndex)
	rightChildIndex := h.getRightChildIndex(parentIndex)
	leftChild := h.items.ValueAt(leftChildIndex)
	rightChild := h.items.ValueAt(rightChildIndex)

	var childIndex, childKey int
	if hasBothChildren {
		if leftChild.key > rightChild.key {
			childIndex = leftChildIndex
			childKey = leftChild.key
		} else {
			childIndex = rightChildIndex
			childKey = rightChild.key
		}
	} else if hasLeftChild {
		childIndex = leftChildIndex
		childKey = leftChild.key
	} else {
		childIndex = rightChildIndex
		childKey = rightChild.key
	}

	if childKey > parentKey {
		h.items.Swap(childIndex, parentIndex)
		h.shiftDown(childIndex, parentKey)
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
