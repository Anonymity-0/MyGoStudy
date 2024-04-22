package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 使用 rand.Intn 函数生成一个 [0,100) 范围内的随机整数
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// 使用 rand.Float64 函数生成一个 [0.0,1.0) 范围内的随机浮点数
	fmt.Println(rand.Float64())

	// 生成一个 [5.0,10.0) 范围内的随机浮点数
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 使用当前时间的 Unix 时间戳创建一个新的随机数源 s1
	s1 := rand.NewSource(time.Now().Unix())
	// 使用随机数源 s1 创建一个新的 rand.Rand 实例 r1
	r1 := rand.New(s1)

	// 使用 r1 生成一个 [0,100) 范围内的随机整数
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 创建一个新的随机数源 s2，其种子值为 42
	s2 := rand.NewSource(42)
	// 使用随机数源 s2 创建一个新的 rand.Rand 实例 r2
	r2 := rand.New(s2)
	// 使用 r2 生成一个 [0,100) 范围内的随机整数
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	// 创建一个新的随机数源 s3，其种子值为 42
	s3 := rand.NewSource(42)
	// 使用随机数源 s3 创建一个新的 rand.Rand 实例 r3
	r3 := rand.New(s3)
	// 使用 r3 生成一个 [0,100) 范围内的随机整数
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
