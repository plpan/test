package main

import (
	"github.com/valyala/gorpc"
	"log"
	"strings"
)

func main() {
	c := &gorpc.Client{
		// tcp address to connect
		Addr: "127.0.0.1:12345",
	}
	c.Start()

	message := "Hello, World"
	resp, err := c.Call(message)
	if err != nil {
		log.Fatalf("Error when sending message to server: %s", err)
	}
	if resp.(string) != strings.ToLower(message) {
		log.Fatalf("Error response received from server: %v", resp)
	}
	log.Printf("response: %v", resp)
}
