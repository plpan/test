/**
 * function f catch panic, after handling this signal, it will be dismissed
 * so function main will continue run
 */
package main

import (
    "fmt"
    "log"
)

func main() {
    defer func() {
        fmt.Println("3")
        if err := recover(); err != nil {
            fmt.Println(err)
        }
        fmt.Println("4")
    }()
    f()
    fmt.Println("9")
}

func f() {
    defer func() {
        fmt.Println("6")
        if err := recover(); err != nil {
            fmt.Println("7")
        }
        fmt.Println("8")
    }()
    fmt.Println("1")
    log.Panic("2")
    fmt.Println("5")
}
