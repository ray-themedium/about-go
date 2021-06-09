package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5) // 요소가 5개인 링 생성
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}
	fmt.Println()

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
}