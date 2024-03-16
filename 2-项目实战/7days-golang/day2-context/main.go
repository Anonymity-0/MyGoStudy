// main 包是程序的入口
package main

import (
	"example/gee" // 导入自定义的 gee HTTP 框架
	"net/http"    // 用于处理 HTTP 请求和响应
)

// main 函数是程序的入口点
func main() {
	// 创建一个新的 gee Engine 实例
	r := gee.New()

	// 添加一个处理 GET 请求的路由，当访问 "/" 时，返回一个 HTML 响应
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
	})
	// 添加一个处理 POST 请求的路由，当访问 "/login" 时，返回一个 JSON 响应
	// JSON 数据包含从表单数据中获取的 "username" 和 "password"
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// 启动 HTTP 服务器，监听并在 9999 端口上接收请求
	r.Run(":9999")
}
