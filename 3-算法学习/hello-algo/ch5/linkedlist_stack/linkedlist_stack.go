package main

import "container/list"

type linkedListStack struct {
	data *list.List
}

// 创建栈
func newLinkedListStack() *linkedListStack {
	return &linkedListStack{
		data: list.New(),
	}
}

// 入栈
func (s *linkedListStack) push(val int) {
	s.data.PushBack(val)
}

// 出栈
func (s *linkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	return e.Value
}

// 获取栈顶元素
func (s *linkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.data.Back().Value
}

// 获取栈长度
func (s *linkedListStack) size() int {
	return s.data.Len()
}

// 判断是否为空
func (s *linkedListStack) isEmpty() bool {
	return s.data.Len() == 0
}

// 获取list用于打印
func (s *linkedListStack) toList() *list.List {
	return s.data
}
func main() {

}
