package main

// LinkNode 链表节点
type LinkNode struct {
	Val  int       // 节点值
	Next *LinkNode // 下一个节点
}

// NewLinkNode 创建链表节点
// val: 节点值
// return: 链表节点
func NewLinkNode(val int) *LinkNode {
	return &LinkNode{
		Val:  val,
		Next: nil,
	}
}

// insertNode 插入节点
// n0: 节点
// P: 插入的节点
func insertNode(n0 *LinkNode, P *LinkNode) {
	P.Next = n0.Next
	n0.Next = P
}

// removeNode 删除n0之后的第一个节点
// n0: 节点
func removeNode(n0 *LinkNode) {
	if n0.Next == nil {
		return
	}
	P := n0.Next
	n0.Next = P.Next
}

// 访问链表中索引为i的节点
// head: 链表头节点
// index: 索引
func access(head *LinkNode, index int) *LinkNode {
	if head == nil {
		return nil
	}
	for i := 0; i < index; i++ {
		head = head.Next
	}
	return head
}

// 查找链表中值为target的节点的索引值
// head: 链表头节点
// target: 查找的值
// return: 索引值，-1表示未找到
func findNode(head *LinkNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}
func main() {
	///* 初始化链表 1 -> 3 -> 2 -> 5 -> 4 */
	n0 := NewLinkNode(1)
	n1 := NewLinkNode(3)
	n2 := NewLinkNode(2)
	n3 := NewLinkNode(5)
	n4 := NewLinkNode(4)

	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
}
