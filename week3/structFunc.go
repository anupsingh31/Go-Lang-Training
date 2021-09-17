package main

import "fmt"

func main() {
	var person1 = person{
		firstName: "Anupam",
		lastName:  "Singh",
		age:       21,
		addresses: []string{"mumbai", "pune"},
	}
	var person2 = person{
		firstName: "Saket",
		lastName:  "Mishra",
		age:       22,
		addresses: []string{"mumbai", "Thane"},
	}
	var person3 = person{
		firstName: "Raj",
		lastName:  "Shah",
		age:       24,
		addresses: []string{"amber", "pune"},
	}
	fmt.Println("before calling : ", person1)
	// modifyPerson(person1)
	// fmt.Println("After calling ModifyPerson: ", person1)
	// modifyPersonByPointer(&person1)
	// fmt.Println("After calling ModifyPerson by pointer: ", person1)

	var allPeroson = []person{person1, person2, person3}
	for _, newPerson := range allPeroson {
		fmt.Println("First Name : ", newPerson.firstName)
		fmt.Println("Last Name : ", newPerson.lastName)
		fmt.Println("age : ", newPerson.age)
		for index, newaddress := range newPerson.addresses {
			fmt.Println("Address ", index+1, " : ", newaddress)
		}

	}
}

func modifyPerson(p person) {
	p.firstName = "newName"
	p.age = 25
	fmt.Println("Inside Modify : ", p)
}

func modifyPersonByPointer(p *person) {
	p.firstName = "newName"
	p.age = 25
	fmt.Println("Inside Modify : ", *p)
}

type person struct {
	firstName string
	lastName  string
	age       int
	addresses []string
}
