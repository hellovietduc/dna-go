package hash_table

import (
	"errors"
)

type Node struct {
	key   int
	value int
	next  *Node
}

type HashTable struct {
	arr  []*Node
	size int
}

func NewHashTable() *HashTable {
	return &HashTable{
		arr: make([]*Node, 10),
	}
}

func (h *HashTable) Insert(key int, value int) {
	// same load factor as Java 10
	if h.getLoadFactor() >= 0.75 {
		h.rehash(true)
	}

	h.size++
	hash := h.getHash(key)
	node := h.arr[hash]

	// first node in the linked list
	if node == nil {
		node := &Node{
			key:   key,
			value: value,
		}
		h.arr[hash] = node
		return
	}

	// find last node in the linked list
	for node.next != nil {
		node = node.next
	}

	node.next = &Node{
		key:   key,
		value: value,
	}

}

func (h *HashTable) Search(key int) (int, error) {
	hash := h.getHash(key)
	node := h.arr[hash]

	if node == nil {
		return -1, errors.New("Key not found")
	}

	for node.key != key {
		node = node.next
	}

	return node.value, nil
}

func (h *HashTable) Delete(key int) (int, error) {
	if h.getLoadFactor() < 0.2 {
		h.rehash(false)
	}

	hash := h.getHash(key)
	node := h.arr[hash]

	if node == nil {
		return -1, errors.New("Key not found")
	}

	var prevBucket *Node
	for node.key != key {
		prevBucket = node
		node = node.next
	}

	if prevBucket == nil {
		// delete the first node in the linked list
		h.arr[hash] = node.next
	} else {
		prevBucket.next = node.next
	}

	// remove the reference from this node
	node.next = nil
	h.size--

	return node.value, nil
}

func (h *HashTable) Size() int {
	return h.size
}

func (h *HashTable) getLoadFactor() float64 {
	return float64(h.Size()) / float64(len(h.arr))
}

func (h *HashTable) getHash(key int) int {
	return key % len(h.arr)
}

func (h *HashTable) rehash(isGrowSize bool) {
	var newLength int
	if isGrowSize == true {
		newLength = len(h.arr) * 2
	} else {
		newLength = len(h.arr) / 2
	}

	newArr := make([]*Node, newLength)
	oldArr := h.arr
	h.arr = newArr
	h.size = 0

	for _, node := range oldArr {
		for node != nil {
			h.Insert(node.key, node.value)
			node = node.next
		}
	}
}
