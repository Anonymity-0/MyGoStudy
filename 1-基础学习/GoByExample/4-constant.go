package main

import (
	"fmt"
	"math"
)

// const 用于声明一个常量，可以是字符、字符串、布尔值或数值
const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500000000

	// 常量表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	// 数值常量没有类型，直到被给定一个类型，比如显式的类型转换
	fmt.Println(int64(d))

	// 一个数值可以获得一个类型，比如 math.Sin 函数需要一个 float64 的参数
	// n会被自动转换为float64
	fmt.Println(math.Sin(n))
}
