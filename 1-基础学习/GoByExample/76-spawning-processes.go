package main

import (
	"fmt"       // 用于格式化输出
	"io/ioutil" // 用于读取数据
	"os/exec"   // 用于创建和管理外部进程
)

func main() {
	// 使用 exec.Command 函数创建一个表示 "date" 命令的 *exec.Cmd
	dateCmd := exec.Command("date")

	// 使用 cmd.Output 方法运行 "date" 命令并收集其输出
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	// 输出 "date" 命令的输出
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 使用 exec.Command 函数创建一个表示 "grep hello" 命令的 *exec.Cmd
	grepCmd := exec.Command("grep", "hello")

	// 使用 cmd.StdinPipe 方法获取 "grep hello" 命令的标准输入的管道
	grepIn, _ := grepCmd.StdinPipe()
	// 使用 cmd.StdoutPipe 方法获取 "grep hello" 命令的标准输出的管道
	grepOut, _ := grepCmd.StdoutPipe()
	// 使用 cmd.Start 方法开始 "grep hello" 命令，但不等待其完成
	grepCmd.Start()
	// 向 "grep hello" 命令的标准输入写入数据
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	// 关闭 "grep hello" 命令的标准输入的管道
	grepIn.Close()
	// 从 "grep hello" 命令的标准输出读取数据
	grepBytes, _ := ioutil.ReadAll(grepOut)
	// 使用 cmd.Wait 方法等待 "grep hello" 命令完成
	grepCmd.Wait()

	// 输出 "grep hello" 命令的输出
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 使用 exec.Command 函数创建一个表示 "ls -a -l -h" 命令的 *exec.Cmd
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	// 使用 cmd.Output 方法运行 "ls -a -l -h" 命令并收集其输出
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	// 输出 "ls -a -l -h" 命令的输出
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
