package main

import "fmt"

func ExampleCapture() {
	i := 0

	f := func() {
		// 함수 리터럴에서 외부 변수를 내부 상태로 가져올 때 값 복사가 아닌
		// 인스턴스 참조로 가져오게 됨
		// 이 것을 캡쳐(capture)라고 함
		i += 10
	}

	i++
	f()
	fmt.Println(i)
}

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i, _ := range f {
		f[i] = func() {
			// i를 참조로 가져옴
			fmt.Println(i)
		}
	}

	for _, function := range f {
		function()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for _, function := range f {
		function()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
