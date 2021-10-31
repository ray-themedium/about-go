package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// 모든 노드를 재귀 호출로 방문
func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

//
func create(n int) *Tree {
	var t *Tree
	// rand.Seed(time.Now().Unix())
	for _, v := range []int{7, 3, 9, 2, 1, 8, 6, 4, 10} {
		t = insert(t, v)
		traverse(t)
		fmt.Println()
	}
	// for i := 0; i < 2*n; i++ {
	// 	temp := rand.Intn(n * 2)
	// 	t = insert(t, temp)
	// }
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	tree := create(10)
	fmt.Println("root:", tree.Value)
	// traverse(tree)
	// fmt.Println()
	// tree = insert(tree, -10)
	// tree = insert(tree, -2)
	// traverse(tree)
	// fmt.Println()
	// fmt.Println("root:", tree.Value)
}
