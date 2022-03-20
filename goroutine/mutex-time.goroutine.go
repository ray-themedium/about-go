/*
sync.RWMutex 사용이 훨씬 효율 적이다.
*/

package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var Password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	defer c.RWM.Unlock()

	fmt.Println("LChange")
	time.Sleep(5 * time.Second)
	c.password = pass
}

func show(c *secret) string {
	c.RWM.RLock()
	defer c.RWM.RUnlock()

	fmt.Println("show")
	time.Sleep(3 * time.Second)
	return c.password
}

func showWithLock(c *secret) string {
	c.M.Lock()
	defer c.M.Unlock()

	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	return c.password
}

func main() {
	var showFunction = func(c *secret) string { return "" }

	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex")
		showFunction = show
	} else {
		fmt.Println("Using symc.Mutex")
		showFunction = showWithLock
	}

	var w sync.WaitGroup
	fmt.Println("Pass:", showFunction(&Password))

	for i := 0; i < 15; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			fmt.Println("Go Pass:", showFunction(&Password))
		}()
	}

	go func() {
		w.Add(1)
		defer w.Done()
		Change(&Password, "123456")
	}()

	w.Wait()
	fmt.Println("Pass:", showFunction(&Password))
}
