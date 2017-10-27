package main

import (
    "fmt"
    "time"
)

func main() {
    t, err := time.Parse("2006-01-02", "2011-01-19")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(t.String())
}
