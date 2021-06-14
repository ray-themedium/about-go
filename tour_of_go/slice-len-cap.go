package main

import "fmt"

func printSlice(s []int) {
	fmt.Println(s)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13} // len -> 6, cap -> 6
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:] // 앞에서부터 자르면 용량도 줄어듬
	printSlice(s)
}