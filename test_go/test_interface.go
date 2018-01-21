package main

import "fmt"

// 接口类型等于nil当且仅当其类型和data都为nil

type Empty interface{}

type Inter interface {
	Print()
}

type Str struct{}

func (Str) Print() {}

func main() {
	var i interface{}
	fmt.Println(i == nil)
	var inter Inter
	fmt.Println(inter == nil)
	var empty Empty
	fmt.Println(empty == nil)

	var str Str

	i = str
	fmt.Println(i == nil)
	empty = str
	// 类型为Str，值为nil，所以不为nil
	fmt.Println(empty == nil)

	// canfail案例测试，接口转换不成功，但是程序仍然继续执行
	var empty1 Empty
	inter, ok := empty1.(Inter)
	fmt.Println(empty == nil, ok)
}
