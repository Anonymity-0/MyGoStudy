package main

import (
	"errors"
	"fmt"
)

// 定义一个名为 f1 的函数，它接受一个整型参数 arg 并返回一个整型值和一个 error
func f1(arg int) (int, error) {
	// 如果 arg 等于 42，返回 -1 和一个新的 error
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	// 否则，返回 arg + 3 和 nil
	return arg + 3, nil
}

// 定义一个名为 argError 的结构体，它有两个字段：arg 和 prob
type argError struct {
	arg  int
	prob string
}

// 在 argError 结构体上定义一个 Error 方法，它返回一个字符串
// 这个方法使用 fmt.Sprintf 格式化一个字符串，包含 argError 的 arg 和 prob 字段的值
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// 定义一个名为 f2 的函数，它接受一个整型参数 arg 并返回一个整型值和一个 error
func f2(arg int) (int, error) {
	// 如果 arg 等于 42，返回 -1 和一个 argError
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	// 否则，返回 arg + 3 和 nil
	return arg + 3, nil
}

func main() {
	// 遍历一个整型切片，包含两个元素：7 和 42
	for _, i := range []int{7, 42} {
		// 调用 f1 函数，将返回值赋值给 r 和 e
		if r, e := f1(i); e != nil {
			// 如果 e 不是 nil，打印 "f1 failed:" 和 e 的值
			fmt.Println("f1 failed:", e)
		} else {
			// 如果 e 是 nil，打印 "f1 worked:" 和 r 的值
			fmt.Println("f1 worked:", r)
		}
	}
	// 遍历同样的整型切片，调用 f2 函数
	for _, i := range []int{7, 42} {
		// 调用 f2 函数，将返回值赋值给 r 和 e
		if r, e := f2(i); e != nil {
			// 如果 e 不是 nil，打印 "f2 failed:" 和 e 的值
			fmt.Println("f2 failed:", e)
		} else {
			// 如果 e 是 nil，打印 "f2 worked:" 和 r 的值
			fmt.Println("f2 worked:", r)
		}
	}
}
