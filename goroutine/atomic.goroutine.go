/*
atomic 변수에 대한 읽기 연산과 쓰기 연산은 모두 atomic 패키지에서 제공하는
atomic 함수로 처리해야 한다.
*/
package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
)

type atomCounter struct {
	val int64
}

func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	minusX := flag.Int("x", 100, "Goroutines")
	minusY := flag.Int("y", 200, "Value")
	flag.Parse()
	x := *minusX
	y := *minusY

	var w sync.WaitGroup
	counter := atomCounter{}

	for i := 0; i < x; i++ {
		w.Add(1)
		go func(no int) {
			defer w.Done()
			for i := 0; i < y; i++ {
				atomic.AddInt64(&counter.val, 1)
			}
		}(i)
	}

	w.Wait()
	fmt.Println(counter.Value())
}
