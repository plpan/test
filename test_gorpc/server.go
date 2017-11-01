package main

import (
	"github.com/valyala/gorpc"
	"log"
	"strings"
)

func main() {
	s := &gorpc.Server{
		// tcp addr to listen
		Addr: ":12345",

		// handle
		Handler: func(clientAddr string, request interface{}) (response interface{}) {
			log.Printf("Received request %+v from the client %s\n", request, clientAddr)
			return strings.ToLower(request.(string))
		},
	}

	if err := s.Serve(); err != nil {
		log.Fatalf("Cannot start rpc server: %s", err)
	}
	//time.Sleep(time.Second * 2)
	//s.Stop()
}