package main

import "fmt"

// mayPanic 函数使用 panic 函数抛出一个异常
func mayPanic() {
	panic("panic")
}

func main() {
	// 使用 defer 关键字来注册一个匿名函数，这个函数会在 main 函数结束时执行
	defer func() {
		// 使用 recover 函数来捕获 panic 异常
		if r := recover(); r != nil {
			// 如果捕获到了 panic 异常，打印异常信息
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	// 调用 mayPanic 函数，这个函数会抛出一个 panic 异常
	mayPanic()
	// 这里的代码不会被执行，因为 mayPanic 函数中使用 panic 函数抛出了一个异常，
	// main 程序的执行在 panic 点停止，并在继续处理完 defer 后结束。
	fmt.Println("After panic")
}
