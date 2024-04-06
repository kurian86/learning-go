package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	p := MakePerson("Angel", "Siles", 36)
	fmt.Println(p)

	p2 := MakePersonPointer("Angel", "Siles", 36)
	fmt.Println(p2)
}
