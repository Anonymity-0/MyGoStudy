package main

import "fmt"

func main() {

	// 默认情况下，channel 是无缓冲的，这意味着只有在对应的接收（<- chan）已经准备好的情况下，才允许进行发送（chan <-）

	// 有缓冲的 channel 允许在没有对应接收方的情况下，缓存限定数量的值
	// 这里我们创建了一个带有缓冲区大小为 2 的 channel
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println("msg1:", <-messages)
	fmt.Println("msg2:", <-messages)
}
