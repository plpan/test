package main

import (
	"crypto/tls"
	"fmt"

	"github.com/pplonepiece/test/test_thrift/simple/gen-go/thrift/example"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func handleClient(client *example.AdderClient) error {
	var a int32 = 1
	var b int32 = 3
	sum, err := client.Add(a, b)
	if err != nil {
		fmt.Println("error sum: ", err)
	}
	fmt.Printf("sum of %d, %d is : %d\n", a, b, sum)

	an := &example.Numbers{1, 2}
	bn := &example.Numbers{3, 4}
	sn, err := client.AddNumber(an, bn)
	if err != nil {
		fmt.Println("error sum: ", err)
	}
	fmt.Printf("sum of %#v, %#v is : %#v\n", an, bn, sn)

	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TTransport
	var err error

	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("error opening socket: ", err)
		return err
	}
	transport = transportFactory.GetTransport(transport)

	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err
	}

	return handleClient(example.NewAdderClientFactory(transport, protocolFactory))
}
