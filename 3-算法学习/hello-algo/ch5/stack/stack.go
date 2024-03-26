package main

import "fmt"

func main() {
	var stack []int
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)
	stack = append(stack, 5)

	//访问栈顶元素
	peek := stack[len(stack)-1]
	fmt.Println("peek:", peek)

	//弹出栈顶元素
	pop := stack[:len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println("pop:", pop)

	//获取栈长度
	size := len(stack)
	fmt.Println("size:", size)

	//判断是否为空
	isEmpty := len(stack) == 0
	fmt.Println("isEmpty:", isEmpty)
}
