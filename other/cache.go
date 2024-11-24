package main

import (
	"math/rand"
	"sync"
	"time"
)

type Cache interface {
	Get() int
	Invalidate(value int)
}

type MyCache struct {
	value int
	m     sync.RWMutex
}

func (c *MyCache) Get() int {
	defer c.m.RUnlock()
	c.m.RLock()
	return c.value
}

func (c *MyCache) Invalidate(value int) {
	defer c.m.Unlock()
	c.m.Lock()
	c.value = value
}

func NewCache() Cache {
	return &MyCache{
		value: 1,
		m:     sync.RWMutex{},
	}
}

func long() int {
	time.Sleep(time.Second * 4)

	return rand.Intn(10)
}

func main() {
	cache := NewCache()

	ticker := time.NewTicker(time.Minute * 5)

	go func() {
		for {
			select {
			case <-ticker.C:
				cache.Invalidate(long())
			}
		}
	}()

}
