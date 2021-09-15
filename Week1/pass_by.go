package main

import "fmt"

func main() {
	number := 20
	case1(number)
	fmt.Println("address of number in main : ", &number)
	fmt.Println("After case1, in main1", number)
	case2(&number)
	fmt.Println("After case2, in main1", number)
}

// pass by value

func case1(number int) {
	number += 2
	fmt.Println("In case1, value is: ", number)
	fmt.Println("In case1, address is: ", &number)
}

// pass by reference

func case2(number *int) {
	*number += 2
	fmt.Println("In case2, value is: ", number)
	fmt.Println("In case2, address is: ", &number)
}
