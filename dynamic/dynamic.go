package main

import (
	"fmt"
	"time"
)

var m = make(map[int]int)

func fib(n int) int {
	if n < 2 {
		return n
	}
	if val, ok := m[n]; ok {
		return val
	}
	m[n] = fib(n-1) + fib(n-2)
	return fib(n-1) + fib(n-2)
}

func oldFib(n int) int {
	if n < 2 {
		return n
	}
	return oldFib(n-1) + oldFib(n-2)
}

func wrapTime(f func(int) int) func(int) (int, time.Duration) {
	return func(n int) (int, time.Duration) {
		now := time.Now()
		val := f(n)
		elapsed := time.Since(now)
		return val, elapsed
	}
}

func main() {
	decoratedFib := wrapTime(fib)
	fib, elapsed := decoratedFib(12)
	fmt.Printf("fib of 12 is: %d and took %s\n", fib, elapsed)
	fib, elapsed = decoratedFib(15)
	fmt.Printf("fib of 15 is: %d and took %s\n", fib, elapsed)
	fib, elapsed = decoratedFib(150)
	fmt.Printf("fib of 150 is: %d and took %s\n", fib, elapsed)
	fmt.Printf("================================\n")
	decoratedOldFib := wrapTime(oldFib)
	fib, elapsed = decoratedOldFib(12)
	fmt.Printf("fib of 12 is: %d and took %s\n", fib, elapsed)
	fib, elapsed = decoratedOldFib(15)
	fmt.Printf("fib of 15 is: %d and took %s\n", fib, elapsed)
	ch := make(chan string)
	go func() {
		fib, elapsed := decoratedOldFib(150)
		ch <- fmt.Sprintf("fib of 150 is: %d and took %s\n", fib, elapsed)
	}()
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(5 * time.Second):
		fmt.Println("Timed out")
	}

}
