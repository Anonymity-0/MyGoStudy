package gee

import (
	"net/http"
)

// HandlerFunc 类型定义了一个处理 HTTP 请求的函数签名
type HandlerFunc func(c *Context)

// Engine 结构体是框架的主体，包含一个 router 用于路由映射
type Engine struct {
	router *router
}

// New 函数创建并返回一个新的 Engine 实例
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 方法将一个处理函数添加到 Engine 的路由映射中
// 它使用 HTTP 方法和 URL 路径作为键
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET 方法是一个便捷函数，用于添加一个处理 GET 请求的路由
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 方法是一个便捷函数，用于添加一个处理 POST 请求的路由
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 方法启动 HTTP 服务器，并使用 Engine 的路由映射处理请求
// 如果服务器启动失败，它会返回一个错误
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// ServeHTTP 方法实现了 http.Handler 接口
// 在 HTTP 请求到来时，会创建一个 Context，并将其传递给对应的处理函数
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}
