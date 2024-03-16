package main

import (
	"math/rand"
)

// randomAccess 随机访问元素
// nums: 数组
// return: 随机访问的元素
func randomAccess(nums []int) (randomNum int) {
	// 生成随机数
	randomIndex := rand.Intn(len(nums))
	// 访问元素
	randomNum = nums[randomIndex]
	return
}

/* 在index处插入元素 */
// nums: 数组
// num: 插入的元素
// index: 插入的位置
func insert(nums []int, num int, index int) {
	// 从后往前移动元素
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	// 插入元素
	nums[index] = num
}

/* 删除index处的元素 */
// nums: 数组
// index: 删除的位置
func remove(nums []int, index int) {
	// 从前往后移动元素
	for i := index; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
}

/*遍历数组*/
// nums: 数组
func traverse(nums []int) {
	count := 0
	// 传统的for循环
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	count = 0
	// 使用range关键字
	for _, num := range nums {
		count += num
	}
	count = 0
	// 使用range关键字，同时获取索引
	for i, num := range nums {
		count += nums[i]
		count += num
	}
}

/*查找元素*/
// nums: 数组
// target: 查找的元素
// return: 元素的索引
func find(nums []int, target int) (index int) {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			index = i
			break
		}
	}
	return
}

/*拓展数组长度*/
// nums: 数组
// enlarge: 拓展的长度
// return: 拓展后的数组
func extend(nums []int, enlarge int) []int {
	// 创建新的数组,
	res := make([]int, len(nums)+enlarge)

	// 将原数组的元素复制到新数组
	for i, num := range nums {
		res[i] = num
	}
	return res
}

func main() {
	/* 初始化数组 */
	// 在 Go 中，指定长度时（[5]int）为数组，不指定长度时（[]int）为切片
	//var arr [5]int
	nums := []int{1, 2, 3, 4, 5}

	// 遍历数组
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}

	// 随机访问元素
	randomNum := randomAccess(nums)
	println(randomNum)

}
