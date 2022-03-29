package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var mutex sync.Mutex

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Give me a natural number!")
		os.Exit(1)
	}
	numGR, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var w sync.WaitGroup
	var i int
	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numGR; i++ {
		w.Add(1)
		go func(j int) {
			defer w.Done()
			mutex.Lock() // lock이 없으면 경쟁 상태 발생
			k[j] = j // race condition
			mutex.Unlock()
		}(i)
	}

	w.Wait()
	k[2] = 10
	fmt.Printf("k = %#v\n", k)
}
