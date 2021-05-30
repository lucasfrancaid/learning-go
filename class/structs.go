package class

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func Structs() {
	// use person strict to declare variable
	var person Person
	person.Name = "Lucas"
	person.Age = 25
	fmt.Println("[use person strict to declare variable] ->", person)

	// self assign person from struct Person
	otherPerson := Person{
		Name: "Vinicius",
		Age:  25,
	}
	fmt.Println("[self assign person from struct Person] ->", otherPerson)

}
