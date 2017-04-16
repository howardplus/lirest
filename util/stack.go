package util

type Stack []interface{}

const (
	StackInitSize = 5
)

func NewStack() Stack {
	return make([]interface{}, 0, StackInitSize)
}

func (s *Stack) Push(elem interface{}) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() interface{} {
	elem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return elem
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(*s)
}
