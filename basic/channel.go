package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	// var wg sync.WaitGroup
	ch := make(chan int)

	// wg.Add(1)
	// go square(&wg, ch)
	ch <- 9
	// wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <- ch // 데이터가 없다면 대기한다.

	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}