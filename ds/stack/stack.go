package stack

import "fmt"

type stack struct {
	data []int
}

type stackOps interface {
	createStack()
	push()
	pop()
}

func (s stack) createStack() *stack {
	//s.data = []int{}
	return &s
}

func (s *stack) push(n int) *stack {
	s.data = append(s.data, n)
	return s
}

func (s *stack) pop() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("Stack Empty")
	}
	element := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return element, nil
}
