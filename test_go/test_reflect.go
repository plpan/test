package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Age int
}

func testType() {
	var a User
	at := reflect.TypeOf(a)
	if at.Kind() != reflect.Struct {
		fmt.Println("false")
	}
}

func testValue() {
	var b User
	bt := reflect.ValueOf(b)
	if bt.Kind() != reflect.Struct {
		fmt.Println("false")
	}
}

func main() {
	testType()
	testValue()
}
