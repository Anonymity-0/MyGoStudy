package main

import (
	"fmt" // 用于格式化输出
	"os"  // 用于操作系统相关的操作，如结束程序
)

func main() {
	// 使用 defer 语句确保在函数返回时输出 "!"
	defer fmt.Println("!")

	// 使用 os.Exit 函数结束程序并返回状态码 3
	// 注意，defer 语句不会被执行，！不会被输出
	os.Exit(3)
}
