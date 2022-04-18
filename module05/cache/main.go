package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	storage map[string]int
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

func main() {
	cache := Cache{storage: make(map[string]int)}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) //добавляем по 1 на каждую горутину
		go func() {
			defer wg.Done() //вызываем Done, когда горутина закончит выполнение
			cache.Increase(k1, step)
			//time.Sleep(100 * time.Millisecond)
		}()
	}
	for i := 0; i < 10; i++ {
		i := i //copy variable
		wg.Add(1)
		go func() {
			defer wg.Done()
			cache.Set(k1, step*i)
			//time.Sleep(100 * time.Millisecond)
		}()
	}
	wg.Wait() //ждем окончания работы всех горутин
	fmt.Println(cache.Get(k1))
}
