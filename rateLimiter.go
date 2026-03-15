package main

import (
	"sync"
	"time"
)

type RateLimiter struct {
    mu          sync.Mutex
    maxCapacity int
    inUse       int
    ticker      *time.Ticker
    done        chan struct{}
}

func NewRateLimiter(rps int) *RateLimiter {
    r := &RateLimiter{
        maxCapacity: rps,
        ticker:      time.NewTicker(time.Second),
        done:        make(chan struct{}),
    }
    go func() {
        for {
            select {
            case <-r.ticker.C:
                r.mu.Lock()
                r.inUse = 0
                r.mu.Unlock()
            case <-r.done:
                return
            }
        }
    }()
    return r
}

func (r *RateLimiter) Allow() bool {
    r.mu.Lock()
    defer r.mu.Unlock()
    if r.inUse < r.maxCapacity {
        r.inUse++
        return true
    }
    return false
}

func (r *RateLimiter) Stop() {
    r.ticker.Stop()
    close(r.done)
}
