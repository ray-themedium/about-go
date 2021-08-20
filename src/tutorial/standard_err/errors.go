// cla.go float 타입 검사하는 코드를 추가한 것
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one more floats.")
		os.Exit(1)
	}

	args := os.Args // a b c -1 100 -20 a
	var err error = errors.New("An Error")
	k := 1
	var n float64

	for err != nil {
		if k >= len(args) {
			fmt.Println("None of the args is a float")
			return
		}
		n, err = strconv.ParseFloat(args[k], 64)
		k++
	}
	min, max := n, n

	for i := 2; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err == nil {
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
