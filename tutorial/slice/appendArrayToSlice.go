package main

import "fmt"

func printSlice(sl ...[]int) {
	for _, s := range sl {
		fmt.Println(s)
	}
}

func main() {
	a := []int{1, 2, 3}
	b := [3]int{4, 5, 6}
	ref := b[:]
	fmt.Println("Existing Array:\t", ref)

	c := append(a, ref...)
	fmt.Println("New Slice:\t", c)
	a = append(a, ref...)
	fmt.Println("Existing Slice:\t", a)
	a = append(a, a...)
	fmt.Println("a+a:\t\t", a)

	println()
	printSlice(a, ref, c)
	fmt.Println(b)
}
