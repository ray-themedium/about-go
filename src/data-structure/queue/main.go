package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue adds a value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes a value at the front
// and RETURNs the removed value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	queue := Queue{}

	fmt.Println(queue)
	queue.Enqueue(100)
	queue.Enqueue(200)
	queue.Enqueue(300)
	fmt.Println(queue)

	value := queue.Dequeue()
	fmt.Println(value)
	fmt.Println(queue)

}
