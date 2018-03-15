package main

import (
	"fmt"
	"reflect"
	"strings"
)

// 基于reflect.MakeFunc实现泛型编程
// 根据参数类型选择合适的计算函数

func add(args []reflect.Value) (result []reflect.Value) {
	if len(args) == 0 {
		return nil
	}

	var res reflect.Value
	switch args[0].Kind() {
	case reflect.Int:
		n := 0
		for _, a := range args {
			n += int(a.Int())
		}
		res = reflect.ValueOf(n)
	case reflect.String:
		ss := make([]string, 0, len(args))
		for _, arg := range args {
			ss = append(ss, arg.String())
		}
		res = reflect.ValueOf(strings.Join(ss, " "))
	default:
		panic("unsupported element type")
	}

	result = append(result, res)
	return
}

func makeAdd(f interface{}) {
	fn := reflect.ValueOf(f).Elem()
	nf := reflect.MakeFunc(fn.Type(), add)
	fn.Set(nf)
}

func main() {
	var intAdd func(a, b int) int
	var stringAdd func(a, b string) string

	// must be pointer type
	makeAdd(&intAdd)
	makeAdd(&stringAdd)

	fmt.Println(intAdd(1, 2))
	fmt.Println(stringAdd("1", "2"))
}
