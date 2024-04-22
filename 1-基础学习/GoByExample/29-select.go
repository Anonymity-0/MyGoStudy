package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string) // 创建一个 string 类型的 channel
	c2 := make(chan string) // 创建另一个 string 类型的 channel

	go func() { // 启动一个新的 goroutine
		time.Sleep(1 * time.Second) // 等待 1 秒
		c1 <- "one"                 // 向 c1 channel 发送一个字符串 "one"
	}()

	go func() { // 启动另一个新的 goroutine
		time.Sleep(2 * time.Second) // 等待 2 秒
		c2 <- "two"                 // 向 c2 channel 发送一个字符串 "two"
	}()

	for i := 0; i < 2; i++ { // 循环两次，因为我们有两个 goroutine，每个都会向其各自的 channel 发送数据
		select { // 使用 select 语句等待多个 channel 的操作
		case msg1 := <-c1: // 如果 c1 channel 准备好了，接收数据并赋值给 msg1
			fmt.Println("received", msg1) // 打印接收到的数据
		case msg2 := <-c2: // 如果 c2 channel 准备好了，接收数据并赋值给 msg2
			fmt.Println("received", msg2) // 打印接收到的数据
		}
	}
}
