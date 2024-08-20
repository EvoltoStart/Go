package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	num  = 100
	lock = sync.Mutex{}
)

func main() {
	var ww = sync.WaitGroup{}
	ww.Add(2)
	go func(name string) {
		defer ww.Done()
		for {
			lock.Lock()
			if num > 0 {
				time.Sleep(1 * time.Second)
				num--
				fmt.Printf("窗口%s：卖出一张票，剩余%d张票\n", name, num)
			} else {
				lock.Unlock()
				break
			}
			lock.Unlock()

		}
	}("A")
	go func(name string) {
		defer ww.Done()
		for {
			lock.Lock()
			if num > 0 {
				time.Sleep(1 * time.Second)
				num--
				fmt.Printf("窗口%s：卖出一张票，剩余%d张票\n", name, num)
			} else {
				lock.Unlock()
				break
			}
			lock.Unlock()

		}
	}("B")
	ww.Wait()

}
