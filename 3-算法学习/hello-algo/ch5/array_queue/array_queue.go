package main

type arrayQueue struct {
	nums        []int //用于存储队列元素的数组
	front       int   //队首指针
	queSize     int   //队列大小
	queCapacity int   //队列容量
}

func newArrayQueue(queCapacity int) *arrayQueue {
	return &arrayQueue{
		nums:        make([]int, queCapacity),
		front:       0,
		queSize:     0,
		queCapacity: queCapacity,
	}
}

func (q *arrayQueue) size() int {
	return q.queSize
}

func (q *arrayQueue) isEmpty() bool {
	return q.queSize == 0
}

func (q *arrayQueue) push(val int) {
	if q.queSize == q.queCapacity {
		return
	}
	rear := (q.front + q.queSize) % q.queCapacity
	q.nums[rear] = val
	q.queSize++
}

func (q *arrayQueue) peek() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

func (q *arrayQueue) pop() any {
	num := q.peek()
	q.front = (q.front + 1) % q.queCapacity
	q.queSize--
	return num
}

func (q *arrayQueue) toSlice() []int {
	rear := (q.front + q.queSize)
	if rear >= q.queCapacity {
		rear %= q.queCapacity
		return append(q.nums[q.front:], q.nums[:rear]...)
	}
	return q.nums[q.front:rear]
}
func main() {

}
