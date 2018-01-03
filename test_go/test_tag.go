package main

import (
	"fmt"
	"reflect"
)

const tagName = "validate"

// tag 可以借助反射获取
type User struct {
	Id    int    `validate:"-"`
	Name  string `validate:"presence,min=2,max=10"`
	Email string `validate:"email,required"`
}

func main() {
	user := User{
		1,
		"bilibili",
		"bili@163.com",
	}

	t := reflect.TypeOf(user)

	fmt.Println("Type: ", t.Name())
	fmt.Println("Kind: ", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get(tagName)

		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}
