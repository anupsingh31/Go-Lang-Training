package structure

func NewPerson(fName, lName string, age int) employee {
	var personType = employee{
		EmployeeId: 101,
		person: person{
			FirstName: fName,
			lastName:  lName,
			age:       age,
		},
	}

	return personType
}

type employee struct {
	EmployeeId int
	person
}

type person struct {
	FirstName string
	lastName  string
	age       int
}
