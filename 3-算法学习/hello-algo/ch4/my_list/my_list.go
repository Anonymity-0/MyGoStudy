package main

type myList struct {
	arrCapacity int   // 列表容量
	arr         []int // 数组
	arrSize     int   //列表大小
	extendRatio int   //扩容比例
}

// 创建列表
// return: 列表
func newMyList() *myList {
	return &myList{
		arrCapacity: 10,
		arr:         make([]int, 10),
		arrSize:     0,
		extendRatio: 2,
	}
}

// 获取列表长度
// return: 列表长度
func (l *myList) size() int {
	return l.arrSize
}

// 获取列表容量
// return: 列表容量
func (l *myList) capacity() int {
	return l.arrCapacity
}

// 访问元素
// index: 索引
func (l *myList) getItem(index int) int {
	if index >= l.arrSize || index < 0 {
		panic("index out of range")
	}
	return l.arr[index]
}

// 更新元素
// index: 索引
// val: 值
func (l *myList) setItem(index int, val int) {
	if index >= l.arrSize || index < 0 {
		panic("index out of range")
	}
	l.arr[index] = val
}

// 在尾部追加元素
// val: 值
func (l *myList) append(val int) {
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	l.arr[l.arrSize] = val
	l.arrSize++
}

// 在中间插入元素
func (l *myList) insert(index int, val int) {
	if l.arrSize == l.arrCapacity {
		l.extendCapacity()
	}
	if index < 0 || index > l.arrSize {
		panic("index out of range")
	}
	for j := l.arrSize - 1; j >= index; j-- {
		l.arr[j+1] = l.arr[j]
	}
	l.arr[index] = val
	l.arrSize++
}

// 删除元素
func (l *myList) delete(index int) int {
	if index < 0 || index >= l.arrSize {
		panic("index out of range")
	}
	num := l.arr[index]
	for j := index; j < l.arrSize-1; j++ {
		l.arr[j] = l.arr[j+1]
	}
	l.arrSize--
	return num
}

// 扩容
// make函数用于创建切片、映射（map）或通道（chan）。在这里，它被用来创建一个新的整数切片。make函数的第一个参数是要创建的类型，第二个参数是切片的长度。在这里，切片的长度是l.arrCapacity*l.extendRatio。
// ...操作符用于将一个切片的所有元素作为参数传递给函数。在这里，它将make函数创建的切片的所有元素作为参数传递给append函数。
func (l *myList) extendCapacity() {
	l.arr = append(l.arr, make([]int, l.arrCapacity*l.extendRatio)...)
	l.arrCapacity = len(l.arr)
}

// 返回有效长度的列表
func (l *myList) toArray() []int {
	return l.arr[:l.arrSize]
}
func main() {

}
