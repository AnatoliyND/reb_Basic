package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	storage map[string]int
}

type Semaphore interface {
	Acquire(i int) error
	Release() error
	Len() int
}

type semaphore struct {
	chann chan int
}

const (
	k1   = "key1"
	step = 7
)

func (c *Cache) Increase(key string, value int) {
	c.Lock()
	defer c.Unlock()
	time.Sleep(1 * time.Millisecond)
	c.storage[key] += value
}

func (c *Cache) Set(key string, value int) {
	c.Lock()
	defer c.Unlock()
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	c.RLock()
	defer c.RUnlock()
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	c.RLock()
	defer c.RUnlock()
	delete(c.storage, key)
}

func (s *semaphore) Len() int {
	return len(s.chann)
}

func NewSemaphore(tickets int) Semaphore {
	return &semaphore{chann: make(chan int, tickets)}
}

func (s *semaphore) Acquire(i int) error {

	s.chann <- i * step
	return nil

}

func (s *semaphore) Release() error {

	Cache1.storage[k1] = <-s.chann
	//fmt.Println("Get:", Cache1.storage[k1])
	return nil
}

func runTicket(i int, s Semaphore) {
	defer func() {
		if err := s.Release(); err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}()
	time.Sleep(30 * time.Millisecond)
	//fmt.Println("Get:", Cache1.storage[k1])
}

var Cache1 = Cache{storage: make(map[string]int)}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)
	s := NewSemaphore(3)
	go func() {
		for i := 0; i < 10; i++ {
			if err := s.Acquire(i); err != nil {
				fmt.Printf("error: %s\n", err)
			}
			go runTicket(i, s)
		}
		for s.Len() > 0 {
			time.Sleep(1 * time.Millisecond)
		}

	}()

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context deadline exceeded")
			break LOOP
		default:

			fmt.Println("waiting")
			time.Sleep((time.Millisecond * 20))
		}
	}
	fmt.Println("k1 :", Cache1.Get(k1))
}
