package main

import "fmt"

//嵌入（embedding）是一种简单的代码复用方式。
//通过嵌入，一个结构体可以直接使用另一个结构体的所有字段和方法，就像它们是自己的一样。

// 定义一个名为 base 的结构体，它有一个整型字段 num
type base struct {
	num int
}

// 在 base 结构体上定义一个方法 describe，它返回一个字符串
// 这个方法使用 fmt.Sprintf 格式化一个字符串，包含 base 的 num 字段的值
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// 定义一个名为 container 的结构体，它有一个 base 字段和一个字符串字段 str
// 这里的 base 字段是匿名的，这意味着 container 结构体继承了 base 的所有字段和方法
type container struct {
	base
	str string
}

func main() {
	// 创建一个 container 实例 co，base 字段的 num 为 42，str 字段为 "some name "
	co := container{
		base: base{num: 42},
		str:  "some name ",
	}
	// 打印 co 的 num 和 str 字段的值
	fmt.Printf("co={num:%v, str:%v}\n", co.num, co.str)

	// 打印 co 的 base 字段的 num 值
	fmt.Println("also", co.base.num)

	// 调用 co 的 describe 方法，这个方法实际上是继承自 base 结构体的
	fmt.Println("describe:", co.describe())

	// 定义一个名为 describer 的接口，它有一个方法 describe
	type describer interface {
		describe() string
	}

	// 定义一个 describer 类型的变量 d
	var d describer
	// 将 co 赋值给 d，因为 co 实现了 describer 接口的所有方法
	d = co
	// 调用 d 的 describe 方法，实际上调用的是 co 的 describe 方法
	fmt.Println("d:", d.describe())
}
