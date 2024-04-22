package main

import (
	"flag" // 用于解析命令行参数
	"fmt"  // 用于格式化输出
)

func main() {
	// 使用 flag.String 函数定义一个字符串命令行参数
	// 第一个参数是命令行参数的名字，第二个参数是默认值，第三个参数是说明
	wordPtr := flag.String("word", "foo", "a string")

	// 使用 flag.Int 函数定义一个整数命令行参数
	numbPtr := flag.Int("numb", 42, "an int")
	// 使用 flag.Bool 函数定义一个布尔值命令行参数
	forkPtr := flag.Bool("fork", false, "a bool")

	// 定义一个字符串变量
	var svar string
	// 使用 flag.StringVar 函数定义一个字符串命令行参数，第一个参数是字符串变量的地址
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 使用 flag.Parse 函数解析命令行参数
	flag.Parse()

	// 输出所有命令行参数
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	// 使用 flag.Args 函数获取非标志命令行参数
	fmt.Println("tail:", flag.Args())
}
