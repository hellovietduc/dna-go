package heap

const (
	defaultCapacity  = 8
	growShrinkFactor = 2
	upperLoadFactor  = 1.0
	lowerLoadFactor  = 0.25
)

// NodeArray is a struct for storing linear equal-size data.
type NodeArray struct {
	items []*Node
	size  int
}

// New creates a new NodeArray instance.
func NewNodeArray() *NodeArray {
	return &NodeArray{
		items: make([]*Node, defaultCapacity),
	}
}

// Size returns the number of items being stored in the NodeArray.
func (a *NodeArray) Size() int {
	return a.size
}

// Capacity returns the current maximum number of items
// that can be stored in the NodeArray.
func (a *NodeArray) Capacity() int {
	return len(a.items)
}

// IsEmpty returns whether the NodeArray has no items or not.
func (a *NodeArray) IsEmpty() bool {
	return a.size == 0
}

// ValueAt returns the value stored at the given index.
func (a *NodeArray) ValueAt(index int) *Node {
	if index < 0 || index >= a.size {
		return nil
	}
	return a.items[index]
}

// Insert adds the value to the given index and shift
// the items from that index to the right.
func (a *NodeArray) Insert(node *Node, index int) {
	if index < 0 || index > a.size {
		return
	}

	if a.getLoadFactor() >= upperLoadFactor {
		a.grow()
	}

	// shift array to the right
	for i := a.size; i > index; i-- {
		a.items[i] = a.items[i-1]
	}

	a.items[index] = node
	a.size++
}

// Append adds the value to the end of the NodeArray.
func (a *NodeArray) Append(node *Node) {
	if a.getLoadFactor() >= upperLoadFactor {
		a.grow()
	}

	a.items[a.size] = node
	a.size++
}

// RemoveAt removes the value at the given index and
// shift the items from that index to the left.
func (a *NodeArray) RemoveAt(index int) {
	if index < 0 || index >= a.size {
		return
	}

	if a.getLoadFactor() < lowerLoadFactor {
		a.shrink()
	}

	// shift array to the left
	for i := index; i < a.size; i++ {
		a.items[i] = a.items[i+1]
	}

	a.size--
}

// Pop removes that last item in the NodeArray.
func (a *NodeArray) Pop() *Node {
	if a.IsEmpty() {
		return nil
	}

	if a.getLoadFactor() < lowerLoadFactor {
		a.shrink()
	}

	a.size--
	return a.items[a.size]
}

// Swap swaps the values at the given indices in-place.
func (a *NodeArray) Swap(index1, index2 int) {
	if index1 < 0 || index1 >= a.size || index2 < 0 || index2 >= a.size {
		return
	}
	a.items[index1], a.items[index2] = a.items[index2], a.items[index1]
}

func (a *NodeArray) getLoadFactor() float64 {
	return float64(a.size) / float64(len(a.items))
}

func (a *NodeArray) grow() {
	newCapacity := len(a.items) * growShrinkFactor
	a.resize(newCapacity)
}

func (a *NodeArray) shrink() {
	newCapacity := len(a.items) / growShrinkFactor
	a.resize(newCapacity)
}

func (a *NodeArray) resize(newCapacity int) {
	newArr := make([]*Node, newCapacity)
	for i := 0; i < a.size; i++ {
		newArr[i] = a.items[i]
	}
	a.items = newArr
}
