package main

import "fmt"

// 测试golang闭包
// 闭包的简单定义：有数据的行为
// 闭包的底层数据结构就是一个struct {FuncPtr, *ElemType}
// 其中FuncPtr存储的是闭包函数的地址，而*ElemType存储的是闭包数据的地址
//
// 更信息的信息可以根据以下命令查看
// go tool compiler -S test_closure.go

func test(a int) func(int) int {
	return func(i int) int {
		a = a + i
		return a
	}
}

func main() {
	f := test(1)
	i1 := f(2)
	i2 := f(3)
	fmt.Println(i1, i2)
}
