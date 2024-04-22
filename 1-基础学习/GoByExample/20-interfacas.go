package main

import "fmt"

// Go 语言中的接口是一种抽象的类型
// 它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合
// 接口只会展示出它自己的方法
// 一个对象只要全部实现了接口中的方法，那么这个对象就实现了这个接口
type geometry interface {
	area() float64
	perim() float64
}

// 在这个例子中，我们将让 rect 和 circle 实现 geometry 接口
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// 要在 Go 中实现一个接口，我们只需要实现接口中的所有方法
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// circle 的实现
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * 3.14 * c.radius
}

// 如果一个变量的是接口类型，那么我们可以调用这个被命名的接口中的方法
func measure(g geometry) {
	fmt.Println("g:", g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}
func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// 结构体类型 circle 和 rect 都实现了 geometry 接口，所以我们可以使用它们的实例作为 measure 函数的参数
	measure(r)
	measure(c)
}
