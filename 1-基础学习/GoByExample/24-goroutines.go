package main

import (
	"fmt"
	"time"
)

// 定义一个名为 f 的函数，它接受一个字符串参数 from
// 这个函数会打印出 from 的值和一个从 0 到 2 的整数
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println("from", from, ":", i)
	}
}

func main() {
	// 直接调用 f 函数，传入 "direct" 作为参数
	f("direct")

	// 使用 go 关键字启动一个新的 goroutine，然后在这个 goroutine 中调用 f 函数，传入 "goroutine" 作为参数
	go f("goroutine")

	// 使用 go 关键字启动一个新的 goroutine，然后在这个 goroutine 中调用一个匿名函数
	// 这个匿名函数接受一个字符串参数 msg，并打印出 msg 的值
	// 我们传入 "going" 作为参数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 让 main 函数暂停 1 秒，以确保上面的两个 goroutine 有足够的时间执行
	// 如果没有这行代码，main 函数可能会在上面的 goroutines 还没来得及执行就结束，因为 main 函数结束时，所有的 goroutines 都会被立即停止
	time.Sleep(time.Second)
	fmt.Println("done")
}
