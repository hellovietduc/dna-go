package queue

// Empty is a constant indicates that
// the returned value is empty.
const Empty = 0

// Queue is a struct for storing values
// which follow FIFO order.
type Queue struct {
	items []int
	front int
	rear  int
}

// NewQueue creates a new Queue instance
// with specified capacity.
func NewQueue(capacity int) *Queue {
	return &Queue{
		items: make([]int, capacity+1),
	}
}

// Enqueue adds a new value to the Queue
// and returns if the value is added
// successfully.
func (q *Queue) Enqueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.items[q.rear] = value
	q.rear = (q.rear + 1) % len(q.items)
	return true
}

// Dequeue removes a value and returns its
// value. It returns Empty if the Queue
// is empty.
func (q *Queue) Dequeue() int {
	if q.IsEmpty() == true {
		return Empty
	}
	value := q.items[q.front]
	q.front = (q.front + 1) % len(q.items)
	return value
}

// IsEmpty returns true if the Queue has
// no values, else it returns false.
func (q *Queue) IsEmpty() bool {
	return q.front == q.rear
}

// IsFull returns true if the Queue has stored
// maximum number of values, else it returns false.
func (q *Queue) IsFull() bool {
	return (q.rear+1)%len(q.items) == q.front
}
