package connect

import (
    "github.com/silenceper/pool"
    "fmt"
    "time"
)

func InitThread(min, max int, factory func() (interface {}, error), close func(v interface{}) error) (pool.Pool, error) {
    poolConfig := &pool.PoolConfig {
        InitialCap : min,
        MaxCap : max,
        Factory : factory,
        Close : close,
        IdleTimeout : 5 * time.Second,
    }
    p, err := pool.NewChannelPool(poolConfig)
    if err != nil {
        fmt.Printf("create pool failed: %v", err)
        return nil, err
    }
    return p, nil
}
