
/*递归实现*/
func recur() int {
	if n == 1 {
		return 1
	}
	res := recur(n - 1)
	return n + res
}

/* 尾递归实现 */
func tailRecur(n int, res int) int {
	//终止条件
	if n == 0 {
		return res
	}
	return tailRecur(n-1, n+res)
}

/* 斐波那契数列 */
func fib(n int) int {
	if n == 1|n == 2 {
		return n - 1
	}

	res := fib(n-1) + fib(n-2)
	return res
}

/* 使用迭代模拟递归 */
func forLoopRecur(n int) int {
	stack := list.New()
	res := 0
	for i := n; i > 0; i-- {
		stack.PushBack(i)
	}
	for stack.Len() > 0 {
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}
	return res
}