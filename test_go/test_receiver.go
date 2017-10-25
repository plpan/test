package main

import (
    "fmt"
)

type User struct {
    Name string
}

func (u *User) Echo() {
    fmt.Println(u.Name)
    u.Name = "hahah"
}

/*func (u User) Echo() {
    fmt.Println(u.Name)
    u.Name = "ahaha"
}*/

func main() {
    u := User{Name: "onepiece"}
    u.Echo()
    fmt.Println(u.Name)
}
