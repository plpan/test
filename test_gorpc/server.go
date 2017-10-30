package main

import (
	"github.com/valyala/gorpc"
	"log"
)

func main() {
	s := &gorpc.Server{
		// tcp addr to listen
		Addr: ":12345",

		// handle
		Handler: func(clientAddr string, request interface{}) (response interface{}) {
			log.Printf("Received request %+v from the client %s\n", request, clientAddr)
			return request
		},
	}

	if err := s.Serve(); err != nil {
		log.Fatalf("Cannot start rpc server: %s", err)
	}
}