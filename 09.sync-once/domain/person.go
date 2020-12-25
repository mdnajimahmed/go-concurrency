package domain

import "sync"

// Person is a struct
type Person struct {
	firstName string
	lastName  string
}

var person *Person
var once sync.Once

func initPerson() {
	person = &Person{
		firstName: "Md Najim",
		lastName:  "Ahmed",
	}
}

// NewPerson is to make person singleton
func NewPerson() *Person {
	once.Do(initPerson)
	return person
}
