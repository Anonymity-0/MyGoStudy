package main

import "fmt"

func main() {
	//简单if else
	if 7%2 == 0 {
		fmt.Println("7是偶数")
	} else {
		fmt.Println("7是奇数")
	}
	//可以不要else
	if 8%4 == 0 {
		fmt.Println("8可以被4整除")
	}
	// 条件语句可以在条件表达式之前执行一个语句，变量在整个if else语句块中都可用
	if num := 9; num < 0 {
		fmt.Println(num, "是负数")
	} else if num < 10 {
		fmt.Println(num, "是一位数")
	} else {
		fmt.Println(num, "是多位数")
	}

}
