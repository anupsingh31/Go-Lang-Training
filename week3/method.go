package main

import "fmt"

func main() {
	var person1 = person{
		firstName: "Anupam",
		lastName:  "Singh",
		age:       21,
	}
	var person2 = person{
		firstName: "Amit",
		lastName:  "Raj",
		age:       21,
	}
	fmt.Println("person 1 is : ", person1)
	fmt.Println("person 1's full name is : ", person1.getFullName())
	fmt.Println("person 2's full name is : ", person2.getFullName())

}

func (p *person) getFullName() string {
	return p.firstName + " " + p.lastName
}

type person struct {
	firstName string
	lastName  string
	age       int
}
