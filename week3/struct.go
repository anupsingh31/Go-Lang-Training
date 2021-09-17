package main

import "fmt"

func main() {
	var person1 = person{
		firstName: "Anupam",
		lastName:  "Singh",
		age:       21,
	}
	fmt.Println(person1)
	person2 := newPerson("Arish", "Raj", 13)
	fmt.Println(*person2)
}

func newPerson(fName, lName string, age int) *person {
	var personType = &person{
		firstName: fName,
		lastName:  lName,
		age:       age,
	}
	return personType
}

type person struct {
	firstName string
	lastName  string
	age       int
}
