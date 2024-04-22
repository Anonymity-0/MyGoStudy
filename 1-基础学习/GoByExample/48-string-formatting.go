package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	// %v 以默认的格式输出结构体
	fmt.Printf("struct1: %v\n", p)
	// %+v 会包括结构体的字段名
	fmt.Printf("struct2: %+v\n", p)
	// %#v 输出这个值的 Go 语法表示
	fmt.Printf("struct3: %#v\n", p)

	// %T 输出值的类型
	fmt.Printf("type: %T\n", p)

	// %t 格式化布尔值
	fmt.Printf("bool: %t\n", true)

	// %d 格式化整型
	fmt.Printf("int: %d\n", 123)
	// %b 输出二进制表示
	fmt.Printf("bin: %b\n", 123)
	// %c 输出对应的字符
	fmt.Printf("char: %c\n", 123)
	// %x 输出十六进制表示
	fmt.Printf("hex: %x\n", 123)
	// %f 输出浮点数
	fmt.Printf("float1: %f\n", 78.9)

	// %e 和 %E 输出科学记数法表示的浮点数
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// %s 输出字符串
	fmt.Printf("str1: %s\n", "\"string\"")
	// %q 输出带双引号的字符串
	fmt.Printf("str2: %q\n", "\"string\"")
	// %x 输出字符串的十六进制表示
	fmt.Printf("str3: %x\n", "hex this")

	// %p 输出指针
	fmt.Printf("point: %p\n", &p)

	// 可以用宽度和精度来控制输出的格式
	// 默认宽度是右对齐的，左边加空格
	fmt.Printf("witdth1: |%6d|%6d|\n", 12, 345)
	// 宽度为 0 时，不会填充空格
	fmt.Printf("witdth2: |%6.2f|%6.2f|\n", 1.2, 3.45)
	// 宽度为 - 时，左对齐
	fmt.Printf("witdth3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("witdth4: |%-6s|%-6s|\n", "foo", "b")
	fmt.Printf("witdth5: |%6s|%6s|\n", "foo", "b")

	// Sprintf 格式化并返回一个字符串而不带任何输出
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)
}