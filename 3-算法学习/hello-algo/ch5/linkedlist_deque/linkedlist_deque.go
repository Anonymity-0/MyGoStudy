package main

import "container/list"

type linkedListDeque struct {
	data *list.List
}

func newLinkedListDeque() *linkedListDeque {
	return &linkedListDeque{
		data: list.New(),
	}
}

// 入队
func (q *linkedListDeque) pushFirst(value any) {
	q.data.PushFront(value)
}

func (q *linkedListDeque) pushLast(value any) {
	q.data.PushBack(value)
}

func (q *linkedListDeque) isEmpty() bool {
	return q.data.Len() == 0
}

// 出队
func (q *linkedListDeque) popFirst() any {
	if q.isEmpty() {
		return nil

	}
	e := q.data.Front()
	q.data.Remove(e)
	return e.Value
}

func (q *linkedListDeque) popLast() any {
	if q.isEmpty() {
		return nil
	}
	e := q.data.Back()
	q.data.Remove(e)
	return e.Value
}

func (q *linkedListDeque) peekFirst() any {
	if q.isEmpty() {
		return nil

	}
	e := q.data.Front()
	return e.Value
}

func (q *linkedListDeque) peekLast() any {
	if q.isEmpty() {
		return nil

	}
	e := q.data.Back()
	return e.Value
}

func (q *linkedListDeque) size() int {
	return q.data.Len()
}

func (q *linkedListDeque) toList() *list.List {
	return q.data
}
func main() {

}
