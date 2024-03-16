package main

import (
	"fmt"
	"time"
)

// main 是程序的入口点。
// 它首先创建一个新的 Trace，并使用它来创建一个新的文件。
// 然后，它获取并执行测试文件中的操作，并计算缓冲区的命中率和输入/输出操作的数量。
// 最后，它打印出这些统计信息和程序的运行时间。
func main() {
	trace := NewTrace()

	trace.CreateFile()

	fmt.Println("data.dbf创建完成")

	startTime := time.Now()
	trace.GetStatistics()
	runTime := time.Since(startTime)

	fmt.Printf("\nTrace程序运行完成,最终测试数据:\nI/O总次数: %d\t命中率: %f\t运行时间: %v\n", trace.IOCounter, trace.HitRate, runTime)
}
