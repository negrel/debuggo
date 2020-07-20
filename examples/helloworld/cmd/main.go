package main

import (
	"log"

	"github.com/negrel/debuggo/examples/helloworld/internal/debug/person"
	"github.com/negrel/debuggo/examples/helloworld/internal/person"
)

func main() {
	bob := person.New("Bob")

	cat := &person.Pet{
		Name:  "Snowball",
		Breed: "Persian",
	}

	bob.Adopt(cat)

	log.Println(bob)
}
