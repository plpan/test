package main

import (
	"bytes"
	"fmt"
	"reflect"
)

type User struct {
	Age   int
	Name  string
	Grade int
	Class string
}

func StructIter(u *User) string {
	var buf bytes.Buffer

	uType := reflect.TypeOf(u).Elem()
	uValue := reflect.ValueOf(u)

	elem := uValue.Elem()
	for i := 0; i < elem.NumField(); i++ {
		buf.WriteString(uType.Field(i).Name)
		buf.WriteString("\t")
		buf.WriteString(elem.Field(i).Kind().String())
		buf.WriteString("\t")
	}

	return buf.String()
}

func main() {
	u := User{1, "stupig", 2, "2"}
	res := StructIter(&u)
	fmt.Println(res)
}
