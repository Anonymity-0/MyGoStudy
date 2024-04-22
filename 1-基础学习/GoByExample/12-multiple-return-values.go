package main

import "fmt"

func vals() (int, int) {
	return 1, 2
}

func main() {
	// 接收多个返回值
	a, b := vals()
	fmt.Println("a:", a, "b:", b)

	// 只接收其中一个返回值
	_, c := vals()
	fmt.Println("c:", c)

}
