package person

import (
	"fmt"
	"math"
)

const x = math.Pi

// Println formats using the default formats for its operands
// and writes to standard output. Spaces are always added
// between operands and a newline is appended. It returns
// the number of bytes written and any write error encountered.
func Println(a ...interface{}) {
	fmt.Println(a...)
}

// String return the string of a Stringer object.
func String(a fmt.Stringer) string {
	return a.String()
}

// PrintPi print the math.Pi constant.
func PrintPi() {
	fmt.Println(x)
}
