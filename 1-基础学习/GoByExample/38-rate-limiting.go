package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5) // 创建一个可以存储 5 个 int 类型值的 channel
	for i := 1; i <= 5; i++ {
		requests <- i // 将值发送到 channel
	}
	close(requests) // 关闭 channel

	limiter := time.Tick(time.Millisecond * 200) // 创建一个定时触发器，每 200 毫秒触发一次

	for req := range requests { // 从 channel 中接收值，直到 channel 被关闭
		<-limiter                               // 从 limiter channel 中接收值，如果 limiter channel 没有值，这会阻塞，直到有值可接收
		fmt.Println("request", req, time.Now()) // 打印请求信息和当前时间
	}

	burstyLimiter := make(chan time.Time, 3) // 创建一个可以存储 3 个 time.Time 类型值的 channel

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now() // 将当前时间发送到 channel
	}
	go func() { // 启动一个新的 goroutine
		for t := range time.Tick(time.Millisecond * 200) { // 创建一个定时触发器，每 200 毫秒触发一次
			burstyLimiter <- t // 将当前时间发送到 channel
		}
	}()

	burstyLimiterRequests := make(chan int, 5) // 创建一个可以存储 5 个 int 类型值的 channel
	for i := 1; i <= 5; i++ {
		burstyLimiterRequests <- i // 将值发送到 channel
	}
	close(burstyLimiterRequests)             // 关闭 channel
	for req := range burstyLimiterRequests { // 从 channel 中接收值，直到 channel 被关闭
		<-burstyLimiter                         // 从 burstyLimiter channel 中接收值，如果 burstyLimiter channel 没有值，这会阻塞，直到有值可接收
		fmt.Println("request", req, time.Now()) // 打印请求信息和当前时间
	}
}
