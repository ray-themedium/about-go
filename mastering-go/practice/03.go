package main

import "fmt"

const (
	n1 int = 4 << iota
	n2
	n3
	n4
)

func main() {
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
}
