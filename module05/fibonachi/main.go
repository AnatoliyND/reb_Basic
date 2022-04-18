package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go spinner(100 * time.Millisecond)
	n := 32
	go fib(n)
	go fib(n + 1)

	time.Sleep(1 * time.Second)

	fmt.Printf("\rFibonacci(%d) = %d\n", n, fib(n))
	fmt.Printf("\rFibonacci(%d) = %d\n", n+1, fib(n+1))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)

}
