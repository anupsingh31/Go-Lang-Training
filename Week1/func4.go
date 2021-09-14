package main

import "fmt"

func main() {
	fmt.Println(mathOperation(add))
	fmt.Println(mathOperation(subtract))

}

func add(num1, num2 int) int {
	return num1 + num2
}

func subtract(num1, num2 int) int {
	return num1 - num2
}

func mathOperation(funcName func(num1, num2 int) int) int {
	return funcName(20, 10)
}
