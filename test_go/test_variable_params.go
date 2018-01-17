package main

import "fmt"

func fun(a int, b ...int) int {
	sum := a
	for _, c := range b {
		sum += c
	}
	return sum
}

func main() {
	fmt.Println(fun(1, 2, 3, 4))
}
