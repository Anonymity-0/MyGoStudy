package main

import "os"

func main() {
	// 当 main 中触发第一个 panic 时，程序就会退出而不会执行代码的其余部分。 如果你想看到程序尝试创建 /tmp/file 文件，请注释掉第一个panic。
	panic("a problem") // 使用 panic 函数抛出一个异常

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err) // 如果创建文件时发生错误，我们使用 panic 函数抛出这个错误
	}
}
