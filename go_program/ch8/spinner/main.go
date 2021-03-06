// Author: magician
// File:   demo1.go
// Date:   2020/7/20
package main

import (
	"fmt"
	"time"
)

func main()  {
	go spinner(100 * time.Microsecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\nFibnacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int  {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

