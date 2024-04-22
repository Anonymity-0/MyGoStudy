package main

import (
	"fmt"     // 用于格式化输出
	"os"      // 用于操作系统相关的操作，如操作环境变量
	"strings" // 用于操作字符串
)

func main() {
	// 使用 os.Setenv 函数设置环境变量 "FOO" 的值为 "1"
	os.Setenv("FOO", "1")
	// 使用 os.Getenv 函数获取环境变量 "FOO" 的值，并输出
	fmt.Println("FOO:", os.Getenv("FOO"))
	// 使用 os.Getenv 函数获取环境变量 "BAR" 的值，并输出
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	// 使用 os.Environ 函数获取所有环境变量
	for _, e := range os.Environ() {
		// 使用 strings.Split 函数将环境变量的键和值分开
		pair := strings.Split(e, "=")
		// 输出环境变量的键和值
		fmt.Println("Key:", pair[0], "Value:", pair[1])
	}
}
