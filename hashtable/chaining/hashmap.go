package chaining

const (
	// KeyNotFound is a constant indicates that
	// a key cannot be found in the HashMap.
	KeyNotFound      = 0
	defaultCapacity  = 8
	upperLoadFactor  = 0.75
	lowerLoadFactor  = 0.2
	growShrinkFactor = 2
)

type node struct {
	key   int
	value int
	next  *node
	prev  *node
}

// HashMap is a struct for storing key-value data.
type HashMap struct {
	items []*node
	size  int
}

// NewHashMap creates a new HashMap instance.
func NewHashMap() *HashMap {
	return &HashMap{
		items: make([]*node, defaultCapacity),
	}
}

// Insert adds a new pair of key-value to the HashMap.
// If the key already exists, it overwrites the key's value.
func (h *HashMap) Insert(key int, value int) {
	if h.getLoadFactor() >= upperLoadFactor {
		h.grow()
	}

	newNode := &node{
		key:   key,
		value: value,
	}

	foundNode, index := h.searchNode(key)
	if foundNode != nil {
		// overwrite value of this key
		foundNode.value = value
		return
	}

	h.size++
	if curNode := h.items[index]; curNode == nil {
		// first node in the linked list
		h.items[index] = newNode
	} else {
		// last node in the linked list
		for curNode.next != nil {
			curNode = curNode.next
		}
		curNode.next = newNode
		newNode.prev = curNode
	}
}

// Delete removes a key from the HashMap.
func (h *HashMap) Delete(key int) {
	if h.getLoadFactor() < lowerLoadFactor {
		h.shrink()
	}

	foundNode, index := h.searchNode(key)
	if foundNode == nil {
		return
	}

	h.size--

	if foundNode.prev == nil {
		// first node in the linked list
		h.items[index] = foundNode.next
		return
	}

	if foundNode.next == nil {
		// last node in the linked list
		foundNode.prev.next = nil
	}

	prevNode := foundNode.prev
	nextNode := foundNode.next
	prevNode.next = nextNode
	nextNode.prev = prevNode
}

// Search returns the value for the given key.
// If the key cannot be found, it returns KeyNotFound.
func (h *HashMap) Search(key int) int {
	if foundNode, _ := h.searchNode(key); foundNode != nil {
		return foundNode.value
	}
	return KeyNotFound
}

func (h *HashMap) searchNode(key int) (*node, int) {
	index := h.hash(key)
	curNode := h.items[index]

	for curNode != nil {
		if curNode.key == key {
			return curNode, index
		}

		curNode = curNode.next
	}

	return nil, index
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

func (h *HashMap) grow() {
	newLength := len(h.items) * growShrinkFactor
	h.rehash(newLength)
}

func (h *HashMap) shrink() {
	newLength := len(h.items) / growShrinkFactor
	h.rehash(newLength)
}

func (h *HashMap) rehash(newLength int) {
	newArr := make([]*node, newLength)
	oldArr := h.items
	h.items = newArr
	h.size = 0

	for _, curNode := range oldArr {
		for curNode != nil {
			h.Insert(curNode.key, curNode.value)
			curNode = curNode.next
		}
	}
}
