package main

type arrayDeque struct {
	nums          []int //用于存储队列元素的数组
	front         int   //队首指针
	queueSize     int   //队列大小
	queueCapacity int   //队列容量
}

func newArrayDeuqe(queCapacity int) *arrayDeque {
	return &arrayDeque{
		nums:          make([]int, queCapacity),
		front:         0,
		queueSize:     0,
		queueCapacity: queCapacity,
	}
}

// 获取队列长度
func (q *arrayDeque) size() int {
	return q.queueSize
}

// 判断队列是否为空
func (q *arrayDeque) isEmpty() bool {
	return q.size() == 0
}

// 计算环形数组索引
func (q *arrayDeque) index(idx int) int {
	return (idx + q.queueCapacity) % q.queueCapacity
}

// 入队
func (q *arrayDeque) pushFirst(val int) {
	if q.size() == q.queueCapacity {
		panic("queue is full")
	}
	// front指针前移 自减1
	q.front = q.index(q.front - 1)
	// 插入元素
	q.nums[q.front] = val
	// 队列大小自增1
	q.queueSize++
}

// 队尾入队
func (q *arrayDeque) pushLast(val int) {
	if q.size() == q.queueCapacity {
		panic("queue is full")
	}

	rear := q.index(q.front + q.size())
	q.nums[rear] = val
	q.queueSize++
}

func (q *arrayDeque) peekFirst() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

func (q *arrayDeque) peekLast() any {
	if q.isEmpty() {
		return nil
	}
	rear := q.index(q.front + q.queueSize - 1)
	return q.nums[rear]
}

// 队首出队
func (q *arrayDeque) popFirst() any {
	num := q.peekFirst()
	q.front = q.index(q.front + 1)
	q.queueSize--
	return num
}

// 队尾出队
func (q *arrayDeque) popLast() any {
	num := q.peekLast()
	q.queueSize--
	return num
}

// 获取slice用于打印
func (q *arrayDeque) toSlice() []int {
	res := make([]int, q.queueSize)
	for i, j := 0, q.front; i < q.queueSize; i++ {
		res[i] = q.nums[q.index(j)]
		j++
	}
	return res
}

func main() {

}
