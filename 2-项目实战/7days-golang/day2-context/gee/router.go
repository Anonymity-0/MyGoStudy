package gee

import (
	"log"
	"net/http"
)

// router 结构体包含一个处理函数映射，用于存储 URL 路径和对应的处理函数
type router struct {
	handlers map[string]HandlerFunc
}

// newRouter 函数创建并返回一个新的 router 实例
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}

}

// addRoute 方法将一个处理函数添加到 router 的处理函数映射中
// 它使用 HTTP 方法和 URL 路径作为键
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// handle 方法根据请求的 HTTP 方法和 URL 路径，找到并调用对应的处理函数
// 如果找不到处理函数，它会返回一个 404 错误
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusFound, "404 NOT FOUND:%s\n", c.Path)
	}
}
