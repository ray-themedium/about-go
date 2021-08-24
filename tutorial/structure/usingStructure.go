package main

import "fmt"

func main() {
	type S struct {
		X int
		Y int
		Z int
	}

	var s1 S
	fmt.Println(s1.Y, s1.Z)
	p1 := S{13, 14, 01}
	p2 := S{5, 6, 7}
	fmt.Println(p1)
	fmt.Println(p2)

	pSlice := [4]S{}
	pSlice[2] = p1
	pSlice[0] = p2
	fmt.Println(pSlice)
	p2 = S{1, 2, 3}
	fmt.Println(pSlice)
}
