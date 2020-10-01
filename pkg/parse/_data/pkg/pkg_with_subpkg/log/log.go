package log

import (
	"os"
)

// Print the given msg to stdout.
func Print(msg string) {
	_, _ = os.Stdout.Write([]byte(msg))
}
