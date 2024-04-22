package main

import (
	"bufio" // 用于缓冲 I/O
	"fmt"   // 用于格式化输出
	"os"    // 用于操作系统相关的操作，如文件操作
)

// check 函数用于检查错误，如果有错误就 panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 创建一个字节切片
	d1 := []byte("hello\ngo\n")
	// 将字节切片写入到文件 /tmp/dat1 中，如果文件不存在则创建，文件权限设置为 0644
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err) // 检查错误

	// 创建一个新的文件 /tmp/dat2
	f, err := os.Create("/tmp/dat2")
	check(err) // 检查错误

	// defer 关键字用于确保 f.Close() 会在 main 函数结束时执行
	defer f.Close()

	// 创建一个字节切片
	d2 := []byte{115, 111, 109, 101, 10}
	// 将字节切片写入到文件 f 中
	n2, err := f.Write(d2)
	check(err)                         // 检查错误
	fmt.Printf("wrote %d bytes\n", n2) // 打印写入的字节数

	// 将字符串 "writes\n" 写入到文件 f 中
	n3, err := f.WriteString("writes\n")
	check(err)                         // 检查错误
	fmt.Printf("wrote %d bytes\n", n3) // 打印写入的字节数

	// Sync 方法用于将所有缓冲的文件数据写入到硬盘
	f.Sync()

	// 创建一个新的 Writer 对象
	w := bufio.NewWriter(f)
	// 将字符串 "buffered\n" 写入到 Writer 中
	n4, err := w.WriteString("buffered\n")
	check(err)                         // 检查错误
	fmt.Printf("wrote %d bytes\n", n4) // 打印写入的字节数

	// Flush 方法用于将所有缓冲的数据写入到硬盘
	w.Flush()
}
