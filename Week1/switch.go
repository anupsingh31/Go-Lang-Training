package main

import "fmt"

func main() {
	cgpa := 8
	switch cgpa {
	case 10:
		fmt.Println("Excellent")
	case 9, 8:
		fmt.Println("Very Good")
	case 7:
		fmt.Println("Good")
	default:
		fmt.Println("Inside Default")
	}
}
