package linked_list

const (
	// NodeNotFound is a constant indicates that
	// a node cannot be found in the Linked List.
	NodeNotFound = 0
)

// Node is a struct representing a data node
// in the Linked List.
type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

// Front returns the value of the previous node
// to this node.
func (n *Node) Front() int {
	if n.Prev == nil {
		return NodeNotFound
	}
	return n.Prev.Value
}

// Back returns the value of the next node
// to this node.
func (n *Node) Back() int {
	if n.Next == nil {
		return NodeNotFound
	}
	return n.Next.Value
}

// LinkedList is a struct which provides
// functions to manipulate the Linked List.
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// Size returns the number of nodes being stored
// in the Linked List.
func (l *LinkedList) Size() int {
	return l.size
}

// IsEmpty returns whether the Linked List has
// any nodes or not.
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// ValueAt returns the value of the node
// at the specified position in the Linked List.
func (l *LinkedList) ValueAt(pos int) int {
	if pos < 0 || pos >= l.size {
		return NodeNotFound
	}

	node := l.head
	for i := 0; i < pos; i++ {
		node = node.Next
	}
	return node.Value
}

// PushFront sets a new node as the head of the Linked List.
func (l *LinkedList) PushFront(value int) {
	node := &Node{
		Value: value,
		Prev:  nil,
		Next:  l.head,
	}
	l.head = node
	l.size++
}

// PushBack sets a new node as the tail of the Linked List.
func (l *LinkedList) PushBack(value int) {
	if l.IsEmpty() {
		l.PushFront(value)
		return
	}

	node := &Node{
		Value: value,
		Prev:  l.tail,
		Next:  nil,
	}

	if l.tail == nil {
		node.Prev = l.head
		l.head.Next = node
	}

	l.tail = node
	l.size++
}

// Insert adds a node at the specified position in the
// Linked List and returns whether the operation succeeds
// or not.
func (l *LinkedList) Insert(value int, pos int) bool {
	if pos < 0 || pos >= l.size {
		return false
	}

	if pos == 0 {
		l.PushFront(value)
		return true
	}

	node := l.head
	for i := 0; i < pos; i++ {
		node = node.Next
	}

	prevNode := node.Prev
	nextNode := node.Next

	newNode := &Node{
		Value: value,
		Prev:  prevNode,
		Next:  node,
	}

	node.Prev = newNode
	if prevNode != nil {
		prevNode.Next = newNode
	}

	if nextNode != nil {
		nextNode.Prev = newNode
	}

	l.size++
	return true
}

// PopFront removes and returns the value
// of the head of the Linked List.
func (l *LinkedList) PopFront() int {
	oldHead := l.head
	if oldHead == nil {
		return NodeNotFound
	}

	l.head = oldHead.Next
	l.head.Prev = nil

	oldHead.Next = nil
	oldHead.Prev = nil

	l.size--
	return oldHead.Value
}

// PopBack removes and returns the value
// of the tail of the Linked List.
func (l *LinkedList) PopBack() int {
	oldTail := l.tail
	if oldTail == nil {
		return NodeNotFound
	}

	l.tail = oldTail.Prev
	l.tail.Next = nil

	oldTail.Next = nil
	oldTail.Prev = nil

	l.size--
	return oldTail.Value
}

// RemoveAt removes a node at the specified position
// in the Linked List and returns whether the operation
// succeeds or not.
func (l *LinkedList) RemoveAt(pos int) bool {
	if pos < 0 || pos >= l.size {
		return false
	}

	if pos == 0 {
		l.PopFront()
		return true
	}

	if pos == l.size-1 {
		l.PopBack()
		return true
	}

	node := l.head
	for i := 0; i < pos; i++ {
		node = node.Next
	}

	prevNode := node.Prev
	nextNode := node.Next

	node.Prev = nil
	node.Next = nil

	if prevNode != nil {
		prevNode.Next = nextNode
	}
	if nextNode != nil {
		nextNode.Prev = prevNode
	}

	l.size--
	return true
}
