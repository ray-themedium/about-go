package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select { // ch, quit 채널을 모두 기다림
		case n := <- ch:
			fmt.Printf("Square: %d\n", n * n)
			time.Sleep(time.Second)
		case <- quit:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup // 고루틴 그룹 선언
	ch := make(chan int) // int 타입의 채널 선언 (버퍼 0)
	quit := make(chan bool) // bool 타입의 채널 선언 (버퍼 0)

	wg.Add(1) // 메인 루틴을 제외하고 추가한 고루틴 한 개
	go square(&wg, ch, quit) // 다른 루틴에서 동작시킬 함수

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}