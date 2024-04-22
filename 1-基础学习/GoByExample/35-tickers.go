package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond) // 创建一个新的定时触发器，定时触发器的时间间隔是 500 毫秒
	done := make(chan bool)                          // 创建一个 bool 类型的 channel

	go func() { // 启动一个新的 goroutine
		for { // 无限循环
			select { // 开始一个 select 语句，用于同时等待多个 channel
			case <-done: // 如果 done channel 发送了一个值
				return // 结束 goroutine
			case t := <-ticker.C: // 如果定时触发器发送了一个值
				fmt.Println("Tick at", t) // 打印 "Tick at" 和接收到的值
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond) // 主 goroutine 等待 1600 毫秒

	ticker.Stop()                 // 停止定时触发器
	done <- true                  // 向 done channel 发送一个 true，以便新的 goroutine 可以退出
	fmt.Println("Ticker stopped") // 打印 "Ticker stopped"
}
