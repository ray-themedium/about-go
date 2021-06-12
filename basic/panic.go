package main

import "fmt"

func divide(a, b int) float64 {
	if b == 0 {
		panic("b는 0일 수 없습니다.")
	}
	return float64(a / b)
}

func f() {
	fmt.Println("f 시작")
	defer func () {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g()
	fmt.Println("f 함수 끝")
}

func g() {
	fmt.Printf("9 / 3 = %f\n", divide(9, 3))
	fmt.Printf("9 / 3 = %f\n", divide(9, 0))
}

func main() {
	// divide(9, 3)
	// divide(9, 0)

	f()
	fmt.Println("프로그램이 계속 실행됨")
}