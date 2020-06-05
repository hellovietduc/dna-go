package stack

// Empty is a constant indicates that
// the returned value is empty.
const Empty 0

// Stack is a struct for storing values
// which follow LIFO order.
type Stack struct {
	items  []int
	size int
}

// NewStack creates a new Stack instance
// with specified capacity.
func NewStack(capacity int) *Stack {
	return &Stack{
		items: make([]int, capacity),
	}
}

// Push adds a new value to the Stack
// and returns true if the value is
// added successfully.
func (s *Stack) Push(value int) bool {
	if s.IsFull() == true {
		return false
	}
	s.items[s.size] = value
	s.size++
	return true
}

// Pop removes a value from the Stack.
// It returns Empty if there's no values
// to pop out.
func (s *Stack) Pop() int {
	if s.IsEmpty() == true {
		return Empty
	}
	value := s.items[s.size-1]
	s.size--
	return value
}

// IsEmpty returns true if the Stack
// has no values, else it returns false.
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// IsFull returns true if the Stack
// has stored maximum number of values,
// else it returns false.
func (s *Stack) IsFull() bool {
	return s.size > len(s.items)-1
}
