package main

import (
	"container/heap"
	"fmt"
)

type heapFloat32 []float32

// container/heap에 필요한 메서드 구현
func (n *heapFloat32) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new
	return x
}

func (n *heapFloat32) Push(x interface{}) {
	*n = append(*n, x.(float32))
}

// sort.Interface에 필요한 메서드 구현
func (n heapFloat32) Len() int {
	return len(n)
}

func (n heapFloat32) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapFloat32) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func main() {
	// heap 구조는 일종의 트리이기 때문에 출력해서 확인할 때 첫 번째 요소를 제외하고 순서가 달라보일 수 있다.

	myHeap := &heapFloat32{1.2, 2.1, 3.1, -100.1}
	heap.Init(myHeap) // 힙으로 초기화하여 슬라이스를 힙 구조로 변경
	size := len(*myHeap)
	fmt.Printf("Heap size: %d\n", size)
	fmt.Printf("%v\n", myHeap)

	myHeap.Push(float32(-99.2))
	myHeap.Push(float32(0.2))
	fmt.Printf("Heap size: %d\n", len(*myHeap))
	fmt.Printf("%v\n", myHeap)
	heap.Init(myHeap) // 변경된 리스트를 다시 힙 구조로 변경
	fmt.Printf("%v\n", myHeap)
}
