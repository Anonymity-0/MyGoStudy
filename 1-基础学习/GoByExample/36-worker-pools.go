package main

import (
	"fmt"
	"time"
)

// worker 函数定义了一个工作函数，它从 jobs channel 接收任务，处理任务，然后将结果发送到 results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // 从 jobs channel 接收任务，当 jobs channel 被关闭时，循环会结束
		fmt.Println("worker", id, "processing job", j) // 打印工作信息
		time.Sleep(time.Second)                        // 模拟任务处理时间
		fmt.Println("worker", id, "finished job", j)   // 打印工作完成信息
		results <- j * 2                               // 将处理结果发送到 results channel
	}
}

func main() {
	const numJobs = 5                  // 定义任务数量
	jobs := make(chan int, numJobs)    // 创建一个可以存储 numJobs 个 int 类型值的 channel
	results := make(chan int, numJobs) // 创建一个可以存储 numJobs 个 int 类型值的 channel

	for w := 1; w <= 3; w++ { // 启动 3 个 worker goroutine
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ { // 将任务发送到 jobs channel
		jobs <- j
	}
	close(jobs) // 关闭 jobs channel，表示所有的任务都已经发送

	for a := 1; a <= numJobs; a++ { // 从 results channel 接收结果，如果 results channel 没有数据，这会阻塞，直到有数据可接收
		<-results
	}
}
