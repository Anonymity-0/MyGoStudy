package main

type arrayStack struct {
	data []int
}

// 创建栈
func newArrayStack() *arrayStack {
	return &arrayStack{
		data: make([]int, 0, 16),
	}
}

func (s *arrayStack) size() int {
	return len(s.data)
}

func (s *arrayStack) isEmpty() bool {
	return len(s.data) == 0
}

// 入栈
func (s *arrayStack) push(val int) {
	s.data = append(s.data, val)
}

func (s *arrayStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

// 出栈
func (s *arrayStack) pop() any {
	val := s.peek()
	s.data = s.data[:len(s.data)-1]
	return val
}

// 获取sclice用于打印
func (s *arrayStack) toSlice() []int {
	return s.data
}

func main() {

}
