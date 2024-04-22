package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// 遍历数组，第一个返回值是索引，第二个是该索引的值，_表示忽略索引
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 遍历数组
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	// 遍历map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// 遍历键
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 遍历字符串
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
