package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// 定义一个字符串 s，其内容是一个 URL
	s := "postgres://user:pawd@localhost:5432/path?k=v#f"

	// 使用 url.Parse 函数将字符串 s 解析为一个 url.URL 结构体实例 u
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	// 打印 u 的 Scheme 字段，即 URL 的协议部分
	fmt.Println("scheme:", u.Scheme)
	// 打印 u 的 User 字段，即 URL 的用户信息部分
	fmt.Println("user:", u.User)
	// 打印 u 的 User 字段的 Username 方法的返回值，即 URL 的用户名部分
	fmt.Println("user name:", u.User.Username())
	// 打印 u 的 User 字段的 Password 方法的返回值，即 URL 的密码部分
	p, _ := u.User.Password()
	fmt.Println("password:", p)

	// 打印 u 的 Host 字段，即 URL 的主机部分
	fmt.Println("host:", u.Host)
	// 使用 net.SplitHostPort 函数将 u 的 Host 字段分割为主机名和端口号
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host:", host)
	fmt.Println("port:", port)

	// 打印 u 的 Path 字段，即 URL 的路径部分
	fmt.Println("path:", u.Path)
	// 打印 u 的 Fragment 字段，即 URL 的片段标识符部分
	fmt.Println("fragment:", u.Fragment)

	// 打印 u 的 RawQuery 字段，即 URL 的查询参数部分
	fmt.Println("raw query:", u.RawQuery)
	// 使用 url.ParseQuery 函数将 u 的 RawQuery 字段解析为一个映射 m
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("query params:", m)
	// 打印映射 m 中 "k" 键对应的值
	fmt.Println("query params k:", m["k"][0])
}
