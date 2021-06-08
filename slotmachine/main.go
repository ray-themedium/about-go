package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Balance = 1000 // 초기 잔액
	EarnPoint = 500 // 맞췄을 때 포인트
	LosePoint = 100 // 틀렸을 때 포인트
	VictoryPoint = 5000 // 게임 승리 포인트
	GameOverPoint = 0 // 게임 오버 포인트
)

var stdin = bufio.NewReader(os.Stdin)

func InputInt() (int, error) {
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		stdin.ReadString('\n')
	}
	return input, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	balance := Balance

	for {
		fmt.Print("1~5 사이 값을 입력하세요:")
		n, err := InputInt()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else if n < 1 || n > 5 {
			fmt.Println("1~5 사이의 값만 입력하세요.")
		} else {
			r := rand.Intn(5) + 1 // 1~5 랜덤 발생
			if n == r {
				balance += EarnPoint
				fmt.Println("맞춤. 남은 돈:", balance)
				if balance >= VictoryPoint {
					fmt.Println("게임 승리")
					break
				}
			} else {
				balance -= LosePoint
				fmt.Println("꽝. 남은 돈:", balance)
				if balance <= GameOverPoint {
					fmt.Println("게임 오버")
					break
				}
			}
		}
	}
}