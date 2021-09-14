package main

import (
	"fmt"
	"strings"
)

func main() {
	name, size := doSomething("Anupam", "Singh")
	fmt.Println(name, " ", size)
}

func doSomething(firstName string, lastName string) (string, int) {
	name := firstName + " " + lastName
	size := len(name)
	return strings.ToUpper(name), size
}
