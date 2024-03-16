package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	Path   string
	Method string
	Params map[string]string

	StatusCode int
}

// newContext 函数创建并返回一个新的 Context 实例
// 它接收一个 HTTP 响应写入器和一个 HTTP 请求作为参数
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 方法从 HTTP 请求的表单数据中获取指定的值
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 方法从 HTTP 请求的 URL 查询参数中获取指定的值
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 方法设置 HTTP 响应的状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 方法设置 HTTP 响应的头部信息
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 方法发送一个文本格式的 HTTP 响应
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 方法发送一个 JSON 格式的 HTTP 响应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 方法发送一个二进制格式的 HTTP 响应
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 方法发送一个 HTML 格式的 HTTP 响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
