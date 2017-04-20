package util

// Queue is simply a slice of any type
type Queue []interface{}

const (
	queueInitSize = 5
)

// NewQueue creates an empty queue
func NewQueue() Queue {
	return make([]interface{}, 0, queueInitSize)
}

// Enqueue adds element to the queue
func (q *Queue) Enqueue(elem interface{}) {
	*q = append(*q, elem)
}

// Dequeue removes element from the queue
func (q *Queue) Dequeue() interface{} {
	elem := (*q)[0]
	*q = (*q)[1:len(*q)]
	return elem
}

// Empty checks if queue is empty
func (q *Queue) Empty() bool {
	return q.Size() == 0
}

// Size returns size of queue
func (q *Queue) Size() int {
	return len(*q)
}
