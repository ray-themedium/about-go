package main

import "fmt"

type Node struct {
	Prev  *Node
	Value int
	Next  *Node
}

var root = new(Node)

func add(t *Node, v int) int {
	if root == nil {
		root = &Node{nil, v, nil}
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}
	if t.Next == nil {
		temp := t
		t.Next = &Node{temp, v, nil}
		return -2
	}
	return add(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func reverseTraverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}
	tail := t
	for t != nil {
		tail = t
		t = t.Next
	}

	for tail.Prev != nil {
		fmt.Printf("%d -> ", tail.Value)
		tail = tail.Prev
	}
	fmt.Printf("%d -> ", tail.Value)
	fmt.Println()
}

func main() {
	root = nil
	fmt.Println(root)
	traverse(root)
	add(root, 1)
	add(root, 2)
	add(root, 3)
	add(root, 4)
	add(root, 5)
	traverse(root)
	add(root, 3)
	add(root, 9)
	add(root, 8)
	add(root, 7)
	traverse(root)
	reverseTraverse(root)
}
