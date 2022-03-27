package main

import (
	"context"
	"fmt"
	"sync"
)

var w sync.WaitGroup

func goroutineFunc(ctx context.Context) {
	if value := ctx.Value("key"); value != "" {
		v := value.(string)
		fmt.Println(v)
	}
	w.Done()
}

func main() {
	w.Add(1)

	ctx := context.WithValue(context.Background(), "key", "value")
	go goroutineFunc(ctx)

	w.Wait()
}
