package main

import (
    "github.com/coreos/etcd/clientv3"
    "context"
    "time"
    "log"
)

func main() {

    endpoints := []string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"}
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   endpoints,
        DialTimeout: 3 * time.Second,
    })

    ctx, cancel := context.WithTimeout(context.Background(), 46 * time.Second)
    defer cancel()

    // minimum lease TTL is 5-second
    resp, err := cli.Grant(ctx, 5)
    if err != nil {
        log.Fatal(err)
    }

    // after 5 seconds, the key 'foo' will be removed
    _, err = cli.Put(ctx, "foo", "bar", clientv3.WithLease(resp.ID))
    if err != nil {
        log.Fatal(err)
    }

    ch, err := cli.KeepAlive(ctx, resp.ID)
    if err != nil {
        log.Fatal(err)
    }

    for i := 0; i <= 5; i++ {
        ttl :=  <- ch
        log.Println("ttl", ttl.TTL)
    }
}
