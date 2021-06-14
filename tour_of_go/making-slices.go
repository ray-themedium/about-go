package main

import "fmt"

func main() {
	a := make([]int, 5)	
	printSlice("a", a)

	b := make([]int, 0, 5) // len -> 0 이라서 값이 할당되지 않았음
	printSlice("b", b)

	c := b[:2] // [0, 0]
	printSlice("c", c)

	d := c[2:5] // [0, 0, 0]	
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}