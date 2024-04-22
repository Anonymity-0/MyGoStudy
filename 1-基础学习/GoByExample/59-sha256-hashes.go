package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// 定义一个字符串 s
	s := "sha256 this string"
	// 使用 sha256.New 函数创建一个新的 hash.Hash 实例 h
	h := sha256.New()

	// 使用 h 的 Write 方法将字符串 s 的字节切片写入 h
	h.Write([]byte(s))
	// 使用 h 的 Sum 方法生成哈希值，参数 nil 表示不添加任何前缀
	bs := h.Sum(nil)

	// 打印原始字符串 s
	fmt.Println(s)
	// 打印哈希值 bs，%x 格式说明符表示以十六进制格式打印
	fmt.Printf("%x\n", bs)
}
