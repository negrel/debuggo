package main

import (
	"github.com/negrel/debuggo/examples/assert/debug/assert"
)

func main() {
	var X, Y int
	X = 0
	Y = 1

	X += Y
	assert.Equal(X == Y, "X and Y should be equal")

	// Should panic if compiled with '-tags assert'
	X *= 2
	assert.Equal(X == Y, "X and Y should be equal")
}
