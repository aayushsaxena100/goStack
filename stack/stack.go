package stack

//Stack is a data structure that follows LIFO
type Stack struct {
	capacity int   // the max amount of elements the stack can hold
	a        []int // the underlying slice to implement stack
	Size     int   // the current size of the stack
}

//New is the stack initializer
func New(cap int) Stack {
	return Stack{capacity: cap}
}

//FullError means the stack is already full and push operation is performed
type FullError struct{}

func (s *FullError) Error() string {
	return "The stack is already full"
}

//EmptyError means the stack is already empty and pop or top operation is performed
type EmptyError struct{}

func (s *EmptyError) Error() string {
	return "The stack is already empty"
}

//Push inserts an element
func (s *Stack) Push(val int) error {
	if len(s.a) == s.capacity {
		return &FullError{}
	}
	s.a = append(s.a, val)
	s.Size = s.Size + 1
	return nil
}

//Pop removes the latest inserted element
func (s *Stack) Pop() error {
	if len(s.a) == 0 {
		return &EmptyError{}
	}
	s.a = s.a[:len(s.a)-1]
	s.Size = s.Size - 1
	return nil
}

//Top returns the latest inserted element
func (s *Stack) Top() (int, error) {
	if len(s.a) == 0 {
		return 0, &EmptyError{}
	}
	return s.a[len(s.a)-1], nil
}
