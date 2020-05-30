package hashtable

const (
	KeyNotFound     = 0
	defaultCapacity = 8
	upperLoadFactor = 0.75
	lowerLoadFactor = 0.2
)

type Node struct {
	key   int
	value int
	next  *Node
}

type HashMap struct {
	items []*Node
	size  int
}

func NewHashMap() *HashMap {
	return &HashMap{
		items: make([]*Node, defaultCapacity),
	}
}

func (h *HashMap) Insert(key int, value int) {
	if h.getLoadFactor() >= upperLoadFactor {
		h.rehash(true)
	}

	newNode := &Node{
		key:   key,
		value: value,
	}

	h.size++
	index := h.hash(key)
	node := h.items[index]

	if node == nil {
		// first node in the linked list
		h.items[index] = newNode
		return
	}

	for true {
		if node.key == key {
			// found key to overwrite value
			node.value = value
			return
		}

		if node.next == nil {
			// or reach the last node
			break
		}

		node = node.next
	}

	node.next = newNode
}

func (h *HashMap) Search(key int) int {
	index := h.hash(key)
	node := h.items[index]

	if node == nil {
		return KeyNotFound
	}

	for node.key != key {
		if node.next == nil {
			// reach the last node
			return KeyNotFound
		}

		node = node.next
	}

	return node.value
}

func (h *HashMap) Delete(key int) {
	if h.getLoadFactor() < lowerLoadFactor {
		h.rehash(false)
	}

	index := h.hash(key)
	node := h.items[index]

	if node == nil {
		return
	}

	var prevNode *Node
	for node.key != key {
		if node.next == nil {
			// reach the last node
			return
		}

		prevNode = node
		node = node.next
	}

	if prevNode == nil {
		// delete the first node in the linked list
		h.items[index] = node.next
	} else {
		prevNode.next = node.next
	}

	// remove the reference from this node
	node.next = nil
	h.size--
}

func (h *HashMap) capacity() int {
	return len(h.items)
}

func (h *HashMap) getLoadFactor() float64 {
	return float64(h.size) / float64(len(h.items))
}

func (h *HashMap) hash(key int) int {
	return key % len(h.items)
}

func (h *HashMap) rehash(isGrowSize bool) {
	currLength := len(h.items)
	var newLength int
	if isGrowSize == true {
		newLength = currLength * 2
	} else {
		newLength = currLength / 2
	}

	newArr := make([]*Node, newLength)
	oldArr := h.items
	h.items = newArr
	h.size = 0

	for _, node := range oldArr {
		for node != nil {
			h.Insert(node.key, node.value)
			node = node.next
		}
	}
}
