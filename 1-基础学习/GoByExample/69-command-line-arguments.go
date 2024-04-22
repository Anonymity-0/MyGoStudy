package main

import (
	"fmt" // 用于格式化输出
	"os"  // 用于操作系统相关的操作，如获取命令行参数
)

func main() {
	// os.Args 是一个包含程序启动命令行参数的字符串切片
	// os.Args[0] 是程序的路径
	argsWithProg := os.Args

	// os.Args[1:] 是除了程序路径以外的命令行参数
	argsWithoutProg := os.Args[1:]

	// os.Args[3] 是第三个命令行参数
	arg := os.Args[3]

	// 输出所有命令行参数，包括程序的路径
	fmt.Println(argsWithProg)
	// 输出除了程序路径以外的命令行参数
	fmt.Println(argsWithoutProg)
	// 输出第三个命令行参数
	fmt.Println(arg)
}
