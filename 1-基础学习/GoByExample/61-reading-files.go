package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// check 函数用于检查错误，如果存在错误则 panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 使用 ioutil.ReadFile 函数读取文件中的所有内容
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// 使用 os.Open 函数打开一个文件，返回一个文件对象
	f, err := os.Open("/tmp/dat")
	check(err)

	// 读取文件的前 5 个字节
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// 将文件的读取位置移动到第 6 个字节，然后读取 2 个字节
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2[:n2]))

	// 将文件的读取位置移动到第 6 个字节，然后至少读取 2 个字节
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3[:n3]))

	// 将文件的读取位置移动到开始位置
	_, err = f.Seek(0, 0)
	check(err)

	// 使用 bufio.NewReader 创建一个新的 Reader，然后预读取前 5 个字节
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 关闭文件
	f.Close()
}
