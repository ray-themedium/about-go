package main

import "fmt"

func main() {
	arr := []int{10, 20, 3, 1, 6, 9, 15, 8}

	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr) - i; j++ {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	fmt.Println(arr)
}
