package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// readOp 和 writeOp 结构体分别表示读操作和写操作
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps, writeOps uint64 // 记录读操作和写操作的次数

	reads := make(chan readOp)   // 创建一个 readOp 类型的 channel
	writes := make(chan writeOp) // 创建一个 writeOp 类型的 channel

	// 启动一个 goroutine 来处理所有的读操作和写操作
	go func() {
		var state = make(map[int]int) // 创建一个 map 来保存状态
		for {
			select {
			case read := <-reads: // 处理读操作
				read.resp <- state[read.key]
			case write := <-writes: // 处理写操作
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动 100 个 goroutine 来发送读操作
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),   // 随机选择一个 key
					resp: make(chan int)} // 创建一个新的 channel 来接收响应
				reads <- read                 // 发送读操作
				<-read.resp                   // 等待响应
				atomic.AddUint64(&readOps, 1) // 增加读操作的次数
				time.Sleep(time.Millisecond)  // 等待一段时间
			}
		}()
	}

	// 启动 10 个 goroutine 来发送写操作
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),    // 随机选择一个 key
					val:  rand.Intn(100),  // 随机选择一个 value
					resp: make(chan bool)} // 创建一个新的 channel 来接收响应
				writes <- write                // 发送写操作
				<-write.resp                   // 等待响应
				atomic.AddUint64(&writeOps, 1) // 增加写操作的次数
				time.Sleep(time.Millisecond)   // 等待一段时间
			}
		}()
	}

	time.Sleep(time.Second) // 等待一段时间

	// 获取并打印最终的读操作和写操作的次数
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
