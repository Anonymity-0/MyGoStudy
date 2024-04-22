package main

import (
	"fmt"
	"sort"
)

// byLength 类型是一个字符串切片，它实现了 sort.Interface 接口的三个方法：Len、Swap 和 Less。Len 方法返回切片的长度，Swap 方法交换切片中的两个元素，Less 方法比较切片中的两个元素。sort.Sort 函数使用这三个方法来对字符串切片进行排序。这样，我们就可以自定义排序规则，例如在这个例子中，我们按照字符串的长度进行排序。
// byLength 类型是一个字符串切片
type byLength []string

// Len 方法返回切片的长度
func (s byLength) Len() int {
	return len(s)
}

// Swap 方法交换切片中的两个元素
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less 方法比较切片中的两个元素，如果第一个元素的长度小于第二个元素的长度，返回 true
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"} // 定义一个字符串切片
	sort.Sort(byLength(fruits))                   // 使用 sort.Sort 函数和 byLength 类型的方法来对字符串切片进行排序
	fmt.Println(fruits)                           // 打印排序后的字符串切片，输出：[kiwi peach banana]
}
