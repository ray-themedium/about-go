package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square2(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <- ch // 데이터가 없다면 대기한다.

	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}

func square2(wg *sync.WaitGroup, ch chan int) {
	// 채널에 데이터가 들어오기를 계속 기다리고 있다는 것
	for n := range ch { // for range 구문을 사용하면 채널에서 데이터를 계속 기다릴 수 있음
		fmt.Printf("Square: %d\n", n * n)
		time.Sleep(time.Second)
	}
	wg.Done()
}