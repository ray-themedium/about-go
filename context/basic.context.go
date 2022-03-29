package main

import (
	"context"
	"fmt"
	"time"
)

// Done 채널이 닫히는 시점은 앞에 나오는 코드처럼 cancel() 함수가 호출될 때나
// 부모 컨텍스트의 Done 채널이 닫힐 때다.
func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1) // Done 채널도 생성

	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1():", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}
	return
}

func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second) // timeout되면 자동으로 cancel 실행
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second) // deadline 시간이 되면 자동으로 cancel 실행
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3():", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
	return
}

func main() {
	delay := 2
	f1(delay)
	f2(delay)
	f3(delay)
}
