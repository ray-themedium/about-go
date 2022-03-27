package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			w.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}

func main() {
	w.Add(1)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // 강제로 종료할 수 있는 함수를 제공
	go PrintEverySecond(ctx)

	w.Wait()
}
