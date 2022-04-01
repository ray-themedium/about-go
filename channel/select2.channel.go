package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			fmt.Println("Close end channel")
			return
		case <-time.After(4 * time.Second): // default 역할을 함
			fmt.Println("\ntime.After()!")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an inteager")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers,\n", n)
	go gen(0, 2*n, createNumber, end)

	for i := 0; i < 30; i++ {
		fmt.Printf("%d \n", <-createNumber)
	}
	time.Sleep(3 * time.Second)
	// fmt.Println("Exiting...")
	end <- true // 이 코드가 실행되면 데드락이 발생 -> all goroutines are asleep
	time.Sleep(1 * time.Second)
}
