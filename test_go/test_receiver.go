package main

import (
	"fmt"
)

type User struct {
	Name string
}

// receiver是对象，则是临时对象
func (u User) EchoObject() {
	u.Name = "ahaha"
	fmt.Printf("receiver: %p name: %s\n", &u, u.Name)
}

// receiver是指针，则是引用
func (u *User) EchoPointer() {
	u.Name = "hahah"
	fmt.Printf("receiver: %p name: %s\n", u, u.Name)
}

// 对receiver的替换不会起作用
func (u *User) Rename() {
	user := User{
		"stupig",
	}
	fmt.Printf("receiver: %p temp: %p name: %s\n", u, &user, u.Name)
	// assign to receiver directly
	// do not work: assignment to method receiver propagates only to callees but not to callers
	u = &user
	fmt.Printf("receiver: %p temp: %p name: %s\n", u, &user, u.Name)
}

func main() {
	u := User{Name: "onepiece"}
	fmt.Println("---echo object---")
	fmt.Printf("receiver: %p  name: %s\n", &u, u.Name)
	u.EchoObject()
	fmt.Printf("receiver: %p  name: %s\n", &u, u.Name)

	fmt.Println("---echo pointer---")
	fmt.Printf("receiver: %p  name: %s\n", &u, u.Name)
	u.EchoPointer()
	fmt.Printf("receiver: %p  name: %s\n", &u, u.Name)

	fmt.Println("---rename---")
	fmt.Printf("receiver: %p  name: %s\n", &u, u.Name)
	u.Rename()
	fmt.Printf("receiver: %p  name: %s\n", &u, u.Name)
}
