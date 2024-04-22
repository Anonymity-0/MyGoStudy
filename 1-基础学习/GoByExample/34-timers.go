package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second) // 创建一个新的定时器，定时器的时间间隔是 2 秒

	<-timer.C // 阻塞等待定时器的 channel，直到它发送一个值，表示定时器已经到期

	fmt.Println("Timer 1 expired") // 打印 "Timer 1 expired"

	timer2 := time.NewTimer(time.Second) // 创建一个新的定时器，定时器的时间间隔是 1 秒

	go func() { // 启动一个新的 goroutine
		<-timer2.C                     // 阻塞等待定时器的 channel，直到它发送一个值，表示定时器已经到期
		fmt.Println("Timer 2 expired") // 打印 "Timer 2 expired"
	}()

	//这是因为在新的 goroutine 还来不及接收 timer2.C 的值（即定时器还未到期）时，主 goroutine 就已经调用了 timer2.Stop() 方法来停止定时器。因此，timer2.C 不会发送值，新的 goroutine 在 <-timer2.C 这行代码上会一直阻塞，所以 "Timer 2 expired" 不会被打印。
	stop2 := timer2.Stop() // 停止定时器，如果定时器还没有到期，那么返回 true；否则返回 false

	if stop2 { // 如果定时器被成功停止
		fmt.Println("Timer 2 stopped") // 打印 "Timer 2 stopped"
	}

	time.Sleep(2 * time.Second) // 让主 goroutine 等待 2 秒，以便新的 goroutine 有足够的时间来打印 "Timer 2 expired"
}
