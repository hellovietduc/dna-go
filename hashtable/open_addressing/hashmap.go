package open_addressing

const (
	// KeyNotFound is a constant indicates that
	// a key cannot be found in the HashMap.
	KeyNotFound      = 0
	defaultCapacity  = 8
	defaultPrime     = 7
	upperLoadFactor  = 0.75
	lowerLoadFactor  = 0.2
	growShrinkFactor = 2
)

type node struct {
	key     int
	value   int
	deleted bool
}

// HashMap is a struct for storing key-value data.
type HashMap struct {
	items []*node
	size  int
	prime int
}

// NewHashMap creates a new HashMap instance.
func NewHashMap() *HashMap {
	return &HashMap{
		items: make([]*node, defaultCapacity),
		prime: defaultPrime,
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

	h.size++

	tries := 1
	for true {
		index := h.hash(key, tries)
		curNode := h.items[index]

		if curNode == nil || curNode.deleted == true {
			h.items[index] = newNode
			return
		}

		if curNode.key == key {
			curNode.value = value
			return
		}

		tries++
	}
}

// Search returns the value for the given key.
// If the key cannot be found, it returns KeyNotFound.
func (h *HashMap) Search(key int) int {
	tries := 1
	for true {
		index := h.hash(key, tries)
		curNode := h.items[index]

		if curNode == nil {
			return KeyNotFound
		}

		if curNode.key == key && curNode.deleted == false {
			return curNode.value
		}

		tries++
	}

	return KeyNotFound
}

// Delete removes a key from the HashMap.
func (h *HashMap) Delete(key int) {
	if h.getLoadFactor() < lowerLoadFactor {
		h.shrink()
	}

	h.size--

	tries := 1
	for true {
		index := h.hash(key, tries)
		curNode := h.items[index]

		if curNode == nil {
			return
		}

		if curNode.key == key {
			curNode.deleted = true
			return
		}

		tries++
	}
}

func (h *HashMap) capacity() int {
	return len(h.items)
}

func (h *HashMap) getLoadFactor() float64 {
	return float64(h.size) / float64(len(h.items))
}

func (h *HashMap) hash(key int, tries int) int {
	originIndex := key % len(h.items)
	stepSize := h.prime - (key % h.prime)
	return (originIndex + tries*stepSize) % len(h.items)
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
	h.prime = h.nearestPrime(newLength)

	for _, curNode := range oldArr {
		if curNode != nil && curNode.deleted == false {
			h.Insert(curNode.key, curNode.value)
		}
	}
}

func (h *HashMap) nearestPrime(num int) int {
	prime := num - 1
	for h.isPrime(prime) == false {
		prime = prime - 1
	}
	return prime
}

func (h *HashMap) isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
