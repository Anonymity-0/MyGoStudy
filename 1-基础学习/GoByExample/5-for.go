package main

import "fmt"

func main() {
	i := 1
	// for 循环是 Go 中唯一的循环结构
	// 这里有三种基本的循环类型
	// 最基本的循环类型，一个条件
	for i <= 3 {
		println(i)
		i = i + 1
	}

	// 初始化语句；条件表达式；后续语句
	for j := 7; j <= 9; j++ {
		fmt.Println("j:", j)
	}

	// 不带条件的 for 循环将一直执行，直到在循环体内使用了 break 或者 return 来跳出循环
	for {
		fmt.Println("loop")
		break
	}

	// continue 可以跳过当前循环的剩余代码，开始下一次循环
	// 输出奇数
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n:", n)
	}
}
