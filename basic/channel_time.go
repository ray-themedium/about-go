package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second) // 1초 간격 시그널이 오는 채널
	terminate := time.After(10 * time.Second) // 10초 이후 시그널이 오는 채널

	for {
		select { // 하나만 통과되면 나머지는 무시됨
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square(&wg, ch)

	// 채널에 데이터 담기
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	wg.Wait()
}