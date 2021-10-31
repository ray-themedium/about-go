package main

import "fmt"

const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunction(i, size int) int {
	return i % size
}

func insert(table *HashTable, value int) int {
	index := hashFunction(value, table.Size)
	element := Node{value, table.Table[index]}
	table.Table[index] = &element
	return index
}

func traverse(table *HashTable) {
	for k := range table.Table {
		if table.Table[k] != nil {
			t := table.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func lookup(table *HashTable, value int) bool {
	index := hashFunction(value, SIZE)
	node := table.Table[index]
	if node == nil {
		return false
	}
	current := node
	for current.Next != nil {
		if value == current.Value {
			return true
		}
		current = current.Next
	}
	return false
}

func main() {
	table := make(map[int]*Node, SIZE)
	hashTable := &HashTable{table, SIZE}
	for i := 0; i < 100; i++ {
		insert(hashTable, i)
	}
	traverse(hashTable)

	fmt.Println(lookup(hashTable, 50))
}
