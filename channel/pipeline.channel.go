package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

// tireCh 쓰기 전용 체널
func MakeBody(wg *sync.WaitGroup, tireCh chan<- *Car) {
	defer wg.Done()

	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			car := &Car{Body: "Sports car"}
			tireCh <- car
		case <-after:
			close(tireCh)
			return
		}
	}
}

// tireCh 읽기 전용 채널, paintCh 쓰기 전용 체널
func InstallTire(wg *sync.WaitGroup, tireCh <-chan *Car, paintCh chan<- *Car) {
	defer wg.Done()

	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car
	}
	close(paintCh)
}

// paintCh 읽기 전용 체널
func paintCar(wg *sync.WaitGroup, paintCh <-chan *Car, startTime *time.Time) {
	defer wg.Done()

	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(*startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
}

// 생산자 소비자 패턴 (파이프라인 패턴)
// 한쪽에서 데이터를 넣어주고 다른 쪽에서 데이터를 빼서 사용하는 방식
// 채널을 이용해 역할을 나눠 데드락을 피하면서 동시성 프로그래밍하기
func main() {
	var wg sync.WaitGroup
	startTime := time.Now()

	tireCh := make(chan *Car)  // 타이어 생산 채널
	paintCh := make(chan *Car) // 도색 채널

	wg.Add(3)
	go MakeBody(&wg, tireCh)
	go InstallTire(&wg, tireCh, paintCh)
	go paintCar(&wg, paintCh, &startTime)

	wg.Wait()
	fmt.Println("Close the factory")
}
