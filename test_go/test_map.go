package main

import "fmt"

type Bool bool

/**
 * 用作map key的struct不能包含map、func、slice
 * 因为golang没法用自己定义hash或者重载==运算符
 */
type T1 struct {
	a [3]string
	b map[string]int
}

type T2 struct {
	a [3]string
}

func main() {
	a := make(map[T2]interface{})
	a[T2{}] = "b"
	fmt.Printf("%v\n", a)

	// ERROR
	c := make(map[T1]interface{})
	c[T1{}] = "b"
	fmt.Printf("%v\n", c)
}
