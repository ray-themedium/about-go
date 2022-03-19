/*
들어온 요청은 모두 채널에 전달되고 각각을 하나씩 처리한다.
채널이 어떤 요청에 대한 처리 작업을 끝내면 호출한 측으로 새로운 작업을 처리할
준비가 됐다는 메세지를 보낸다.
따라서 채널에서 사용하는 버퍼의 크기에 따라 동시에 처리할 수 있는 요청의 수가
결정된다.
*/
package main

import "fmt"

func main() {
	numbers := make(chan int, 5)
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	for i := 0; i < counter+5; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to be done!")
			break
		}
	}
}
