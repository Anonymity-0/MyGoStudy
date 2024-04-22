package main

import "fmt"

// 这个 intSeq 函数返回另一个在 intSeq 函数体内定义的匿名函数
// 这个返回的函数使用闭包的方式隐藏变量 i
// 闭包是一个函数值，它引用了其函数体之外的变量
// 本例中，我们使用闭包来隐藏变量 i

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// 我们调用 intSeq 函数，将返回值（一个函数）赋给 nextInt
	nextInt := intSeq()

	// 通过多次调用 nextInt 来看看闭包的效果
	// nextInt 的值是一个函数，每次调用都会返回一个递增的数字
	// 但是这个数字是隐藏在 intSeq 函数内部的局部变量 i
	// 这个局部变量对于 intSeq 函数来说是私有的
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())
	fmt.Println("nextInt:", nextInt())

	newInts := intSeq()
	fmt.Println("newInts:", newInts())
}
