package main

import "fmt"

func main() {
	var userNumber int
test:
	fmt.Println("Enter the number")
	fmt.Scanln(&userNumber)
	fmt.Println("The number is ", userNumber)
	if userNumber > 200 {
		fmt.Println("Return")
		return
	}

	if userNumber > 100 {
		fmt.Println("Number is very High")
		goto test
	}
	fmt.Println("Number is in range")

}
