package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// 通过 *int 参数来传递 i 的内存地址
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	// zeroval 有一个 int 型参数，所以使用值传递。 zeroval 将从调用它的那个函数中得到一个实参的拷贝：ival
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i 语法来取得 i 的内存地址，即指向 i 的指针
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 指针也是可以被打印的
	fmt.Println("pointer:", &i)

}
