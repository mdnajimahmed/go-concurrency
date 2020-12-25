package domain2

// Person is a struct
type Person struct {
	firstName string
	lastName  string
}

var person *Person

func init() {
	person = &Person{
		firstName: "Md Najim",
		lastName:  "Ahmed",
	}
}

// NewPerson is to make person singleton
func NewPerson() *Person {
	return person
}
