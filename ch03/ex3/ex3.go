package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	em1 := Employee{
		"John",
		"Doe",
		1,
	}
	em2 := Employee{
		firstName: "Jane",
		lastName:  "Rodriguez",
		id:        2,
	}
	var em3 Employee
	em3.firstName = "David"
	em3.lastName = "Ortiz"
	em3.id = 3

	fmt.Println(em1, em2, em3)
}
