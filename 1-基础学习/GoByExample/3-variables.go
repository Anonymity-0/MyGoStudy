package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	// 一次性声明多个变量
	var b, c int = 1, 2
	fmt.Println("b:", b, "c:", c)

	// Go 会自动推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	// 未初始化的变量会被赋予零值
	var e int
	fmt.Println(e)

	// := 语法是声明并初始化变量的简写
	f := "short"
	fmt.Println(f)
}
