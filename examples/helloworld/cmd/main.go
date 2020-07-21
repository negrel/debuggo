package main

import (
	"fmt"

	"github.com/negrel/debuggo/examples/helloworld/internal/person"
)

func main() {
	bob := person.New("Bob")

	cat := &person.Pet{
		Name:  "Snowball",
		Breed: "Persian",
	}

	bob.Adopt(cat)

	fmt.Println(bob)
}
