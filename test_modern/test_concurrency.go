package main

import (
	"context"
	"fmt"
	"time"

	"github.com/modern-go/concurrent"
)

func test_normal() {
	e := concurrent.NewUnboundedExecutor()
	e.Go(func(ctx context.Context) {
		fmt.Println("ok")
	})
	time.Sleep(time.Second)
	e.StopAndWaitForever()
	fmt.Println("executor stoped")
}

func test_special() {
	e := concurrent.NewUnboundedExecutor()
	e.Go(func(ctx context.Context) {
		timer := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-timer.C:
				fmt.Println("timer")
			case <-ctx.Done():
				fmt.Println("goroutine exit")
				return
			}
		}
	})
	time.Sleep(time.Second)

	e.StopAndWaitForever()
	fmt.Println("executer stoped")
}

func main() {
	test_normal()
	test_special()
}
