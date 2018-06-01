package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
)

var wg sync.WaitGroup

func testRequest() {
	r := gorequest.New()
	r.Transport.DisableKeepAlives = true
	r.Timeout(time.Millisecond * 1000).Get("http://www.baidu.com")
	_, _, err := r.End()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Request done, open files =", countOpenFiles())

	wg.Done()
}

func countOpenFiles() int {
	out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -p %v", os.Getpid())).Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(out), "\n")
	return len(lines) - 1
}

func main() {
	for i := 0; i < 1000; i = i + 1 {
		wg.Add(1)
		go testRequest()
	}
	wg.Wait()
}
