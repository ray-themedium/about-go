package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func ExampleGoroutine() {
	defer wg.Done()
	temp := 0
	for i := 0; i <= 1000000000; i++ {
		temp += i
	}
	fmt.Println(temp)
}

func main() {
	// 작업 개수 설정
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go ExampleGoroutine()
	}
	// 모든 작업을 기다림
	wg.Wait()
	fmt.Println("Finish")
}
