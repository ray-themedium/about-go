package main

import "fmt"

func main() {
	example1()
	example2()
}

// 요소 추가하기1
func example1() {
	slice := []int{1,2,3,4,5,6}
	idx := 2
	value := 100

	slice = append(slice[:idx], append([]int{value}, slice[idx:]...)...)
	fmt.Println(slice)
}

// 요소 추가하기2
func example2() {
	slice := []int{1,2,3,4,5,6}
	idx := 2
	value := 100

	slice = append(slice, 0) // 공간 하나 확보
	copy(slice[idx+1:], slice[idx:]) // 삽입할 위치를 제외한 뒤쪽 요소들을 복사
	slice[idx] = value
	fmt.Println(slice)
}