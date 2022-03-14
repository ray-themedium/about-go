package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)            // 1초 간격 시그널 채널
	terminate := time.After(10 * time.Second) // 10초 이후 시그널

	for {
		select {
		case <-tick:
			fmt.Println(time.Now())
		case <-terminate:
			fmt.Println("Terminated")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
			// case <-quit:
			// 	wg.Done()
			// 	return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	// quit := make(chan bool) // 종료 체널

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	// quit <- true
	// close(quit)
	wg.Wait()

}
