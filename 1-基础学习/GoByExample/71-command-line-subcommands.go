package main

import (
	"flag" // 用于解析命令行参数
	"fmt"  // 用于格式化输出
	"os"   // 用于操作系统相关的操作，如获取命令行参数
)

func main() {
	// 使用 flag.NewFlagSet 函数创建一个新的 flag.FlagSet，用于解析 "foo" 子命令的命令行参数
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// 定义 "foo" 子命令的 "enable" 参数
	fooEnable := fooCmd.Bool("enable", false, "enable")
	// 定义 "foo" 子命令的 "name" 参数
	fooName := fooCmd.String("name", "", "name")

	// 使用 flag.NewFlagSet 函数创建一个新的 flag.FlagSet，用于解析 "bar" 子命令的命令行参数
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	// 定义 "bar" 子命令的 "level" 参数
	barLevel := barCmd.Int("level", 0, "level")

	// 如果命令行参数的数量小于 2，就输出错误信息并退出程序
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// 根据第一个命令行参数判断是哪个子命令
	switch os.Args[1] {
	case "foo":
		// 解析 "foo" 子命令的命令行参数
		fooCmd.Parse(os.Args[2:])
		// 输出 "foo" 子命令的信息
		fmt.Println("subcommand 'foo'")
		// 输出 "foo" 子命令的 "enable" 参数
		fmt.Println("  enable:", *fooEnable)
		// 输出 "foo" 子命令的 "name" 参数
		fmt.Println("  name:", *fooName)
		// 输出 "foo" 子命令的非标志参数
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		// 解析 "bar" 子命令的命令行参数
		barCmd.Parse(os.Args[2:])
		// 输出 "bar" 子命令的信息
		fmt.Println("subcommand 'bar'")
		// 输出 "bar" 子命令的 "level" 参数
		fmt.Println("  level:", *barLevel)
		// 输出 "bar" 子命令的非标志参数
		fmt.Println("  tail:", barCmd.Args())
	default:
		// 如果第一个命令行参数既不是 "foo" 也不是 "bar"，就输出错误信息并退出程序
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
