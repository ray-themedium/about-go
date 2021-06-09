package main

import "fmt"

func main() {
	arr := []int{10, 20, 3, 1, 6, 9, 15, 8}

	for i := 0; i < len(arr) - 1; i++ {
		targetIdx := i
		for j := 1 + i; j < len(arr); j++ {
			if arr[j] < arr[targetIdx] {
				targetIdx = j
			}
		}
		arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
	}
	fmt.Println(arr)
}

