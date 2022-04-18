package main

import (
	"sync"
	"time"
)

type Cache struct {
	sync.Mutex
	storage map[string]int
}

func (c *Cache) Increase(key string, value int) {
	c.Lock()
	defer c.Unlock()
	time.Sleep(1 * time.Millisecond)
	c.storage[key] += value
}

func (c *Cache) Set(key string, value int) {
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	delete(c.storage, key)
}
