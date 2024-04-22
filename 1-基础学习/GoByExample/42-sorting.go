package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"} // 定义一个字符串切片
	sort.Strings(strs)              // 使用 sort.Strings 函数对字符串切片进行排序
	fmt.Println("Strings:", strs)   // 打印排序后的字符串切片，输出：Strings: [a b c]

	ints := []int{7, 2, 4}        // 定义一个整数切片
	sort.Ints(ints)               // 使用 sort.Ints 函数对整数切片进行排序
	fmt.Println("Ints:   ", ints) // 打印排序后的整数切片，输出：Ints:    [2 4 7]

	s := sort.IntsAreSorted(ints) // 使用 sort.IntsAreSorted 函数检查整数切片是否已排序
	fmt.Println("Sorted: ", s)    // 打印检查结果，输出：Sorted:  true
}
