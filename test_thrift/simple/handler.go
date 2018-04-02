package main

import (
	"fmt"

	"github.com/pplonepiece/test/test_thrift/simple/gen-go/thrift/example"
)

type AdderHandler struct {
}

func NewAdderHandler() *AdderHandler {
	return &AdderHandler{}
}

func (h *AdderHandler) Add(a int32, b int32) (int32, error) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	return a + b, nil
}

func (h *AdderHandler) AddNumber(a *example.Numbers, b *example.Numbers) (*example.Numbers, error) {
	sn := example.NewNumbers()
	sn.A = a.A + b.A
	sn.B = a.B + b.B
	return sn, nil
}
