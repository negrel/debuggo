package person

import (
	"fmt"

	debug "github.com/negrel/debuggo/examples/helloworld/internal/debug/person"
)

var _ fmt.Stringer = &Person{}

// Person define a person.
type Person struct {
	Name string
	Pets []*Pet
}

// New return a new person.
func New(name string) *Person {
	return &Person{
		Name: name,
		Pets: make([]*Pet, 0, 4),
	}
}

// Adopt the given pet.
func (p *Person) Adopt(pet *Pet) {
	p.Pets = append(p.Pets, pet)
	debug.Println(p.Name + " is adopting " + pet.Name)
}

func (p *Person) String() string {
	return fmt.Sprintf("%v own the following pets: %v", p.Name, p.Pets)
}
