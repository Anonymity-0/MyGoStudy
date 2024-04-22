package main

import "fmt"

// 这个函数将递归地调用自身，直到 n 等于 0
// 计算 n 的阶乘
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("fact(7):", fact(7))

	// 闭包也是递归的，但要求定义闭包之前用var声明一个函数变量
	var fib func(n int) int

	// 闭包函数
	// 定义一个匿名函数 fib，它接受一个整型参数 n
	fib = func(n int) int {
		// 如果 n 小于 2，直接返回 n
		// 这是递归的基本情况，也就是递归结束的条件
		// 对于斐波那契数列来说，第 0 个和第 1 个数都是自身
		if n < 2 {
			return n
		} else {
			// 如果 n 大于等于 2，那么返回 fib(n-1) + fib(n-2)
			// 这是递归的递归情况，也就是递归的主体部分
			// 对于斐波那契数列来说，第 n 个数是前两个数的和
			return fib(n-1) + fib(n-2)
		}
	}
	// 递归调用闭包
	fmt.Println("fib(7):", fib(7))

}
