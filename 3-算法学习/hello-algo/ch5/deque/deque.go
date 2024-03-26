package main

import (
	"container/list"
	"fmt"
)

func main() {
	deque := list.New()

	//元素入队
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)
	deque.PushFront(4)
	deque.PushFront(5)

	//
	front := deque.Front()
	rear := deque.Back()

	deque.Remove(front)
	deque.Remove(rear)

	size := deque.Len()

	isEmpty := size == 0

	fmt.Println("isEmpty:", isEmpty)
}
