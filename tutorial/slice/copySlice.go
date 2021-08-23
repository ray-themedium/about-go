package main

import "fmt"

func main() {
	a6 := []int{1, 2, 3, 4, 5, 6}
	a4 := []int{7, 8, 9}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	copy(a6, a4) // 같은 위치의 인덱스의 요소들이 그 자리로 복사되므로 a6: [7 8 9 4 5 6] 가 된다.
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()

	b3 := []int{7, 8, 9}
	b6 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("b3:", b3)
	fmt.Println("b6:", b6)

	copy(b3, b6) // 같은 위치의 인덱스의 요소들이 그 자리로 복사되므로 b3: [1, 2, 3] 가 된다.
	fmt.Println("b3:", b3)
	fmt.Println("b6:", b6)
	fmt.Println()

	c6 := []int{1, 2, 3, 4, 5, 6}
	c4 := [4]int{7, 8, 9, 10} // 배열
	fmt.Println("c6:", c6)
	fmt.Println("c4:", c4)

	copy(c6, c4[0:]) // 슬라이스로 변환 후 카피 -> c6: [1, 2, 3] 가 된다.
	fmt.Println("c6:", c6)
	fmt.Println("c4:", c4)
	fmt.Println()

	d5 := [5]int{5, -5, 5, -5, 5}
	d7 := []int{7, -7, 7, -7, 7, -7, 7}
	fmt.Println("d5:", d5)
	fmt.Println("d7:", d7)

	copy(d5[0:], d7) // 배열을 슬라이스로 변환 후 카피를 받음
	fmt.Println("d5:", d5)
	fmt.Println("d7:", d7)
	// 배열은 용량이 정해져 있고 슬라이스로 변환하더라도 같은 공간을 가리켜 d5: [7 -7 7 -7 7] 이 된다.
}
