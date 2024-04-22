package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 使用 strconv.ParseFloat 函数将字符串 "1.234" 解析为 float64 类型的数字
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 使用 strconv.ParseInt 函数将字符串 "123" 解析为 int64 类型的数字
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// 使用 strconv.ParseInt 函数将十六进制的字符串 "0x1c8" 解析为 int64 类型的数字
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// 使用 strconv.ParseUint 函数将字符串 "789" 解析为 uint64 类型的数字
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// 使用 strconv.Atoi 函数将字符串 "135" 解析为 int 类型的数字
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 尝试使用 strconv.Atoi 函数将非数字的字符串 "wat" 解析为 int 类型的数字
	// 这将失败，并返回一个错误 e
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
