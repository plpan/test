package main

import (
    "google.golang.org/grpc"
    pb "test/test_grpc"
    "test/test_grpc/connect"
    "test/test_grpc/request"
    "fmt"
    "sync"
)

const (
    address = "127.0.0.1:10023"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        fmt.Printf("connect to %v failed.", address)
    }
    defer conn.Close()

    // create connection
    factory := func() (interface{}, error) {
        return pb.NewDataClient(conn), nil
    }

    // close
    close := func(v interface{}) error {
        return conn.Close()
    }

    // initialize connection pool
    p, err := connect.InitThread(10, 30, factory, close)
    if err != nil {
        fmt.Println("initial thread pool failed.")
        return
    }

    var wg sync.WaitGroup
    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // get the connection
            v, _ := p.Get()
            client := v.(pb.DataClient)
            info := &pb.URequest {
                Uid: 10012,
            }
            request.GetUserInfo(client, info)
            // return connection
            p.Put(v)
        }()
        wg.Wait()
    }

    for i := 0; i < 20; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // get the connection
            v, _ := p.Get()
            client := v.(pb.DataClient)
            request.ChangeUserInfo(client)
            // return connection
            p.Put(v)
        }()
        wg.Wait()
    }

    // get the connectionb pool size
    current := p.Len()
    fmt.Printf("len: %v", current)
}
