package tests

import (
	Stack "abstractDataType/stack"
	"testing"
)

func TestPush(t *testing.T) {
	s := Stack.New(4)
	s.Push(1)
	if top, _ := s.Top(); top != 1 {
		t.Error("Pushed element not on top")
	}

	s.Push(2)
	s.Push(3)
	s.Push(4)
	err := s.Push(5)

	if err == nil {
		t.Error("Stack is not throwing error on exceeding capacity")
	}
}

func TestPop(t *testing.T) {
	s := Stack.New(4)
	s.Push(1)
	s.Push(2)

	s.Pop()

	if top, _ := s.Top(); top != 1 {
		t.Error("Expected {}, but the top of stack is {}", 1, top)
	}

	s.Pop()

	if err := s.Pop(); err == nil {
		t.Error("Performing Pop on empty stack not throwing error")
	}
}

func TestTop(t *testing.T) {
	s := Stack.New(4)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	if top, _ := s.Top(); top != 4 {
		t.Error("Expected 4, but the top of stack is ", top)
	}

	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	if _, err := s.Top(); err == nil {
		t.Error("Error not thrown on looking for top of an empty stack")
	}
}
