package main

import (
	"fmt"
)

func main() {

	wg.Add(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		defer wg.Done()
		ch1 <- 1
		close(ch1)
	}()
	go func() {
		defer wg.Done()
		ch2 <- 2
		close(ch2)
	}()

	for {
		select {
		case i, ok := <-ch1:
			if !ok {
				//ch1 = nil
				break
			} else {
				fmt.Println(i)
			}
		case j, ok := <-ch2:
			if !ok {
				//ch2 = nil
				break
			} else {
				fmt.Println(j)
			}
		}
		if ch1 == nil || ch2 == nil {
			break
		}
	}
	wg.Wait()

}
