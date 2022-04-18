package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	/* 	n := 44
	   	fibN := fib(n)
	   	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN) */
	n := 44
	go fib(n)
	n2 := 45
	go fib(n2)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fib(n))
	fmt.Printf("\rFibonacci(%d) = %d\n", n2, fib(n2))
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
