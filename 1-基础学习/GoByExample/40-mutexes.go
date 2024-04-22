package main

import (
	"fmt"
	"sync"
)

// Container 结构体包含一个互斥锁和一个 map 类型的计数器
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// inc 方法用于增加指定名称的计数器的值
func (c *Container) inc(name string) {
	c.mu.Lock()         // 加锁，保护共享资源
	defer c.mu.Unlock() // 在方法结束时解锁，释放共享资源
	c.counters[name]++  // 修改共享资源
}

func main() {
	// 创建一个 Container 实例，其中包含两个计数器 "a" 和 "b"
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup // 创建一个新的 WaitGroup

	// doIncrement 函数用于增加指定名称的计数器的值 n 次
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name) // 调用 inc 方法来增加计数器的值
		}
		wg.Done() // 在函数结束时，调用 wg.Done() 来减少 WaitGroup 的计数器
	}
	wg.Add(3) // 增加 WaitGroup 的计数器
	// 启动 3 个 goroutine 来并发地增加计数器的值
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()               // 阻塞，直到所有的 goroutine 都调用了 wg.Done()
	fmt.Println(c.counters) // 打印最终的计数器值
}
