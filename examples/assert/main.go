package main

import (
	"errors"
	"fmt"

	"github.com/negrel/debuggo/pkg/assert"
)

func cureCOVID() (int, error) {
	return -1, errors.New("failed to cure COVID")
}

func main() {
	days, err := cureCOVID()
	assert.Nil(err, err)

	fmt.Printf("it took %v days to cure the COVID.\n", days)
}
