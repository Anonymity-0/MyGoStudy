package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write", i, "as")
	// Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case
	// Go 的 switch 语句会自动终止，除非使用 fallthrough 语句强制执行后面的 case 语句
	switch i {
	case 1:
		fmt.Println("一")
	case 2:
		fmt.Println("二")
	case 3:
		fmt.Println("三")
	}

	// 一个 case 语句中，你可以使用逗号来分隔多个表达式
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("今天是周末")
	default:
		fmt.Println("今天是工作日")
	}

	t := time.Now()
	// 没有条件的 switch 同 switch true 一样
	// 这一构造使得可以用更清晰的形式来编写长的 if-then-else 链
	switch {
	case t.Hour() < 12:
		fmt.Println("现在是上午")
	default:
		fmt.Println("现在是下午")
	}

	// 类型开关
	// 一个类型开关比较类型而非值，可以用来发现一个接口值的类型
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("我是一个布尔值")
		case int:
			fmt.Println("我是一个整数")
		default:
			fmt.Printf("我是一个 %T", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
