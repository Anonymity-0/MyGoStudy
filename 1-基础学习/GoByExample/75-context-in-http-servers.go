package main

import (
	"fmt"      // 用于格式化输出
	"net/http" // 用于处理 HTTP 请求和响应
	"time"     // 用于处理时间
)

// hello 函数是一个 HTTP 处理函数，用于处理 "/hello" 路径的请求
func hello(w http.ResponseWriter, req *http.Request) {
	// 使用 req.Context 方法获取请求的上下文
	ctx := req.Context()
	// 输出开始处理请求的信息
	fmt.Println("server: hello handler started")
	// 使用 defer 语句确保在函数返回时输出结束处理请求的信息
	defer fmt.Println("server: hello handler ended")

	// 使用 select 语句等待两个可能发生的事件中的一个
	select {
	// 如果 10 秒后还没有收到请求的取消或超时的通知，就向响应写入 "hello\n"
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	// 如果收到了请求的取消或超时的通知，就输出错误信息并向响应写入错误信息
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// 使用 http.HandleFunc 函数注册 hello 函数为 "/hello" 路径的处理函数
	http.HandleFunc("/hello", hello)
	// 使用 http.ListenAndServe 函数启动 HTTP 服务器
	http.ListenAndServe(":8090", nil)
}
