package main

import (
    "github.com/coreos/etcd/clientv3"
    "context"
    "time"
    "log"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
    defer cancel()
    endpoints := []string{"127.0.0.1:2379", "127.0.0.1:22379", "127.0.0.1:32379"}
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   endpoints,
        DialTimeout: 2 * time.Second,

    })
    if err != nil {
        log.Fatal(err)
    }
    defer cli.Close()

    // print cluster counters
    log.Println("cnt:", cli.Endpoints())

    // minimum lease TTL is 5-second
    resp, err := cli.Grant(ctx, 5)
    if err != nil {
        log.Fatal(err)
    }

    // after 5 seconds, the key 'foo' will be removed
    _, err = cli.Put(context.TODO(), "foo", "bar", clientv3.WithLease(resp.ID))
    if err != nil {
        log.Fatal(err)
    }

    ch, err := cli.KeepAlive(context.TODO(), resp.ID)
    if err != nil {
        log.Fatal(err)
    }

    ttl :=  <- ch
    log.Println("ttl", ttl.TTL)

    for i := 1; i <= 4; i++ {
        time.Sleep(time.Second * 4)
        gresp, err := cli.Get(context.TODO(), "foo")
        if err != nil {
            log.Fatal(err)
        }
        log.Println(len(gresp.Kvs))
    }

    _, err = cli.Revoke(context.TODO(), resp.ID)
    if err != nil {
        log.Fatal(err)
    }

    gresp, err := cli.Get(context.TODO(), "foo")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(len(gresp.Kvs))

    _, err = cli.Put(context.TODO(), "foo", "bar")
    if err != nil {
        log.Fatal(err)
    }

    gresp, err = cli.Get(context.TODO(), "foo")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(len(gresp.Kvs))
}
