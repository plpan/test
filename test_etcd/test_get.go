package main

import (
    "github.com/coreos/etcd/clientv3"
    "context"
    "time"
    "fmt"
)

func main() {

    endpoints := []string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"}
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   endpoints,
        DialTimeout: 2 * time.Second,
    })
    ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
    resp, err := cli.Get(ctx, "ticket/57-172.24.24.27_1510817712_NTctMTcyLjI0LjI0LjI3ODE3NzEy1B2M2Y8AsgTpgAmY7PhCfg%3D%3D")
    defer cancel()

    fmt.Println(len(resp.Kvs) != 0, err)
}
