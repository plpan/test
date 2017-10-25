package main

import (
    "fmt"
)

func main() {
    iVal := 290
    sVal := "280"
    value1, ok := (interface{}(iVal).(int))
    fmt.Println(value1, ok)
    value2, ok := (interface{}(sVal).(string))
    fmt.Println(value2, ok)
}
