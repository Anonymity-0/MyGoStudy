package main

import (
	"fmt"
	"sync"
	"time"
)

// worker 函数定义了一个工作函数，它模拟了一些耗时的工作
func worker(id int) {
	fmt.Printf("Worker %d started\n", id) // 打印工作开始信息

	time.Sleep(time.Second) // 模拟耗时的工作

	fmt.Printf("Worker %d done\n", id) // 打印工作完成信息
}

func main() {
	var wg sync.WaitGroup // 创建一个新的 WaitGroup

	for i := 1; i <= 5; i++ { // 启动 5 个 worker goroutine
		wg.Add(1)   // 增加 WaitGroup 的计数器
		i := i      // 创建一个新的变量 i，以便在 goroutine 中使用
		go func() { // 启动一个新的 goroutine
			defer wg.Done() // 在 goroutine 结束时，调用 wg.Done() 来减少 WaitGroup 的计数器
			worker(i)       // 调用 worker 函数
		}()
	}
	wg.Wait() // 阻塞，直到所有的 goroutine 都调用了 wg.Done()
}
