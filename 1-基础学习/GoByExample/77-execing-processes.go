package main

import (
	"os"      // 用于操作系统相关的操作，如获取环境变量
	"os/exec" // 用于创建和管理外部进程
	"syscall" // 用于低级别的操作系统相关的操作，如执行外部命令
)

func main() {
	// 使用 exec.LookPath 函数查找 "ls" 命令的路径
	binary, lookErr := exec.LookPath("ls")
	// 如果发生错误，就使用 panic 函数抛出错误
	if lookErr != nil {
		panic(lookErr)
	}

	// 创建一个包含 "ls" 命令和其参数的切片
	args := []string{"ls", "-a", "-l", "-h"}

	// 使用 os.Environ 函数获取所有环境变量
	env := os.Environ()

	// 使用 syscall.Exec 函数执行 "ls" 命令
	execErr := syscall.Exec(binary, args, env)
	// 如果发生错误，就使用 panic 函数抛出错误
	if execErr != nil {
		panic(execErr)
	}
}
