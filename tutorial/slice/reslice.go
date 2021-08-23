package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)

	// 같은 배열을 참조하고 있기 때문에 원본 슬라이스에도 영향을 미친다.
	reSlice[0] = -100
	reSlice[1] = 12345

	fmt.Println(s1)
	fmt.Println(reSlice)
}
