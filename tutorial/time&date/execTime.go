package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second)
	end := time.Since(start)
	fmt.Println("1초가 걸린 시간 (1):", end)

	start = time.Now()
	time.Sleep(time.Second)
	end = time.Now().Sub(start)
	fmt.Println("1초가 걸린 시간 (2):", end)

	start = time.Now()
	time.Sleep(2 * time.Second)
	end = time.Since(start)
	fmt.Println("2초가 걸린 시간", end)

	sum := 0
	start = time.Now()
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	end = time.Since(start)
	fmt.Println("1 ~ 100000000 합", end)

	sum = 0
	start = time.Now()
	for i := 0; i < 1000000000; i++ {
		sum += i
	}
	end = time.Since(start)
	fmt.Println("1 ~ 1000000000 합", end)
}
