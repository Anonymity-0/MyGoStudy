package main

import "fmt"

func main() {
	// 使用 make 函数创建一个新的 channel，这个 channel 用于传递 string 类型的数据
	messages := make(chan string)

	// 使用 go 关键字启动一个新的 goroutine，然后在这个 goroutine 中执行一个匿名函数
	go func() {
		// 在这个匿名函数中，我们将 "ping" 发送到 messages channel
		// 这个操作会阻塞，直到有其他 goroutine 从这个 channel 中接收数据
		messages <- "ping"
	}()

	// 在 main goroutine 中，我们从 messages channel 接收数据，并将其赋值给 msg
	// 这个操作也会阻塞，直到有其他 goroutine 向这个 channel 发送数据
	msg := <-messages

	// 打印接收到的消息
	fmt.Println("msg: ", msg)
}
