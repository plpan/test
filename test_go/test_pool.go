package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return 1
		},
	}
	fmt.Println(pool.Get()) // 1, gen by New func
	pool.Put(2)
	fmt.Println(pool.Get()) // 2, gen by Put
}
