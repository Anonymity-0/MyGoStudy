package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5) // 创建一个可以存储 5 个 int 类型值的 channel
	done := make(chan bool)   // 创建一个 bool 类型的 channel

	go func() { // 启动一个新的 goroutine
		for { // 无限循环
			//  在 Go 语言中，从 channel 接收数据时，可以返回两个值。第一个值是从 channel 接收到的数据，第二个值是一个 bool 类型的值，表示 channel 是否已经关闭并且没有更多的数据可以接收。
			j, more := <-jobs // 从 jobs channel 接收数据
			if more {         // 如果 jobs channel 还有数据
				fmt.Println("received job", j) // 打印接收到的数据
			} else { // 如果 jobs channel 没有数据
				fmt.Println("received all jobs") // 打印 "received all jobs"
				done <- true                     // 向 done channel 发送 true
				return                           // 结束 goroutine
			}
		}
	}()

	for j := 1; j <= 3; j++ { // 循环 3 次
		jobs <- j                          // 向 jobs channel 发送数据
		fmt.Println("sent job", j)         // 打印发送的数据
		time.Sleep(100 * time.Millisecond) // 等待 100 毫秒
	}
	close(jobs)                  // 关闭 jobs channel
	fmt.Println("sent all jobs") // 打印 "sent all jobs"

	<-done // 从 done channel 接收数据，如果 done channel 没有数据，这会阻塞，直到有数据可接收
}
