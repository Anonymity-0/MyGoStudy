package gee

import (
	"fmt"
	"net/http"
)

// 首先定义了类型HandlerFunc，这是提供给框架用户的，用来定义路由映射的处理方法。
type HandleFunc func(http.ResponseWriter, *http.Request)

// Engine 结构体包含一个路由映射，用于存储 URL 路径和对应的处理函数
type Engine struct {
	router map[string]HandleFunc
}

// New 是 Engine 的构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// addRoute 方法将一个处理函数添加到 Engine 的路由映射中
// 它使用 HTTP 方法和 URL 路径作为键

func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 方法是一个便捷函数，用于添加一个处理 GET 请求的路由
func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 方法是一个便捷函数，用于添加一个处理 POST 请求的路由
func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.addRoute("POST", pattern, handler)
}

// // RUN 方法启动 HTTP 服务器，并使用 Engine 的路由映射处理请求
// 如果服务器启动失败，它会返回一个错误
func (engine *Engine) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 方法实现了 http.Handler 接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}
