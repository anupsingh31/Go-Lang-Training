package main

import "fmt"

var num int

func main() {
	num = 40
	display()
	fmt.Println("In main", num)
	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}
}

func display() {
	var num1 int
	num1 = 20
	fmt.Println(num1)
	fmt.Println("In display", num)
}
