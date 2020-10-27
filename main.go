package main

import (
	stack "abstractDataType/stack"
	"fmt"
)

func main() {
	s := stack.New(4)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	size := s.Size
	for i := 0; i < size; i++ {
		top, _ := s.Top()
		fmt.Println(top)
		s.Pop()
	}
}
