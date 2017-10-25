package main

import (
    "fmt"
    "sync"
)

func gen(done <-chan struct{}, nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            select {
            case out <- n:
            case <-done:
                return
            }
        }
    }()
    return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range(in) {
            select {
            case out <- n * n:
            case <-done:
                return
            }
        }
    }()
    return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    wg.Add(len(cs))

    output := func(c <-chan int) {
        for n := range c {
            select {
            case out <- n:
            case <-done:
            }
        }
        wg.Done()
    }

    for _, c := range cs {
        go output(c)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}

func main() {
    done := make(chan struct{})
    defer close(done)

    in := gen(done, 2, 3)

    c1 := sq(done, in)
    c2 := sq(done, in)

    output := merge(done, c1, c2)
    fmt.Println(<-output)
}
