package util

type Queue []interface{}

const (
	QueueInitSize = 5
)

func NewQueue() Queue {
	return make([]interface{}, 0, QueueInitSize)
}

func (q *Queue) Enqueue(elem interface{}) {
	*q = append(*q, elem)
}

func (q *Queue) Dequeue() interface{} {
	elem := (*q)[0]
	*q = (*q)[1:len(*q)]
	return elem
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}

func (q *Queue) Size() int {
	return len(*q)
}
