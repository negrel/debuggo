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

	debugo.Run(func() {
		logFib(n, result)
	})

	return result
}

// This function will be removed of the final binaries
// if the compiler inline it in the debugo run.
func logFib(n, result int) {
	log.Printf("[DEBUGO] - Fib %v : %v\n", n, result)
}
