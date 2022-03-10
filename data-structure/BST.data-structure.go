package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Left  *Node
	Value int
	Right *Node
}

type BST struct {
	Root *Node
}

func (tree *BST) PreorderTraverse(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.Value, " ")
	tree.PreorderTraverse(n.Left)
	tree.PreorderTraverse(n.Right)
}

func (tree *BST) InorderTraverse(n *Node) {
	if n == nil {
		return
	}
	tree.InorderTraverse(n.Left)
	fmt.Println(n.Value, " ")
	tree.InorderTraverse(n.Right)
}

func (tree *BST) PostorderTraverse(n *Node) {
	if n == nil {
		return
	}
	tree.PostorderTraverse(n.Left)
	tree.PostorderTraverse(n.Right)
	fmt.Println(n.Value, " ")
}

func (tree *BST) Add(value int) bool {
	if tree.Root == nil {
		tree.Root = &Node{nil, value, nil}
		return true
	}

	current := tree.Root
	newNode := &Node{nil, value, nil}
	for true {
		if value == current.Value {
			return false
		}
		if value < current.Value {
			if current.Left != nil {
				current = current.Left
			} else {
				current.Left = newNode
				break
			}
		} else {
			if current.Right != nil {
				current = current.Right
			} else {
				current.Right = newNode
				break
			}
		}
	}
	return true
}

func NewBST() *BST {
	return &BST{}
}

func main() {
	bst := NewBST()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		bst.Add(rand.Intn(100))
	}

	bst.PreorderTraverse(bst.Root)
	fmt.Println()
	bst.InorderTraverse(bst.Root)
	fmt.Println()
	bst.PostorderTraverse(bst.Root)
}
