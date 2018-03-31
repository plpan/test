package main

import "fmt"

type AdderHandler struct {
}

func NewAdderHandler() *AdderHandler {
	return &AdderHandler{}
}

func (h *AdderHandler) Add(a int32, b int32) (int32, error) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	return a + b, nil
}
