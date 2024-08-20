package main

import (
	"fmt"
	"sync"
)

var (
	w = sync.WaitGroup{}
)

func main() {
	w.Add(2)
	c := make(chan int)
	go func(name string) {
		defer w.Done()
		for i := 0; i <= 50; i++ {
			fmt.Printf("%s:%d\n", name, i)
		}
		c <- 1

	}("A")
	go func(name string) {
		defer w.Done()
		<-c
		for i := 51; i <= 100; i++ {
			fmt.Printf("%s:%d\n", name, i)
		}

	}("B")
	w.Wait()
}
