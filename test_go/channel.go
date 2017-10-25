package main

import (
    "fmt"
    "time"
)

func sum(s[] int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
}

func testChannel() {
    s := []int{1, 3, 5, 7, 9}

    c := make(chan int)
    defer close(c)

    go sum(s[:len(s) / 2], c)
    go sum(s[len(s) / 2:], c)

    x, y := <-c, <-c

    fmt.Println(x, y, x + y)
}

func testRange() {
    // this goroutine is used to avoid panic - all goroutines are asleep
    go func() {
        time.Sleep(1 * time.Hour)
    }()

    c := make(chan int)
    go func() {
        for i := 0; i < 10; i = i + 1 {
            c <- i
        }
        // if close is missing, then the range will block
        close(c)
    }()

    // when the channel is closed, range will return 0, and break for
    for i := range(c) {
        fmt.Println(i)
    }

    fmt.Println("finished")
}

func fibonacci(c , quit chan int) {
    x, y := 0, 1
    for {
        select {
            case c <- x:
                x, y = y, x + y
            case <-quit:
                fmt.Println("quit")
                return
        }
    }
}

func testSelect() {
    c := make(chan int)
    quit := make(chan int)

    go func() {
        for i := 1; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()

    fibonacci(c, quit)
}

func testTimeout() {
    c := make(chan string)
    go func() {
        time.Sleep(time.Second * 2)
        c <- "I'm late"
    }()

    select {
    case res := <-c:
        fmt.Println(res)
    // time.After will return current time then
    case <-time.After(time.Second * 1):
        fmt.Println("You are late")
    }
}

func main() {
    testChannel()

    testRange()

    testSelect()

    testTimeout()
}
