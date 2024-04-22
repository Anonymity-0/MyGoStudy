package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 可以使用和数组一样的方式来设置和获取切片的值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len 返回切片的长度
	fmt.Println("len:", len(s))

	// append 可以用来在切片的末尾追加一个或多个值
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 切片也可以被复制
	c := make([]string, len(s))
	// copy 的第二个参数是目标切片，第一个参数是源切片
	copy(c, s)
	fmt.Println("cpy:", c)

	// 切片支持通过 slice[low:high] 语法进行切片操作
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 这个切片从 s[0] 到 s[5]（不包括）的元素
	l = s[:5]
	fmt.Println("sl2:", l)

	// 这个切片从 s[2] 到切片的末尾
	l = s[2:]
	fmt.Println("sl3:", l)

	// 我们可以在一行代码中声明并初始化一个切片变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 切片可以组成多维数据结构
	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
