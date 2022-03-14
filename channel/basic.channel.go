package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	// 채널 큐에서 꺼내옴
	// 채널에 메세지가 없으면 대기함
	// 채널 버퍼가 0이라서 대기하는 것임
	// n := <-ch

	// for문을 사용하면 채널의 데이터를 계속 기다릴 수 있음
	// 채널을 닫아줘야 기다리지 않음
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// 버퍼 크기가 2인 체널을 생성
	// ch := make(chan int, 2)
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	// 채널을 강제로 닫아주는 역할을 함
	close(ch)
	wg.Wait()
}
