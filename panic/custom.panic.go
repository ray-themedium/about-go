package main

import "fmt"

type MyError struct {
	len      int
	required int
}

func ExampleCustomPanic() {
	defer func() {
		if r, ok := recover().(MyError); ok {
			fmt.Printf("r is MyError Type r: %v\n", r)
		}
	}()
	panic(MyError{5, 10})
}

func main() {
	ExampleCustomPanic()
}
