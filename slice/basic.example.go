package main

import "fmt"

// 슬라이스 복사 예제
func ExampleCopySlice() {
	fmt.Println("start ExampleCopySlice()")
	defer fmt.Println("end ExampleCopySlice()\n")
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 3, 10)
	c := make([]int, 10)

	// 복사된 요소의 개수가 리턴
	// 실제 복사되는 요소는 슬라이스 길이 중 가장 작은 개수만큼 복사됨
	cnt1 := copy(b, a)
	cnt2 := copy(c, a)

	fmt.Println(cnt1, b)
	fmt.Println(cnt2, c)

	// 동일한 크기에 복사할 때
	d := make([]int, len(a))
	cnt3 := copy(d, a)
	fmt.Println(cnt3, d)
}

// 슬라이스 요소 삭제 예제
func ExampleDeleteArgument() {
	fmt.Println("start ExampleDeleteArgument()")
	defer fmt.Println("end ExampleDeleteArgument()\n")
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a)
	// 삭제할 요소의 인덱스
	index := 2
	a = append(a[:index], a[index+1:]...)
	fmt.Println(a)
}

// 슬라이스 요소 추가 예제
func ExampleInsertArgument() {
	fmt.Println("start ExampleInsertArgument()")
	defer fmt.Println("end ExampleInsertArgument()\n")
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(a)
	// 삽입할 위치의 인덱스
	index := 2
	a = append(a[:index], append([]int{100}, a[index:]...)...)
	fmt.Println(a)

	// 불필요한 메모리 사용 없이 추가
	b := []int{6, 7, 8, 9, 10}
	index = 2
	b = append(b, 0)
	copy(b[index+1:], b[index:])
	b[index] = 100
	fmt.Println(b)
}

func main() {
	ExampleCopySlice()
	ExampleDeleteArgument()
	ExampleInsertArgument()
}
