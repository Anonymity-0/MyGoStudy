package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// sync/atomic 包中的 AddUint64 函数来安全地增加计数器。这个函数会保证增加操作的原子性，即任何时候都只有一个 goroutine 可以增加计数器，其他的 goroutine 都必须等待，直到这个操作完成。
func main() {
	var ops uint64 // 定义一个无符号64位整数作为原子计数器

	var wg sync.WaitGroup // 创建一个新的 WaitGroup

	for i := 0; i < 50; i++ { // 启动 50 个 goroutine
		wg.Add(1)   // 增加 WaitGroup 的计数器
		go func() { // 启动一个新的 goroutine
			for c := 0; c < 1000; c++ { // 每个 goroutine 都会增加计数器 1000 次
				atomic.AddUint64(&ops, 1) // 使用 atomic.AddUint64 函数来安全地增加计数器
			}
			wg.Done() // 在 goroutine 结束时，调用 wg.Done() 来减少 WaitGroup 的计数器
		}()
	}
	wg.Wait() // 阻塞，直到所有的 goroutine 都调用了 wg.Done()

	fmt.Println("ops:", ops) // 打印最终的计数器值
}
