package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func add(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
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

func lookup(t *Node, v int) bool {
	if root == nil {
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookup(t.Next, v)
}

func getSize(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list")
		return 0
	}
	count := 0
	for t != nil {
		count++
		t = t.Next
	}
	return count
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
	fmt.Println(lookup(root, 3))
	fmt.Println(lookup(root, 100))
}
