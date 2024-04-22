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
	// 创建一个名为 "subdir" 的目录，权限设置为 0755
	err := os.Mkdir("subdir", 0755)
	check(err)

	// defer 关键字用于确保 os.RemoveAll("subdir") 会在 main 函数结束时执行，删除 "subdir" 目录及其所有子目录和文件
	defer os.RemoveAll("subdir")

	// 定义一个函数 createEmptyFile，用于创建一个空文件
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	// 创建一个空文件 "subdir/file1"
	createEmptyFile("subdir/file1")

	// 创建一个目录 "subdir/parent/child"，如果上级目录不存在，会一并创建
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	// 创建三个空文件
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// 读取 "subdir/parent" 目录下的所有文件和目录
	c, err := os.ReadDir("subdir/parent")
	check(err)

	// 列出 "subdir/parent" 目录下的所有文件和目录
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 改变当前工作目录到 "subdir/parent/child"
	err = os.Chdir("subdir/parent/child")
	check(err)

	// 读取当前目录下的所有文件和目录
	c, err = os.ReadDir(".")
	check(err)

	// 列出当前目录下的所有文件和目录
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 改变当前工作目录回到原来的目录
	err = os.Chdir("../../..")
	check(err)

	// 遍历 "subdir" 目录及其所有子目录和文件
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// visit 函数会被 filepath.Walk 调用，每找到一个文件或目录，就会调用一次 visit 函数
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	// 输出文件或目录的路径和是否是目录
	fmt.Println(" ", p, info.IsDir())
	return nil
}
