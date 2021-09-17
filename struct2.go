package main

import (
	"fmt"
	"structure"
)

func main() {
	person1 := structure.NewPerson("Anupam", "Singh", 21)
	fmt.Println(person1)
	fmt.Println("Id of first person : ", person1.EmployeeId)
	fmt.Println("First name of person : ", person1.FirstName)
}
