package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()
	queue.PushBack(1)
	queue.PushBack(2)
	queue.PushBack(3)
	queue.PushBack(4)
	queue.PushBack(5)

	//访问队首元素
	front := queue.Front()
	fmt.Println("front:", front.Value)

	//弹出队首元素
	pop := queue.Front()
	queue.Remove(pop)

	//获取队列长度
	size := queue.Len()
	fmt.Println("size:", size)

	//判断是否为空
	isEmpty := queue.Len() == 0
	fmt.Println("isEmpty:", isEmpty)

}
