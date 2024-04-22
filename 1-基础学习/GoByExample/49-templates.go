package main

import (
	"os"
	"text/template"
)

func main() {
	// 创建一个新的模板 t1
	t1 := template.New("t1")
	// 解析模板字符串，{{.}} 是一个占位符，会被后面的参数替换
	t1, err := t1.Parse("Value1 is {{.}}\n")
	if err != nil {
		// 如果解析模板字符串出错，抛出 panic
		panic(err)
	}
	// 再次解析模板字符串，这次使用 template.Must 函数，如果解析出错，它会抛出 panic
	t1 = template.Must(t1.Parse("Value {{.}}\n"))

	// 执行模板，将 "some text"、5 和一个字符串切片作为参数传入，替换模板中的占位符
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Golang",
		"Python",
		"Java",
		"JavaScript",
	})

	// 定义一个函数 Create，它接受一个模板名和一个模板字符串，创建一个新的模板，
	// 解析模板字符串，并返回这个模板
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	// 使用 Create 函数创建一个新的模板 t2
	t2 := Create("t2", "Name is {{.Name}}\n")

	// 执行模板，将一个结构体和一个映射作为参数传入，替换模板中的占位符
	t2.Execute(os.Stdout, struct{ Name string }{"Jane Doe"})
	t2.Execute(os.Stdout, map[string]string{"Name": "Mickey Mouse"})

	// 使用 Create 函数创建一个新的模板 t3，这个模板中包含一个 if-else 结构
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	// 执行模板，将一个非空字符串和一个空字符串作为参数传入，替换模板中的占位符
	t3.Execute(os.Stdout, "non-empty string")
	t3.Execute(os.Stdout, "")

	// 使用 Create 函数创建一个新的模板 t4，这个模板中包含一个 range 结构
	t4 := Create("t4",
		"Range: {{range .}} {{.}} {{end}}\n")
	// 执行模板，将一个字符串切片作为参数传入，替换模板中的占位符
	t4.Execute(os.Stdout, []string{"Golang", "Python", "Java", "JavaScript"})
}