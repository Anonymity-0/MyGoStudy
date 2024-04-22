package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	// 使用给定的名称创建一个新的 person 结构体
	p := person{name: name}
	// 你可以在初始化结构体时指定字段名
	p.age = 42
	//返回一个指向结构体的指针
	return &p
}

func main() {

	fmt.Println(person{"Bob}", 20})
	// 你可以通过字段名来初始化结构体
	fmt.Println(person{name: "Alice", age: 30})
	// 省略的字段将被初始化为零值
	fmt.Println(person{name: "Fred"})
	// & 前缀生成一个结构体指针
	fmt.Println(&person{name: "Ann", age: 40})
	// 也可以使用 new 函数来创建一个结构体指针
	fmt.Println(newPerson("Jon"))

	// 访问结构体字段
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 你也可以对结构体指针使用 . 操作符 - 指针会被自动解引用
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
