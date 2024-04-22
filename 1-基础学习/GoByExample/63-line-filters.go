package main

import (
	"bufio"   // 用于缓冲 I/O
	"fmt"     // 用于格式化输出
	"os"      // 用于操作系统相关的操作，如标准输入和标准错误
	"strings" // 用于操作字符串
)

func main() {
	// 创建一个从标准输入读取的 Scanner
	scanner := bufio.NewScanner(os.Stdin)

	// 循环读取每一行
	for scanner.Scan() {
		// 将每一行的内容转换为大写
		ucl := strings.ToUpper(scanner.Text())
		// 输出转换后的内容
		fmt.Println(ucl)
	}

	// 如果在读取过程中发生错误，输出错误并退出程序
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
