package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(acount *Account) {
	mutex.Lock()
	defer mutex.Unlock()
	if acount.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", acount.Balance))
	}
	acount.Balance += 1000
	time.Sleep(10 * time.Nanosecond)
	acount.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	acount := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(acount)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
