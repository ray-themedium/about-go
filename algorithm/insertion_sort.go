package main

import "fmt"

func main() {
	arr := []int{10, 20, 3, 1, 6, 9, 15, 8}

	for i := 1; i < len(arr); i++ {
		val := arr[i]
		prevIdx := i - 1
		for (prevIdx >= 0) && (val < arr[prevIdx]) {
			arr[prevIdx + 1] = arr[prevIdx]
			prevIdx--
		}
		arr[prevIdx + 1] = val
	}
	fmt.Println(arr)
}

