// +build debugo

package debugo

import "fmt"

// Print is a shortcut for debug.Run(func() { fmt.Print() })
func Print(a ...interface{}) {
	Run(func() {
		fmt.Print(a...)
	})
}

// Printf is a shortcut for debug.Run(func() { fmt.Printf() })
func Printf(format string, a ...interface{}) {
	Run(func() {
		fmt.Printf(format, a...)
	})
}

// Println is a shortcut for debug.Run(func() { fmt.Println() })
func Println(a ...interface{}) {
	Run(func() {
		fmt.Println(a...)
	})
}
