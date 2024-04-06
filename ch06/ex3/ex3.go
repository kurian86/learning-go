package main

import (
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		firstName,
		lastName,
		age,
	}
}

func main() {
	var people = make([]Person, 0, 10_000_000)

	fmt.Println("Start:", time.Now())

	for i := 0; i < 10_000_000; i++ {
		people = append(people, MakePerson("Angel", "Siles", 36))
	}

	fmt.Println("End:", time.Now())
	fmt.Println(len(people))
}
