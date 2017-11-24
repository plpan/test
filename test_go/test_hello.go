package main

func add(x, y int) int {
    z := x + y
    return z
}

func main() {
    x := 0x10
    y := 0x20
    go add(x, y)
}
