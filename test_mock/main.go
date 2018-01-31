package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser() *User {
	return &User{
		Name: "user",
		Age:  1,
	}
}

func (u *User) Up() int {
	u.Age++
	fmt.Printf("age: %d\n", u.Age)
	return u.Age
}

type UserUp interface {
	Up() int
}

func Up(u UserUp) int {
	return u.Up()
}

func main() {
	u := NewUser()
	Up(u)
}
