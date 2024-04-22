package main

import "fmt"

// 定义一个名为 MapKeys 的泛型函数，它接受一个 map 类型的参数 m
// K 和 V 是类型参数，K 必须是可比较的，V 可以是任何类型
func MapKeys[K comparable, V any](m map[K]V) []K {
	// 创建一个空的切片 r，它的类型是 K
	r := make([]K, 0, len(m))
	// 遍历 map m 的所有键，将它们添加到切片 r 中
	for k := range m {
		r = append(r, k)
	}
	// 返回切片 r，它包含了 map m 的所有键
	return r
}

// 定义一个名为 List 的泛型结构体，它有两个字段：head 和 tail
// T 是类型参数，可以是任何类型
type List[T any] struct {
	head, tail *element[T]
}

// 定义一个名为 element 的泛型结构体，它有两个字段：next 和 value
// T 是类型参数，可以是任何类型
type element[T any] struct {
	next  *element[T]
	value T
}

// 在 List 结构体上定义一个名为 Push 的方法，它接受一个参数 v
// 这个方法将 v 添加到 List 的末尾
func (l *List[T]) Push(v T) {
	if l.tail == nil {
		// 如果 List 是空的，创建一个新的 element 并将其设置为 head 和 tail
		l.head = &element[T]{value: v}
		l.tail = l.head
	} else {
		// 如果 List 不是空的，创建一个新的 element 并将其添加到 tail 的后面
		l.tail.next = &element[T]{value: v}
		l.tail = l.tail.next
	}
}

// 在 List 结构体上定义一个名为 GetAll 的方法，它返回 List 中所有元素的切片
func (l *List[T]) GetAll() []T {
	var elems []T
	// 从 head 开始，遍历 List 中的所有元素
	for e := l.head; e != nil; e = e.next {
		// 将每个元素的值添加到切片 elems 中
		elems = append(elems, e.value)
	}
	// 返回包含了 List 中所有元素的切片
	return elems
}

func main() {
	// 创建一个 map 并调用 MapKeys 函数获取它的所有键
	var m = map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("keys m:", MapKeys(m))

	// 创建一个 List 并添加一些元素
	lst := List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)

	// 调用 GetAll 方法获取 List 中的所有元素
	fmt.Println("list:", lst.GetAll())
}
