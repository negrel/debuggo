package person

import "fmt"

var _ fmt.Stringer = &Pet{}

// Pet define a pet.
type Pet struct {
	Name  string
	Breed string
	owner *Person
}

func (p *Pet) String() string {
	return fmt.Sprintf("%v (%v)", p.Name, p.Breed)
}
