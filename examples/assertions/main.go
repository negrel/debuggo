package main

import (
	"log"
	"math/rand"

	"github.com/negrel/debugo/pkg/debugo"
)

var randomInt = func() func(n int) int {
	r := rand.New(rand.NewSource(99))

	return r.Intn
}()

func main() {
	rand := randomInt(100)

	debugo.Assert(func() (bool, string) {
		log.Println(rand)

		return rand > 50, "the random number should be greater than 50"
	}())

	log.Println("End of the program")
}
