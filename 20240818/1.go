package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type GoroutinePool struct {
	workerCount int
	tasks       chan Task
	wg          sync.WaitGroup
}

func NewGoroutinePool(workerCount int) *GoroutinePool {
	return &GoroutinePool{
		workerCount: workerCount,
		tasks:       make(chan Task),
	}
}

func (p *GoroutinePool) Start() {
	for i := 0; i < p.workerCount; i++ {
		go p.worker()
	}
}

func (p *GoroutinePool) AddTask(task Task) {
	p.wg.Add(1)
	p.tasks <- task
}

func (p *GoroutinePool) Stop() {
	close(p.tasks)
	p.wg.Wait()
}

func (p *GoroutinePool) worker() {

	for task := range p.tasks {
		task()
		p.wg.Done()
	}
}

func main() {
	pool := NewGoroutinePool(5)
	pool.Start()
	for i := 1; i <= 10; i++ {
		taskID := i
		pool.AddTask(func() {
			fmt.Printf("任务 %d 开始执行\n", taskID)
			time.Sleep(2 * time.Second) // 模拟任务执行时间
			fmt.Printf("任务 %d 执行完成\n", taskID)
		})
	}

	pool.Stop()
	fmt.Println("所有任务执行完毕")
}
