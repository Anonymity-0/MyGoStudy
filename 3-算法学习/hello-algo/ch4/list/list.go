package main

import (
	"fmt"
	"sort"
)

func main() {
	//无初始值的列表
	nums := []int{}
	//有初始值的列表
	nums1 := []int{1, 2, 3, 4, 5, 6}

	// 访问元素
	num := nums1[0]
	println(num)

	//清空列表
	nums1 = nil

	//追加元素
	nums1 = append(nums1, 1)
	//在尾部追加多个元素
	nums1 = append(nums1, 2, 3, 4, 5)

	//在中间插入元素
	nums1 = append(nums1[:1], append([]int{6}, nums1[1:]...)...)

	//删除元素
	nums1 = append(nums1[:1], nums1[2:]...)

	//遍历列表1
	for i := 0; i < len(nums1); i++ {
		println(nums1[i])
	}
	//遍历列表2
	for _, num := range nums {
		println(num)
	}
	//拼接列表
	nums3 := append(nums, nums1...)

	//排序列表
	sort.Ints(nums3)

	for _, num := range nums3 {
		fmt.Println(num)
	}
}
