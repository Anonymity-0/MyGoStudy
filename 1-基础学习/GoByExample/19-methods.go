package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

// 这里的 area 方法有一个接收器类型 *rect。
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

/*
当你使用 *rect 作为接收器类型时，你在方法内部可以修改 rect 的值。这是因为 Go 语言中所有的函数参数都是值传递，当你使用一个指针作为接收器时，
你实际上是在操作原始值的引用，因此你可以修改原始值。

当你使用 rect 作为接收器类型时，你在方法内部不能修改 rect 的值。这是因为你在操作的是 rect 的一个副本，而不是原始值。
*/
func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go 自动处理方法调用时的值和指针之间的转化。
	// 你可以使用指针来调用方法来避免在方法调用时产生一个拷贝，或者让方法能够改变接收的结构体。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
