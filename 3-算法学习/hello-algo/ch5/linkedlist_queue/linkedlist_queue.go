package main

import "container/list"

type linkedListQueue struct {
	data *list.List
}

// 创建队列
func newLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{
		data: list.New(),
	}
}

// 入队
func (q *linkedListQueue) push(val int) {
	q.data.PushBack(val)
}

// 出队
func (q *linkedListQueue) pop() any {
	if q.isEmpty() {
		return nil
	}
	e := q.data.Front()
	q.data.Remove(e)
	return e.Value
}

// 获取队首元素
func (q *linkedListQueue) peek() any {
	if q.isEmpty() {
		return nil
	}
	return q.data.Front().Value
}
func (q *linkedListQueue) size() int {
	return q.data.Len()
}

func (q *linkedListQueue) isEmpty() bool {
	return q.data.Len() == 0
}

// 获取list用于打印
func (q *linkedListQueue) toList() *list.List {
	return q.data
}

func main() {

}
