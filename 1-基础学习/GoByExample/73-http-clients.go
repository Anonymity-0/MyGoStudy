package main

import (
	"bufio"    // 用于读取数据
	"fmt"      // 用于格式化输出
	"net/http" // 用于发送 HTTP 请求
)

func main() {
	// 使用 http.Get 函数发送 HTTP GET 请求
	resp, err := http.Get("http://gobyexample.com")
	// 如果发生错误，就使用 panic 函数抛出错误
	if err != nil {
		panic(err)
	}
	// 使用 defer 语句确保响应体在函数返回时关闭
	defer resp.Body.Close()

	// 输出响应状态码
	fmt.Println("Response status:", resp.Status)

	// 使用 bufio.NewScanner 函数创建一个新的扫描器，用于读取响应体
	scanner := bufio.NewScanner(resp.Body)
	// 使用 for 循环和 scanner.Scan 方法读取响应体的前 5 行
	for i := 0; scanner.Scan() && i < 5; i++ {
		// 使用 scanner.Text 方法获取当前行的文本，并输出
		fmt.Println(scanner.Text())
	}

	// 使用 scanner.Err 方法检查扫描过程中是否发生错误
	if err := scanner.Err(); err != nil {
		// 如果发生错误，就使用 panic 函数抛出错误
		panic(err)
	}
}
