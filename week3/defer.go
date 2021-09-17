package main

import "fmt"

func main() {
	var number = 50
	f1()
	defer f2(number)
	number = 100
	fmt.Println("end of main number is :", number)

}

func f1() {
	defer fmt.Println("hello from f1")
	fmt.Println("end of f1")
}

func f2(number int) {
	fmt.Println("end of f2 number is : ", number)
}
