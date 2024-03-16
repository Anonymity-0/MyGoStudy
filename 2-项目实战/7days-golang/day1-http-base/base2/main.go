package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine 是一个空结构体，我们将在其上定义 ServeHTTP 方法
type Engine struct{}

// ServeHTTP 方法实现了 http.Handler 接口
// 它根据请求的 URL 路径，返回不同的响应
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		// 当请求的 URL 路径为 "/" 时，返回 URL 路径
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "hello":
		// 当请求的 URL 路径为 "/hello" 时，返回请求的 Header
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}

}

func main() {
	// 启动 HTTP 服务器，监听 9999 端口
	// 如果服务器启动失败（例如端口已被占用），log.Fatal 会打印错误信息并退出程序
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
