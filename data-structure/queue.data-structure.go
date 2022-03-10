package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(value interface{}) {
	q.v.PushBack(value)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue()

	for i := 0; i < 5; i++ {
		queue.Push(i)
	}

	for v := queue.Pop(); v != nil; v = queue.Pop() {
		fmt.Println(v)
	}
}
