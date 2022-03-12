package main

import (
	"fmt"
	"time"
)

func PrintData() {
	dataList := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	for _, v := range dataList {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNum() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintData()
	go PrintNum()

	time.Sleep(3 * time.Second)
}
