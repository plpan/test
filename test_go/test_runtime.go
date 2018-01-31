package main

import (
	"fmt"
	"runtime"
)

func Test_Maxprocs() {
	fmt.Println("---Test_Maxprocs---")
	// CPU核数限定了goroutine的并行度，多个goroutine同时执行
	runtime.GOMAXPROCS(2)
	// go程序默认只开一个CPU核；除非调用GOMAXPROCS设置最大的CPU核数，这样多个goroutine就可以同时运行了
	fmt.Printf("num goroutine: %d\n", runtime.NumGoroutine())

	// 默认情况下，这个goroutine是没有机会运行的,除非main goroutine执行了GoSched
	go func() {
		fmt.Println("another goroutine")
	}()

	// 这里需要轮训次数多一些，不然上面的goroutine还是没机会执行
	for i := 0; i < 500000; i++ {
		//fmt.Println(i)
		if i == 1 {
			// 让出当前CPU，当CPU核数为1时，上面的goroutine就可以去执行了
			// Gosched：这个函数的作用是让当前 goroutine 让出 CPU，当一个 goroutine 发生阻塞，
			// Go 会自动地把与该 goroutine 处于同一系统线程的其他 goroutine 转移到另一个系统线程上去，以使这些 goroutine 不阻塞
			// runtime.Gosched()
		}
	}
}

func Test_FuncForPC() {
	fmt.Println("---Test_FuncForPC---")
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime caller failed")
		return
	}
	fmt.Printf("file: %v line: %v\n", file, line)

	f := runtime.FuncForPC(pc)
	if f == nil {
		fmt.Println("funcforpc nil")
	}

	fmt.Printf("funcname: %v\n", f.Name())
}

func main() {
	Test_Maxprocs()

	Test_FuncForPC()
}
