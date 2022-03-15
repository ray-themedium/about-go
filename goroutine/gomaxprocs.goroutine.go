package main

import (
	"fmt"
	"runtime"
)

func main() {
	gomaxprocs := runtime.GOMAXPROCS(0)
	fmt.Println("GOMAXPROCS:", gomaxprocs)
}
