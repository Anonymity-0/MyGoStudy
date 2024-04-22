package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1) // 创建一个 string 类型的 channel

	go func() { // 启动一个新的 goroutine
		time.Sleep(2 * time.Second) // 等待 2 秒
		c1 <- "result 1"            // 向 c1 channel 发送一个字符串 "result 1"
	}()

	select { // 使用 select 语句等待多个 channel 的操作
	case res := <-c1: // 如果 c1 channel 准备好了，接收数据并赋值给 res
		fmt.Println("result 1:", res) // 打印接收到的数据
	case <-time.After(1 * time.Second): // 如果 1 秒后没有接收到数据，执行这个 case 分支
		fmt.Println("timeout 1") // 打印 "timeout 1"
	}

	c2 := make(chan string, 1) // 创建另一个 string 类型的 channel

	go func() { // 启动另一个新的 goroutine
		time.Sleep(2 * time.Second) // 等待 2 秒
		c2 <- "result 2"            // 向 c2 channel 发送一个字符串 "result 2"
	}()

	select { // 使用 select 语句等待多个 channel 的操作
	case res := <-c2: // 如果 c2 channel 准备好了，接收数据并赋值给 res
		fmt.Println("result 2:", res) // 打印接收到的数据
	case <-time.After(3 * time.Second): // 如果 3 秒后没有接收到数据，执行这个 case 分支
		fmt.Println("timeout 2") // 打印 "timeout 2"
	}
}
