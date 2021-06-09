package main

import (
	"container/list"
	"fmt"
)

func main() {
	li := list.New() // 리스트 생성
	e4 := li.PushBack(4)

	e1 := li.PushFront(1)

	li.InsertBefore(3, e4)

	li.InsertAfter(2, e1)

	for e := li.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	
	fmt.Println()

	for e := li.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}