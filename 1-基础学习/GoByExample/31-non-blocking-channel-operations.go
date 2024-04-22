package main

import "fmt"

func main() {
	messages := make(chan string) // 创建一个 string 类型的 channel
	signals := make(chan bool)    // 创建一个 bool 类型的 channel

	//这是一个非阻塞接收的例子。 如果在 messages 中存在，然后 select 将这个值带入 <-messages case 中。 否则，就直接到 default 分支中。
	select {
	case msg := <-messages: // 尝试从 messages channel 接收数据
		fmt.Println("received message", msg) // 如果接收到数据，打印接收到的数据
	default:
		fmt.Println("no message received") // 如果没有接收到数据，执行 default 分支，打印 "no message received"
	}

	//一个非阻塞发送的例子，代码结构和上面接收的类似。 msg 不能被发送到 message 通道，因为这是 个无缓冲区通道，并且也没有接收者，因此， default 会执行
	msg := "hi" // 定义一个字符串 "hi"
	select {
	case messages <- msg: // 尝试向 messages channel 发送数据
		fmt.Println("sent message", msg) // 如果发送成功，打印发送的数据
	default:
		fmt.Println("no message sent") // 如果没有发送成功，执行 default 分支，打印 "no message sent"
	}

	select {
	case msg := <-messages: // 尝试从 messages channel 接收数据
		fmt.Println("received message", msg) // 如果接收到数据，打印接收到的数据
	case sig := <-signals: // 尝试从 signals channel 接收数据
		fmt.Println("received signal", sig) // 如果接收到数据，打印接收到的数据
	default:
		fmt.Println("no activity") // 如果两个 channel 都没有数据，执行 default 分支，打印 "no activity"
	}
}