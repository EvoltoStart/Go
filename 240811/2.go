package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	i  = 65
	wg = sync.WaitGroup{}
)

func main() {

	//wg.Add(2)
	//go func(name string) {
	//	defer wg.Done()
	//	for i <= 90 {
	//
	//		lock.Lock()
	//		fmt.Printf("%s:%c\n", name, i)
	//		i++
	//		lock.Unlock()
	//		time.Sleep(time.Second)
	//	}
	//}("A")
	//go func(name string) {
	//	defer wg.Done()
	//	for i <= 90 {
	//		lock.Lock()
	//		fmt.Printf("%s:%c\n", name, i)
	//		i++
	//		lock.Unlock()
	//		time.Sleep(time.Second)
	//	}
	//}("B")
	//
	//wg.Wait()
	wg.Add(2)
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go func(name string) {
		defer wg.Done()
		for i := 65; i <= 90; i += 2 {
			if <-ch1 {
				fmt.Printf("%s:%c\n", name, i)
				ch2 <- true
				time.Sleep(time.Second)

			}
		}
		close(ch1)
	}("A")
	go func(name string) {
		defer wg.Done()
		for i := 66; i <= 90; i += 2 {
			if <-ch2 {
				fmt.Printf("%s:%c\n", name, i)
				if i < 90 {
					ch1 <- true

				}
				time.Sleep(time.Second)
			}
		}
		close(ch2)

	}("B")
	ch1 <- true
	wg.Wait()

}
