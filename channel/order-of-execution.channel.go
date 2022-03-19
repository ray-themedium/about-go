package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	<-a // 채널에 시그널을 기다리는 중 (close 또는 값)
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

func C(a chan struct{}) {
	<-a
	fmt.Println("C()!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go C(z)    // z의 시그널을 기다림
	go A(x, y) // x의 시그널을 기다림 2
	go C(z)    // z의 시그널을 기다림
	go B(y, z) // y의 시그널을 기다림 3
	go C(z)    // z의 시그널을 기다림

	close(x) // 시그널 시작 1
	time.Sleep(3 * time.Second)
}
