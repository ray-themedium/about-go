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
	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()

	w.Wait()
}
