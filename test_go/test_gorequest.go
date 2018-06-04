package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

func testRequest() {
	goreq := gorequest.New()
	goreq.Transport.MaxIdleConns = 2
	goreq.Transport.IdleConnTimeout = 20 * time.Second
	goreq.Transport.MaxIdleConnsPerHost = 2
	goreq.Get("http://www.baidu.com").End()
}

func testHttpReq() {
	tr := &http.Transport{
		MaxIdleConns:        4,
		IdleConnTimeout:     2 * time.Second,
		DisableCompression:  true,
		MaxIdleConnsPerHost: 2,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://www.baidu.com")
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%#v\n", resp.Body)
	fmt.Printf("%#v %v\n", string(body[:100]), err)
}

func main() {
	for {
		go testRequest()
		//go testHttpReq()
		time.Sleep(1 * time.Second)
		fmt.Println("Request done, open files =", countOpenFiles())
	}
}

func countOpenFiles() int {
	out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -p %v", os.Getpid())).Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(out), "\n")
	return len(lines) - 1
}
