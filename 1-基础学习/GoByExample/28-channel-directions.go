package main

import "fmt"

// 定义一个名为 ping 的函数，它接受一个只写channel（chan<- string）和一个 string 类型的参数
// 这个函数将 msg 发送到 pings channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 定义一个名为 pong 的函数，它接受一个只能读数据的 channel（<-chan string）和一个只能写数据的 channel（chan<- string）
// 这个函数从 pings channel 接收数据，然后将接收到的数据发送到 pongs channel
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1) // 创建一个 string 类型的 channel，容量为 1
	pongs := make(chan string, 1) // 创建一个 string 类型的 channel，容量为 1
	ping(pings, "passed message") // 调用 ping 函数，将 "passed message" 发送到 pings channel
	pong(pings, pongs)            // 调用 pong 函数，从 pings channel 接收数据，然后将接收到的数据发送到 pongs channel
	fmt.Println("msg: ", <-pongs) // 从 pongs channel 接收并打印出数据
}
