package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	storage map[string]int
}

type Semaphore interface {
	Acquire() error
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

var R int

func (s *semaphore) Acquire() error {
	R += step
	select {
	case s.chann <- R:
		return nil
	case <-time.After(1500 * time.Millisecond):
		return errors.New("file releaseA")
	}

}

func (s *semaphore) Release() error {

	select {
	case Cache1.storage[k1] = <-s.chann:
		//case Cache1.storage[k1] = <-s.chann:
		return nil
	case <-time.After(500 * time.Millisecond):
		return errors.New("file releaseR")
	}
}

func runTicket(i int, s Semaphore) {
	defer func() {
		if err := s.Release(); err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Get:", step*i)
}

var Cache1 = Cache{storage: make(map[string]int)}

func main() {

	s := NewSemaphore(3)
	for i := 0; i < 10; i++ {
		if err := s.Acquire(); err != nil {
			fmt.Printf("error: %s\n", err)
		}

		go runTicket(i, s)
	}
	for s.Len() > 0 {
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("k1 :", Cache1.Get(k1))

}

/* func main() {
	cache := Cache{storage: make(map[string]int)}
	for i := 0; i < 10; i++ {
		go func() {
			cache.Increase(k1, step)
			//time.Sleep(100 * time.Millisecond)
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			cache.Set(k1, step*i)
			//time.Sleep(100 * time.Millisecond)
		}(i)
	}

	fmt.Println(cache.Get(k1))

}
*/
