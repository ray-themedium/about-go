package main

import "fmt"

type node struct {
	data int64
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) printListData() {
	toPrint := l.head
	if toPrint == nil {
		return
	}
	for toPrint.next != nil {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
	}
	fmt.Printf("%d\n", toPrint.data)
}

func (l *linkedList) deleteWithValue(value int64) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	myList.prepend(&node{data: 50})
	myList.printListData()
	// node2 := &node{data: 42}
	// node3 := &node{data: 65}
	// node4 := &node{data: 77}
	// node5 := &node{data: 88}
	// node6 := &node{data: 15}
	// myList.prepend(node1)
	// myList.prepend(node2)
	// myList.prepend(node3)
	// myList.prepend(node4)
	// myList.prepend(node5)
	// myList.prepend(node6)
	// myList.printListData()
	// myList.deleteWithValue(77)
	// myList.printListData()
	// myList.deleteWithValue(77)
	// myList.printListData()
}
