package main

import (
	"fmt"
	"sync"
)

func main() {
	var ww = sync.WaitGroup{}
	ww.Add(3)
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go func(name string) {
		defer ww.Done()
		for i := 0; i < 5; i++ {
			if <-ch1 {
				fmt.Printf("%s:A->B->C;第%d次\n", name, i+1)
				ch2 <- true
			}
		}

	}("A")
	go func(name string) {
		defer ww.Done()
		for i := 0; i < 5; i++ {
			if <-ch2 {
				fmt.Printf("%s:A->B->C;第%d次\n", name, i+1)
				ch3 <- true
			}

		}
	}("B")
	go func(name string) {
		defer ww.Done()
		for i := 0; i < 5; i++ {
			if <-ch3 {
				fmt.Printf("%s:A->B->C;第%d次\n", name, i+1)
				if i < 4 {
					ch1 <- true
				}
			}

		}
	}("C")

	ch1 <- true

	ww.Wait()
}
