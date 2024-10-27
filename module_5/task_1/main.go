package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go spinner(10 * time.Second)
	n := 33
	channel1 := make(chan int)
	channel2 := make(chan int)
	go func() {
		fibN := fib(n)
		channel1 <- fibN
	}()

	go func() {
		fibM := fib(n - 1)
		channel2 <- fibM
	}()

	fibN := <-channel1
	fibM := <-channel2
	fibNew := fibN + fibM
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	fmt.Printf("\rFibonacci(%d) = %d\n", n+1, fibNew)
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
	a, b := 0, 1
	for i := 2; i <= x; i++ {
		a, b = b, a+b
	}
	return b
}
