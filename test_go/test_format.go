package main

import (
    "fmt"
)

type User struct {
    Name string
}

func main() {
    fmt.Printf("%v\n", "sss")
    fmt.Printf("%v\n", 1212)

    user := User{Name: "Ha"}
    fmt.Printf("%v\n", user)
}
