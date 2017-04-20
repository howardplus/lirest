package util

// Stack is a slice of any type
type Stack []interface{}

const (
	stackInitSize = 5
)

// NewStack
// create an empty stack
func NewStack() Stack {
	return make([]interface{}, 0, stackInitSize)
}

// Push
// add element to stack
func (s *Stack) Push(elem interface{}) {
	*s = append(*s, elem)
}

// Pop
// remove element from stack
func (s *Stack) Pop() interface{} {
	elem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return elem
}

// Empty
// check if stack is empty
func (s *Stack) Empty() bool {
	return s.Size() == 0
}

// Size
// return size of stack
func (s *Stack) Size() int {
	return len(*s)
}
