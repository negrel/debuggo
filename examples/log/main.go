package main

import (
	"log"

	"github.com/negrel/debugo/pkg/debugo"
)

func main() {
	result := fib(10)
	log.Printf("Final result: %v", result)
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	result := fib(n-1) + fib(n-1)

	debugo.Printf("[DEBUGO] - Fib %v : %v\n", n, result)

	return result
}
