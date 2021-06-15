package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body string
	Tire string
	Color string
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick: // Make a body
			car := &Car{}
			car.Body = "Sports Car" // 바디 생성
			tireCh <- car // 타이어 만드는 곳으로 보내기 위한 컨베이어 밸트 (채널)
		case <-after: // 10초 뒤 종료
			close(tireCh) // 채널 닫기
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh { // tire 채널에서 빼올 때 까지 대기
		time.Sleep(time.Second) // 타이어 설치가 1초 걸림
		car.Tire = "Winter tire"
		paintCh <- car // 패인트를 칠할 수 있도록 채널에 전달
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second) // 패인트칠이 1초 걸림
		car.Color = "red"
		duration := time.Now().Sub(startTime) // 한 자동차 당 경과 시간 계산
		fmt.Printf("%.2f Complete Car : %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}