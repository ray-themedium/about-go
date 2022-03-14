// 주어진 범위에서 난수를 생성하다가 같은 값이 두 번 나오면 멈추는 프로그램
// 데이터는 파이프라인으로 구성한 채널을 통해 흐름

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)
var signal chan struct{}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, in chan<- int) {
	for {
		select {
		case <-signal:
			close(in)
			return
		case in <- random(min, max):
		}
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		_, ok := DATA[x]
		if ok {
			signal <- struct{}{}
		} else {
			fmt.Print(x, " ")
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum += x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters")
		os.Exit(1)
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d\n", n1, n2)
		return
	}

	rand.Seed(time.Now().UnixNano())
	IN := make(chan int)
	OUT := make(chan int)
	signal = make(chan struct{})

	go first(n1, n2, IN)
	go second(OUT, IN)
	third(OUT)
}
