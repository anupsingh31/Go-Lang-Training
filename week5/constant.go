package main

import (
	"fmt"
	"reflect"
)

const myConstant = 30
const (
	cat int = iota + 1
	dog
	camel
	horse
)

func main() {
	var num int
	fmt.Println("Type of our constant is : ", reflect.TypeOf(myConstant))
	fmt.Println("Value of my constant is :", myConstant)
	fmt.Println(num == cat)
	fmt.Println(cat, dog, camel, horse)
}
