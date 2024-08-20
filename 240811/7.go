package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	startWatchDogWithContext()
}
func startWatchDogWithContext() {
	var wg sync.WaitGroup
	wg.Add(3)
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDogWithContext(ctx, "watchDog_1")
	}()
	go func() {
		defer wg.Done()
		//ctx1, cancelFunc := context.WithCancel(ctx)
		watchDogWithContext(ctx, "watchDog_2")
	}()
	go func() {
		defer wg.Done()
		watchDogWithContext(ctx, "watchDog_3")
	}()
	time.Sleep(5 * time.Second)
	stop()
	wg.Wait()
}

func watchDogWithContext(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(time.Second)
	}
}
