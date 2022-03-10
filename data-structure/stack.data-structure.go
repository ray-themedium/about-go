package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v    *list.List
	size int
}

func (stack *Stack) Push(value interface{}) {
	stack.v.PushBack(value)
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	if stack.v.Len() == 0 {
		return nil
	}
	return stack.v.Remove(stack.v.Back())
}

func NewStack() *Stack {
	return &Stack{list.New(), 0}
}

func main() {
	stack := NewStack()

	for i := 0; i < 5; i++ {
		stack.Push(i)
	}

	for v := stack.Pop(); v != nil; v = stack.Pop() {
		fmt.Println(v)
	}
}
