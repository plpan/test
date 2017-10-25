package main

import (
    "fmt"
    "time"
)

var c = make(chan int)
var a string

func f() {
    a = "hello, world"
    time.Sleep(1 * time.Second)
    <- c
}

func test1() {
    go f()
    c <- 0
    fmt.Println(a)
}

func main() {
    test1()
}
