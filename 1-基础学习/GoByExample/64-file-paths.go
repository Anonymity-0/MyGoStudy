package main

import (
	"fmt"           // 用于格式化输出
	"path/filepath" // 用于处理文件路径
	"strings"       // 用于操作字符串
)

func main() {
	// 使用 filepath.Join 函数连接路径，它会根据操作系统的路径规则来连接
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// filepath.Join 函数会清理路径，如删除多余的 "/"
	fmt.Println(filepath.Join("dir1//", "filename"))
	// filepath.Join 函数会清理路径，如解析 ".."
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// 使用 filepath.Dir 函数获取路径中的目录部分
	fmt.Println("Dir(p):", filepath.Dir(p))
	// 使用 filepath.Base 函数获取路径中的文件名部分
	fmt.Println("Base(p):", filepath.Base(p))

	// 使用 filepath.IsAbs 函数检查路径是否是绝对路径
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// 定义一个文件名
	filename := "config.json"

	// 使用 filepath.Ext 函数获取文件的扩展名
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// 使用 strings.TrimSuffix 函数删除文件名的扩展名
	fmt.Println("strings.TrimSuffix(filename, ext):", strings.TrimSuffix(filename, ext))

	// 使用 filepath.Rel 函数获取两个路径之间的相对路径
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

}
