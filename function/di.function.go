package main

import (
	"fmt"
	"os"
)

type Writer func(string)

// DI
func WriteLogic(writer Writer) {
	writer("logic")
}

func main() {
	f, err := os.Create("writer-test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	WriteLogic(func(s string) { fmt.Fprintln(f, s) })
}
