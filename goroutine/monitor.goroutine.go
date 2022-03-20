/*
공유 메모리는 스레드끼리 통신하는 기법으로 go 언어는 고루틴이 데이터의 일정 영역
을 소유할 수 있도록 동기화 기능을 제공한다.
어떤 고루틴이 가진 공유 데이터에 다른 고루틴이 접근하려면 반드시 데이터를 보유한
 고루틴에게 먼저 메시지를 보내야 한다.
 이런식으로 데이터가 손상되지 않게 보호하는데 이렇게 공유 데이터를 가진 고루틴을
  모니터 고루틴이라 부른다.
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var readValue = make(chan int)  // 새로운 난수를 읽는데 사용
var writeValue = make(chan int) // 새로운 난수를 가져오는 데 사용

func set(newValue int) {
	writeValue <- newValue
}

func read() int {
	return <-readValue
}

// 공유변수 value는 monitor 함수를 제외한 어느 누구도 건드리지 않음
func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give an inteager!")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Going to create %d random nubmers.\n", n)
	rand.Seed(time.Now().Unix())

	go monitor()

	var w sync.WaitGroup

	for r := 0; r < n; r++ {
		w.Add(1)
		go func() {
			defer w.Done()
			set(rand.Intn(10 * n))
		}()
	}
	w.Wait()
	fmt.Printf("\nLast value: %d\n", read())
}
