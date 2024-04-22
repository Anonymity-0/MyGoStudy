package main

import (
	"fmt"           // 用于格式化输出
	"os"            // 用于操作系统相关的操作，如文件和目录操作
	"path/filepath" // 用于处理文件路径
)

// check 函数用于检查错误，如果有错误就 panic
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 使用 os.CreateTemp 函数创建一个临时文件，第一个参数是临时文件的目录，如果为空字符串，就使用默认的临时文件目录
	f, err := os.CreateTemp("", "sample")
	check(err)
	// 输出临时文件的路径
	fmt.Println("Temp file created:", f.Name())

	// defer 关键字用于确保 os.Remove(f.Name()) 会在 main 函数结束时执行，删除临时文件
	defer os.Remove(f.Name())

	// 向临时文件中写入数据
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// 使用 os.MkdirTemp 函数创建一个临时目录，第一个参数是临时目录的父目录，如果为空字符串，就使用默认的临时文件目录
	dname, err := os.MkdirTemp("", "sampledir")
	// 输出临时目录的路径
	fmt.Println("Temp dir created:", dname)

	// defer 关键字用于确保 os.RemoveAll(dname) 会在 main 函数结束时执行，删除临时目录及其所有子目录和文件
	defer os.RemoveAll(dname)

	// 使用 filepath.Join 函数连接临时目录的路径和文件名，得到文件的完整路径
	fname := filepath.Join(dname, "file1")
	// 使用 os.WriteFile 函数创建文件并写入数据
	err = os.WriteFile(fname, []byte{1, 2, 3, 4}, 0666)
	check(err)
}
