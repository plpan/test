package main

import "fmt"

// to get the details of compiler and stack, run the command below:
// go tool compiler -S test_return.go
//
// we can learn the stack is arranged below:
//
// the last return value
// ......
// the first return value
// the last parameter
// ......
// the first parameter

func test(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	a, b := test(1, 2)
	fmt.Println(a, b)
}
