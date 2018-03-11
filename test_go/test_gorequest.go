package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
)

var wg sync.WaitGroup

func testRequest() {
	r := gorequest.New().Timeout(time.Millisecond * 1000).Get("http://www.baidu.com")
	_, _, err := r.End()
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}

func main() {
	for i := 0; i < 1000; i = i + 1 {
		wg.Add(1)
		go testRequest()
	}
	wg.Wait()
}
