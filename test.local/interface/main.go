package main

import (
	"fmt"
)

// Static structure type to define Person attributes
type Person struct {
	firstname string
	lastname  string
}

// Static structure type to define SecretAgent attibutes
type SecretAgent struct {
	Person
	ltk bool
}

// Speak method for a SecretAgent
func (s SecretAgent) Speak() {
	fmt.Println("[Secret Agent]: I am", s.lastname, "...", s.firstname, s.lastname)
}

// Speak method for a simple Person
func (p Person) Speak() {
	fmt.Println("[Person]: I am", p.firstname, p.lastname)
}

// Polymorphism type
type Human interface {
	Speak()
}

// Function to check if Person and SecretAgent are Human type as well
func Checker(h Human) {
	// User switch function to verify data type
	switch h.(type) {
	case Person:
		fmt.Println("[Checker]: The simple person", h.(Person).firstname, h.(Person).lastname, "is a human")
	case SecretAgent:
		fmt.Println("[Checker]: The secret agent", h.(SecretAgent).firstname, h.(SecretAgent).lastname, "is a human")
	}
}
func main() {
	sa1 := SecretAgent{
		Person: Person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	p1 := Person{
		firstname: "John",
		lastname:  "Doe",
	}
	fmt.Println("============ Print Data Structure ============")
	fmt.Println("SecretAgent:", sa1)
	fmt.Println("Person:", p1)
	fmt.Println("==============================================")
	fmt.Println("************ Conversation ************")
	fmt.Println("What's your name misters ?")
	sa1.Speak()
	p1.Speak()
	fmt.Println("**************************************")
	fmt.Println("############# Checker #############")
	Checker(sa1)
	Checker(p1)
}
