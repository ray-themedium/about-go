package main

import "fmt"

type Stack struct {
	items []int
}

// Push will add a value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	stack := Stack{}

	fmt.Println(stack)
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)
	fmt.Println(stack)

	value := stack.Pop()
	fmt.Println(value)
	fmt.Println(stack)
}
