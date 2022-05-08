package main

import (
	"fmt"
	"sync"
)

var cPool = sync.Pool{
	New: func() interface{} { return Counter{} },
}

type Counter struct {
	А int
	В int
}

func getInt() (c *Counter) {
	data := cPool.Get().(*Counter)
	if data != nil {
		c = data
	}
	return
}

func PutInt(i int) {
	cPool.Put(&Counter{i, i})
}

func SetInt(i int) (c *Counter) {
	data := &Counter{i, i}
	c = data
	return c
}

func main() {
	for i := 0; i < 10; i++ {
		PutInt(i)
		fmt.Println(getInt())
		SetInt(i)
		//fmt.Println(SetInt(i))
	}

}
