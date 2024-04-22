package main

import "fmt"

// 这个函数使用任意数目的整数作为参数
// 通过 ...type 的形式来指定可变参数
func sum(nums ...int) int {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// 可变参数函数的调用方式和其他函数一样
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	// 如果你已经有多个参数在一个 slice 中，应用 func(slice...) 来将它们应用到一个变参函数中
	sum(nums...)
}
