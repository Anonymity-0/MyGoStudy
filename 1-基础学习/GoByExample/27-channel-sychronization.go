package main

import (
	"fmt"
	"time"
)

// 定义一个名为 worker 的函数，它接受一个 bool 类型的 channel 作为参数
func worker(done chan bool) {
	fmt.Println("working...") // 打印 "working..."
	time.Sleep(time.Second)   // 暂停一秒，模拟一个耗时的任务
	fmt.Println("done")       // 打印 "done"
	done <- true              // 将 true 发送到 done channel，表示 worker 函数已经完成
}

func main() {
	done := make(chan bool, 1) // 创建一个 bool 类型的 channel，容量为 1
	go worker(done)            // 使用 go 关键字启动一个新的 goroutine，然后在这个 goroutine 中调用 worker 函数，传入 done channel 作为参数
	<-done                     // 在 main goroutine 中，我们从 done channel 接收数据。这个操作会阻塞，直到 worker goroutine 向 done channel 发送数据
}
